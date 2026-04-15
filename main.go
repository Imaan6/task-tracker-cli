package main

import (
	"bufio"
	"fmt"
	//"strconv"
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

func add(desc string, id int, tasks []Task) []Task {
	fmt.Println("the descriptions is:", desc)
	fmt.Println("the id is", id)
	

	task := Task{
		Id: id,
		Description: desc,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, _ := os.ReadFile("tasks.json")

	if len(tasks) > 0{
		json.Unmarshal(data, &tasks)
	}

	tasks = append(tasks, task)
	return tasks
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

func update(tasks []Task, id int, desc string)[]Task{
	for i, v := range tasks{
		if (v.Id == id){
			tasks[i].Description = desc
		}
	}
	return tasks
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

	var tasks []Task

	id := 0

	scanner := bufio.NewScanner(os.Stdin)
	
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	defer file.Close()

	for scanner.Scan() {
		command := scanner.Text()
		cmd, task, found := strings.Cut(command, " ")
		switch cmd {
		case "add":
			id++;
			if found{
				//TODO: see if i can increment inside the id
				tasks = add(strings.Trim(task, ""), id, tasks)
			}
			newTask, _ := json.MarshalIndent(tasks, "", " ")
			//TODO: everytime a write a new task the whole thing get added to the json file instead of just the actual task, fix.
			os.WriteFile("tasks.json", newTask, 0)
		case "update":
			fmt.Println("update")
			//idconv, _:= strconv.Atoi(cmd[1])
			//tasks = update(tasks, idconv, strings.Trim(cmd[2], `"`))
		case "delete":
			fmt.Println("delete")
		case "list":
			fmt.Println("list")
			/*if len(cmd) ==2{
				list(cmd[1])
			}else{
				list("")
			}	*/		
		}
	}
}