package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func readFile(filename string) error {
	file , err := os.Open(filename)
	if err !=nil{
		return err
	}
	defer file.Close()
	data, err:= io.ReadAll(file)

	if err !=nil{
		return err
	}
	

	fmt.Println(string(data))
	return nil

}

func main() {
	err := readFile("output.txt")
	if err != nil {
		fmt.Println(err)
	}


	data := []int {1,2,3,4,5}
	processData(data)
}



func safeOperation (){


	defer func(){
		if r:=recover(); r !=nil{
			fmt.Println("Recovered from panic",r )
		}

	}()
	panic("Something went wrong")
	fmt.Println("Can't reach here ")
}

func processData(data []int ) {
	start :=time.Now()

	defer func(){
		fmt.Println("Data processing completed in ",time.Since(start))


	}()
	for _ , d :=range data{
		fmt.Println(d)
		time.Sleep(time.Millisecond *100)
	}

	safeOperation()

}