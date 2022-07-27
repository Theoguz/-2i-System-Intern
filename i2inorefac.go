package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)
func main() {
	var primNum int 
	var isPrime bool

	
	file, err := os.Open("rakamlar.txt")
    if err != nil {
        panic(err)
    }
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		isPrime = true
		primNum, err = strconv.Atoi(scanner.Text())
		if err != nil {
            panic(err)
        }
		if primNum < 2 {
			fmt.Println("ot prime")
		} else {
			for i := 2; i <= primNum/2; i++ {
				if primNum%i == 0 {
					isPrime = false
				}
			}
		
		}
		if isPrime == true {
			fmt.Println(primNum, "Number is prime number")
		} else {
			fmt.Println(primNum,"Number is not prime number")
		}
	}
}
