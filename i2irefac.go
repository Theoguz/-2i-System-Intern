package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
func main() {
	opentxt()
	
	
}
func opentxt(){
	file, err := os.Open("numbers.txt")
	if err != nil {
			log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num,_:=strconv.Atoi(lineStr)
		isPrime(num)
	}
	
}
func isPrime(num int){
	var isPrime bool
	isPrime=true

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
		}
			
	}
	if num>=2 && isPrime == true {
		fmt.Println(num,"Number is prime number")
	} else {
		fmt.Println(num,"Number is not prime number")
	}

}

	

