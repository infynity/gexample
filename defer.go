package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("defer.txt")

	defer fmt.Println(66666)
	defer fmt.Println(2323)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")



	fmt.Fprintln(f, "data")
	fmt.Fprintln(f, "data2")
	fmt.Fprintln(f, "data3")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	//os.Exit(123)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}