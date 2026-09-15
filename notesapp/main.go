// notes taking app that will use a struct for storing some data.it does not only store note data in memory,
// while running the program it also store data in a files in json format .

package main

import (
	"bufio"
	"fmt"
	"os"
	"run/notesapp/note"
	"strings"
)
func main() {
	title, content:=getNoteData()
	userNote,err:=note.New(title,content)
	if err != nil{
fmt.Print(err)
return
	}
userNote.Display()	
err=userNote.Save() 
if err!=nil{
	fmt.Print("file failed to save ")
	return
}
	}

func getNoteData() (string,string){
	title:=UserInput("Note Title ")
	content:=UserInput("note content") 
	return title,content 
}
func UserInput(prompt string ) (string){
fmt.Print(" ",prompt)
 var reader=bufio.NewReader(os.Stdin)
 text,err:=reader.ReadString('\n')
 if err!=nil{
	return ""
 }
 text=strings.TrimSuffix(text,"\n")

return text
}

