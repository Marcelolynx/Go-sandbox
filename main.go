package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("tarefa 1") //T1
	go task("tarefa 2") //T2
	task("tarefa 3")    //T3
}
