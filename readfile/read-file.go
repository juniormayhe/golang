/*
Run with: 
go run read-file.go sample.txt
*/
package main

import (
	"fmt"
	"io"
	"os"
	//"reflect"
)

/*
$ go run read-file.go johanna sanchez
Reading file...
arg: 0 C:\Users\Junior\AppData\Local\Temp\go-build883084518\b001\exe\read-file.exe
arg: 1 johanna
arg: 2 sanchez

$ go run read-file.go "johanna sanchez"
Reading file...
arg: 0 C:\Users\Junior\AppData\Local\Temp\go-build639082746\b001\exe\read-file.exe
arg: 1 johanna sanchez
*/
func main() {
	fmt.Println("Reading file...")

	argsWithProg := os.Args

	fmt.Println("//arg contents:")
	for i, arg := range argsWithProg {
		fmt.Println("arg:", i, arg)
	}
	fmt.Println("//arg 1 content:")

	// slice of strings
	//filename := os.Args[1:]

	// a single string
	filename := os.Args[1]

	//fmt.Println("filename \"", filename, "\" of type ", reflect.TypeOf(filename))
	fmt.Printf("filename \"%v\" of type %T \n", filename, filename)

	//f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\n//file contents:\n")
	// copy file content to  console output
	io.Copy(os.Stdout, f)

	fmt.Printf("\n\n//done")
}
