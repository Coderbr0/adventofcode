package main

import (
	"fmt"
	"strconv"
)

//import (
//"fmt"
//"strconv"
//)

func main() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
