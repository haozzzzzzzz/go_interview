package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("ABCDE", 4))
}

func convert(s string, numRows int) (res string) {
	var rows [][]uint8 = make([][]uint8, numRows)

	for col := 0 ; len(s) > 0; col ++{ // 逐列
		for row := 0; row < numRows && len(s) > 0; row ++ { // 逐行
			if rows[row] == nil {
				rows[row] = make([]uint8,0)
			}

			// 默认为0
			rows[row] = append(rows[row], 0)

			if numRows==1 {
				rows[row][col] = s[0]
				s = s[1:]
			} else if numRows > 1 {
				if col % (numRows - 1) == 0 {
					rows[row][col] = s[0]
					s = s[1:]
				} else if col % (numRows - 1) + row == (numRows-1) {
					rows[row][col] = s[0]
					s = s[1:]
				}
			}
		}
	}

	//for _, rowSlice := range rows {
	//	for _, col := range rowSlice {
	//		if col != 0 {
	//			fmt.Print(string(col))
	//		} else {
	//			fmt.Print(" ")
	//		}
	//	}
	//
	//	fmt.Print("\n")
	//}

	for _, rowSlice := range rows {
		for _, col := range rowSlice {
			if col != 0 {
				res += string(col)
			}
		}
	}

	return
}