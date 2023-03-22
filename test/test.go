package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string = "18446744073709551610"
	fmt.Println(strconv.Atoi(num))
}
