package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() { solve() }
func solve() {
	reader := bufio.NewReader(os.Stdin)
	var str1, str2 string
	_, err := fmt.Fscan(reader, &str1, &str2)
	if err != nil {
		fmt.Println(err)
	}
	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)
	fmt.Println(num1 + num2)
}