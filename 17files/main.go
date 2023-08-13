package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "this needs to go in file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	// dataByte, err := ioutil.ReadFile(fileName)
	dataByte, err := os.ReadFile(fileName)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("text data inside the file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
