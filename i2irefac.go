package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main(){

    CheckNumbers("rakamlar.txt")
    

           }

func isPrime(num int) {
    

    var isPrime bool
    isPrime = true

    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
        }
    }
    
        if isPrime == true {
            fmt.Println(num, "Number is prime number")

        } else {
            fmt.Println(num,"Number is not prime number")
    }
}
func CheckNumbers(location string){

    file, err := os.Open("numbers.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        lineStr:=scanner.Text()
        num,_:=strconv.Atoi(lineStr)
    
        isPrime(num)
     }
}