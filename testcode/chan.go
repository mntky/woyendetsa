package main

import (
	"fmt"
	"time"
)

type Data struct {
	name string
}

func main() {
	ch := make(chan string)

	go func() {
		for i:=0; i<10; i++{
			fmt.Println(i)
			time.Sleep(1*time.Second)
		}
		ch <- "aaa"
	}()

	<-ch
}
