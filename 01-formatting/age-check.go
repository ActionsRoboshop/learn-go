package main

import "fmt"

// program to validate whether user to eligible  for vote.

func main() {

var age int
fmt.Printf("Enter a age to check:")
fmt.Scan(&age)

if age >=18 {
fmt.Printf("User is eligible to VOTE")
} else {
fmt.Printf("User is NOT eligible to VOTE")
}
}



