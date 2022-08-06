package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)
func main() {
	var num int 
	var isPrime bool
	file, err := os.Open("numbers.txt")
    
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		isPrime = true
		num, err = strconv.Atoi(scanner.Text())

		if err != nil {
            panic(err)
        }
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
			}

		}
		if num>=2 && isPrime == true {
			fmt.Println(num, "Number is prime number")
		} else {
			fmt.Println(num,"Number is not prime number")
		}

	}
}
