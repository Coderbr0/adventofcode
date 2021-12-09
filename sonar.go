package main // https://adventofcode.com/2021/day/1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strings"
)

func main() {
	strArr := ReadFile()
	// fmt.Println(strArr)
	var convertToSliceOfInt []int // We declare variable outside of for loop (outside the scope of for loop); now the slice of int can be accessible within the main function and not just limited to the for loop
	count := 0
	for _, v := range strArr { // Ranging over string slice to obtain string value (v)
		convertToInt, _ := strconv.Atoi(v) // Converting string to integer
		// fmt.Println(convertToInt)
		convertToSliceOfInt = append(convertToSliceOfInt, convertToInt)
	}
	fmt.Println(convertToSliceOfInt) // fmt.Println(convertToSliceOfInt) declared outside of for loop; all iterations inside for loop take place then we print result
	for i := range convertToSliceOfInt {
		if i != len(convertToSliceOfInt)-1 && convertToSliceOfInt[i+1] > convertToSliceOfInt[i] { // i != len(convertToSliceOfInt)-1 needs to be specified so that we don't iterate over the last index as convertToSliceOfInt[i+1] would lead to out of range index
			count++
			fmt.Println(count)
		}
	}
}

func ReadFile() []string {
	var sonarInput []string
	sonarFile, err := os.Open("sonarsweep.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sonarFile.Close()
	scanner := bufio.NewScanner(sonarFile)
	for scanner.Scan() {
		sonarInput = append(sonarInput, scanner.Text()) // Without appending to the sonarInput slice of string, we would have no stored data => empty string; so we append
		// fmt.Println(scanner.Text()) https://www.devdungeon.com/content/working-files-go#open_close; "Get data from scan with Bytes() or Text()"
	}
	return sonarInput
}

/*
   convertToString := strings.Join(strArr, " ") // Converting string slice to string
   fmt.Println(convertToString)
   for i := 0; i < len(convertToString); i++ { // Length of string equates to every single digit and space
	fmt.Println(convertToString[i])
   }

   for i := 0; i == strArr[i]; i++ { // Length of string slice equates to each element and does not include space
   fmt.Println(strArr[i])
   }

   convertToInt, _ := strconv.Atoi(convertToString) // Converting string to integer
   fmt.Println(convertToInt)

   for _, v := range strArr { // Ranging over string slice to obtain string value (v)
	   fmt.Println(v)
   }

-------------------------------
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("-42") // Converting string to integer
	s := strconv.Itoa(-42) // Converting integer to string
	fmt.Println(i, s)
}
*/
