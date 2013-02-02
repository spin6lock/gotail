package main

import (
	"time"
	"log"
	"os"
)

func main(){
	fh, err := os.Create("test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	t := time.Tick(1 * time.Second)
	for now := range t{
		fh.WriteString(now.String())
	}
}
