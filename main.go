package main

import(
	"fmt"
	//"strings"
)

func main(){
	
	var command string

	fmt.Scanln(&command)

	switch command {
	case "add":
		fmt.Println("add")
	case "update":
		fmt.Println("update")
	case "delete":
		fmt.Println("delete")
	}
}