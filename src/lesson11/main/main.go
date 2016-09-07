package main

import (
	"fmt"
	"time"
)

// Пример использования горутины
// Вывод цифр от 0 до 9 в формате 0:0; 0:1 ....
// time.Sleep() используется для задержки вывода,
// каждой строчки.
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	//go f(0)
	// Без Scanln() результат не отобразиться в консоли,
	// поскольку горутина не дожидается пока завершиться функция
	var input string
	fmt.Scanln(&input)


	// Пример использования каналов
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println("Message from chan:", msg)


	// Пример использования буферизированного канала
	msgBuf := make(chan string, 2)
	msgBuf <- "buf"
	msgBuf <- "chan"

	fmt.Println(<-msgBuf)
	fmt.Println(<-msgBuf)


	// Пример использования оператора select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}

