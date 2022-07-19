package main

import (
	"fmt"
)
func calculate(s string) int {
	if len(s) <= 0 {
		return 0
	}
    
    sign := 1
    num := 0
	res := 0

    for i:=0; i < len(s); i++{
		c := s[i]
        if isNum(c) {
            num = num * 10 + int(c - '0')
        } else if c == '(' {
			cnt := 0
			j := i
			for ; j < len(s); j++ {
				if s[j] == '(' {
					cnt++
				} else if s[j] == ')' {
					cnt--
				}
				if cnt == 0 {
					break
				}
			}
			fmt.Println(s[i+1:j-1])
			if i + 1 == j - 1 {
				num = int(s[i+1] - '0')
			} else {
				num = calculate(s[i+1:j-1])
				fmt.Printf("res: %d\n", num)
			}
			fmt.Println(num)
			i = j + 1
        }
        if c == '-' || c == '+' || i >= len(s) - 1  {
            res += num * sign
			fmt.Printf("%d %d ----- %d \n",num,sign, res)
			num = 0
			if c == '-' {
				sign = -1
			} else {
				sign = 1
			}
        } 
    }
	fmt.Println(res)
    return res
}

func isNum(i byte) bool {
    if i <= '9' && i >= '0' {
        return true
    }   
    return false
}


func main() {
	calculate("2-4-(8+2-6+(8+4-(1)+8-10))")
}