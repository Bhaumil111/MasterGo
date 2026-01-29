package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
	
)

type Saver interface{
	Save() error
}

type Displayer interface{
	Display() error
}

type OutputTable interface{
	Saver
	Display()
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text

}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")
	content := getUserInput("Note Content:")
	return title, content

}
func outputData(data OutputTable)error{
	data.Display()
	return saveData(data)
}

func saveData (data Saver) error {
	err :=data.Save()

	if(err!=nil){
		fmt.Println("Saving data failed")

	}
	fmt.Println("Data saved succesfully")
	return nil
}

func printSomething( value interface{}){

	intVal, ok := value.(int)
	if(ok){
		fmt.Println("Type of value is ", intVal)

	}
	floatVal , ok :=value.(float64)

	if(ok){
		fmt.Println("Type of this value is ", floatVal)
	}

}


func add(a, b interface{})interface{}{
	aint, aisint := a.(int)
	bint, bisint :=b.(int)

	if(aisint&& bisint){
		return aint+bint
	}

	return nil
}


func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return

	}
	outputData(userNote)
	outputData(todo)
	
}
