package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strings"
	// "strconv"
)

func main() {
	strArr := ReadFile()
	// fmt.Println(strArr)
	for _, v := range strArr { // Ranging over string slice to obtain string value (v)
		convertToInt, _ := strconv.Atoi(v) // Converting string to integer
		// fmt.Println(convertToInt)
		var convertToSliceOfInt []int = []int{convertToInt}
		fmt.Println(convertToSliceOfInt)
		convertToSliceOfInt = append(convertToSliceOfInt, convertToInt)
		fmt.Println(convertToInt)
		// 	count := 0
		// 	for i := 0; i < len(convertToSliceOfInt); i++ {
		// 		if convertToSliceOfInt[i+1] > convertToSliceOfInt[i] {
		// 			count++
		// 	}
		// 	fmt.Print(count)
		// }
		// if convertToSliceOfInt[1] > convertToSliceOfInt[0] {
		// 	count++
		// }
	}

	//    convertToString := strings.Join(strArr, " ") // Converting string slice to string
	//    fmt.Println(convertToString)
	//    nums := []int{convertToInt}
	//    fmt.Println(nums)
	//    for i := 0; i == convertToInt; i++ {
	// 	fmt.Println(strArr[i], strArr[i+1])
	//    }
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
func count() {
    var strArr []string
    var line string
    count := 0
    for i := 0; i < len(strArr); i++ {
        if strArr[i+1] > strArr[i] {
            count++
    }
    fmt.Print(line)
}
}
*/

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
*/
