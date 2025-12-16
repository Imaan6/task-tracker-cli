package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"encoding/json"
	"os"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func add(desc string, id int, tasks []Task){
	fmt.Println("the descriptions is:", desc)
	fmt.Println("the id is", id)
	
}

func list(cmd string){
	list, err := os.ReadFile("tasks.json")
	if err != nil{
		return 
	}
	
	var tasks []Task
	var status string 

	switch cmd{
	case "done":
		status = "done"
	case "todo":
		status = "todo"
	case "in-progress":
		status = "in-progress"
	}
	
	err = json.Unmarshal(list, &tasks)
	if err != nil{
		return 
	}

	if cmd == ""{
		for _, v := range tasks{
			fmt.Println(v.Id,"-", v.Description)
		}
	}else{
		for _ , v:= range tasks{
			if v.Status == status{
				fmt.Println(v.Id,"-",v.Description)
			}
		}
	}
}
/*
Add, Update, and Delete tasks

Mark a task as in progress or done

List all tasks
List all tasks that are done
List all tasks that are not done
List all tasks that are in progress

task-cli list done
task-cli list todo
task-cli list in-progress
*/

func main() {
	/*
		Marshall
		verb
		1.assemble and arrange (a group of people, especially troops) in order.
	*/

	
	task1 := Task{
		Id:          1,
		Description: "List all tasks",
		Status:      "done",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task2 := Task{
		Id:          2,
		Description: "List all tasks that are done",
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task3 := Task{
		Id:          3,
		Description: "List all tasks that are not done",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task4 := Task{
		Id:          3,
		Description: "list all tasks that are in progress",
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	var tasks []Task
	tasks = append(tasks, task1, task2, task3, task4)
	newTask, _ := json.Marshal(tasks)
	//////////////////////////////////////

	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	defer file.Close()

	file.WriteString(string(newTask))

	id := 0

	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		command := scanner.Text()
		cmd := strings.SplitN(command, " ", 2)

		switch cmd[0] {
		case "add":
				add(strings.Trim(cmd[1], `"`), id+1, tasks)
		case "update":
			fmt.Println("update")
		case "delete":
			fmt.Println("delete")
		case "list":
			if len(cmd) ==2{
				list(cmd[1])
			}else{
				list("")
			}			
		}
	}
}
