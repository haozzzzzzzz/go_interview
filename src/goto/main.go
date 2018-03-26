package main

import (
	"fmt"
	"time"
)

func main()  {
	for i := 0; i < 10; i ++ {
		loop:
			fmt.Println(i)
	}

	goto loop // failed.
	loop2:
		fmt.Println("1")
		time.Sleep(time.Second)
	goto loop2
}
