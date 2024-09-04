package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	c := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-c.C:
			go Run()
		}
	}
}

func Run() {
	defer func() {
		if res := recover(); res != nil {
			fmt.Println("res:", res)

			// 进行处理
			defer func() {
				if res2 := recover(); res2 != nil {
					fmt.Println("res:", res2)
				}
			}()
			panic(errors.New("第二次panic"))
		}
	}()

	panic(errors.New("panic"))
}
