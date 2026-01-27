package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)


type Note struct{
	Title string 
	Content string  
	CreatedAt time.Time
}

func (note Note) Display(){
	fmt.Printf("Your Note title as %v and below is content of it: \n\n%v" , note.Title  , note.Content)


}


func (note Note) Save () error{
	fileName :=strings.ReplaceAll(note.Title," ", "_")
	fileName= strings.ToLower(fileName)+".json"

	json, err := json.Marshal(note)
	if(err !=nil){
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func  New(title , content string )(Note , error){

	if(title=="" || content==""){
		return  Note{} , errors.New("Invalid Input ")

	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	},nil

}