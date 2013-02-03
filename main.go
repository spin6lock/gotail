package main

import (
	"log"
	"bytes"
	"os"
	"fmt"
	"strings"
	"github.com/howeyc/fsnotify"
)

func LineCount(lines []byte) int {
	return bytes.Count(lines, []byte("\n"))
}

func GetFileSize(filename string) int {
	fileInfo, err := os.Stat(filename)
	if err != nil{
		log.Fatal(err)
	}
	return int(fileInfo.Size())
}

func ByteArrayToMultiLines(bytes []byte) []string{
	lines := string(bytes)
	return strings.Split(lines, "\n")
}

func ReadNBytes(filename string, start int, end int) []byte{
	fh, _ := os.Open(filename)
	defer fh.Close()
	fh.Seek(int64(start), 0)
	size := end - start + 1
	buff := make([]byte, size)
	fh.Read(buff)
	return buff
}

func ReadLastNLines(name string, n int) ([]string, error){
	curr := GetFileSize(name)
	var end int
	count := n
	result := make([]byte, n)
	for count > 0 && curr != 0{
		curr -= n
		end = curr + n - 1
		if curr < 0 {
			curr = 0
		}
		buff := ReadNBytes(name, curr, end)
		result = append(buff, result...)
		count -= LineCount(buff)
	}
	return ByteArrayToMultiLines(result), nil
}

func main(){
	filename := "test.log"
	result, _ := ReadLastNLines(filename, 10)
	fmt.Println(result)
	//size := GetFileSize(filename)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan int)
	go func(){
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				if ev.IsModify(){
					//NewSize := GetFileSize(filename)
					//read content between NewSize and size
					//update size 
					//print out latest line
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
		ch <-1
	}()

	err = watcher.Watch(filename)
	if err != nil {
		log.Fatal(err)
	}
	<-ch
	watcher.Close()
}
