package main

import "fmt"


type Reader interface{
	Read() string
}
type Writer interface{
	Write(data string)
}

type ReaderWriter interface{
	Reader
	Writer
}



type file struct{

}

func (f file) Read()string{
	return "Reading file"
}
func (f file) Write(data string){
	fmt.Println("Writing file", data)
}







func main(){
	var rw ReaderWriter = file{}


	fmt.Printf(rw.Read())

	rw.Write("Personal")
}


