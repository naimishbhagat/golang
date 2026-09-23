package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)



type saver interface {
	Save() error
}

type displayer interface {
	Display() 
}
type outputtable interface{
	saver
	displayer
}

func main(){
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	title,content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.NewTodo(todoText)
	userNote, err := note.NewNote(title,content)
	err = outputData(todo)
	if err != nil {
		return
	}
	outputData(userNote)
	fmt.Println("Saving the note succeeded!")
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if !ok {
		fmt.Println("Integer: ", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if !ok {
		fmt.Println("Float: ", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if !ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type){
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float:",value)
	// case string:
	// 	fmt.Println(value)	
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error{
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string){
	title := getUserInput("Note title:")
	content := getUserInput("Note Content:")
	
	return title, content
}

func getUserInput(prompt string) (string){
	fmt.Printf("%v",prompt)
	// var value string
	// fmt.Scan((&value))
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}