package main

import "fmt"

/**
1024! 末尾有多少个0？

1024的阶乘末尾有多少个0，这个问题只要理清思想就很好解了。

有多少个0取决于有多少个10相乘，即1024拆成小单元后有多少个10。

由于10不是素数，所以直接用10进行计算的话会有很多问题，于是将10分解。 10可以分解成2*5,2和5都是素数，由于每2个相邻的数中一定包含2，所以只要计算出有多少个5就可以了（2会在5之后及时出现）。

于是解法如下：
5 10 15 20 25 30
是5的倍数的数有： 1024 / 5 = 204个

是25的倍数的数有：1024 / 25 = 40个

是125的倍数的数有：1024 / 125 = 8个

是625的倍数的数有：1024 / 625 = 1个

所以1024! 中总共有204+40+8+1=253个因子5。

即1024！后有253个0
 */

func totalZero(n int) int {
 	total := 0
 	for n > 5 {
 		n = ( n - (n%5) ) / 5
 		total += n
	}
	return total
}

func main()  {
	fmt.Println(totalZero(1024))
}
