package main

import (
	"fmt"
	"time"
)

func sleep1(seconds int) {
	t := time.NewTimer(time.Duration(seconds) * time.Second) // создаем таймер, который вернет текущее время через указанное количество времени
	<-t.C                                                    // получаем текущее время, значит указанное время прошло
	t.Stop()
}

func sleep2(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {

	sleep1(2)
	fmt.Println("After 2 seconds")

	sleep2(3)
	fmt.Println("After 3 seconds")
}
