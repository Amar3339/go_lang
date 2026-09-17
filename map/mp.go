package main

import (
	"fmt"


)
func main(){
	url:=map[string]string{
		"google":"https//google.com",
		"amazon_web_service":"https//aws.com",
		
	}
	fmt.Println(url["google"])
	 url["yahoo"]="https//yahoo.com" //you can mutate the map means add some more url 
	 fmt.Println(url)

	 delete(url,"yahoo")
	 	 fmt.Println(url)  //you can also delete some url 
		//  like that you can also update in map 
}

