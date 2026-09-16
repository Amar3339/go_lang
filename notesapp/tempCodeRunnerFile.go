func getNoteData() (string,string){
	title:=UserInput("Note Title",)
	content:=UserInput("note content") 
	return title,content 
}