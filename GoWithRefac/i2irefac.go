package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"

	"os"
	"strconv"
)
var stringNum string
func main(){

    CheckNumbers("numbers.txt")
    ReadFile()
    
        

           }

func isPrime(num int,numstr string) {
   
    var isPrime bool
    isPrime = true
    

    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
        }
    }
    
        if num>=2 && isPrime == true {
            stringNum= stringNum+ numstr +"is prime"

        } else {
            stringNum= stringNum+ numstr +"is not prime"
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
