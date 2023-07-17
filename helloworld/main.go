// package main is used to create an executable file when we run go build <fileName>
// if we don't use package main and run go build <file name> this wont create an executable file
package main

// importing the code from 'fmt' package (fmt=> formatt)

import "fmt"

func main(){
	fmt.Println("Hello World")
}