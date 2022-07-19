package main
// Given a string expression of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. You may return the answer in any order.
import (
	"fmt"
	"strconv"
)

func main() {
	exp := "2*3-4*5"
	fmt.Println(diffWaysToCompute(exp))
}
 

// Example 1:

// Input: expression = "2-1-1"
// Output: [0,2]
// Explanation:
// ((2-1)-1) = 0 
// (2-(1-1)) = 2
// Example 2:

// Input: expression = "2*3-4*5"
// Output: [-34,-14,-10,-10,10]
// Explanation:
// (2*(3-(4*5))) = -34 
// ((2*3)-(4*5)) = -14 
// ((2*(3-4))*5) = -10 
// (2*((3-4)*5)) = -10 
// (((2*3)-4)*5) = 10
 

// Constraints:

// 1 <= expression.length <= 20
// expression consists of digits and the operator '+', '-', and '*'.
// All the integer values in the input expression are in the range [0, 99].

func diffWaysToCompute(expression string) []int {
	n := len(expression)
	res := make([]int, 0)
	if len(expression) <=2 {
		val, _ := strconv.Atoi(expression)
		return []int {
			val,
		}
	}
    for i:=0; i<len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			left := diffWaysToCompute(expression[0:i])
			right := diffWaysToCompute(expression[i+1:n])
			fmt.Println(left)
			fmt.Println(right)
			for j:=0; j<len(left); j++ {
				for k:=0; k<len(right); k++ {
					if expression[i] == '+'  {
						res = append(res, left[j] + right[k])
					} else if expression[i] == '-' {
						res = append(res, left[j] - right[k])
					} else if expression[i] == '*' {
						res = append(res, left[j] * right[k])
					}
				}
			}
		}
	}
	return res
}