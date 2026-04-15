package main

import (
	"bufio"
	"fmt"
	"strconv"
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

func delete(id int, tasks []Task)[]Task{

	for i, v := range(tasks){
		if v.Id == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			updated, _ := json.MarshalIndent(tasks, "", " ")
			os.WriteFile("tasks.json", updated, 0644)
			return tasks
		}
	}

	fmt.Println("No such tasks with this id. Try again.")
	return tasks
}

func add(desc string, id int, tasks []Task) []Task {

	task := Task{
		Id: id,
		Description: desc,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, task)
	newTask, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile("tasks.json", newTask, 0)

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

	for _ , v:= range tasks{
		if cmd == "all"{
			fmt.Println("ID:",v.Id,"-", v.Description)
			continue
		}
		if v.Status == status{
			fmt.Println(v.Id,"-",v.Description)
		}
	}
}

func update(tasks []Task, id int, desc string)[]Task{
	
	for i, v := range tasks{
		if (v.Id == id){
			tasks[i].Description = desc
			updated, _ := json.MarshalIndent(tasks, "", " ")
			os.WriteFile("tasks.json", updated, 0644)
			return tasks
		}
	}
	
	fmt.Println("No such id for any task. Try again.")
	return tasks
}

func mark(tasks []Task, id int, mark string)[]Task{

	for i, v := range tasks{
		if v.Id == id {
			tasks[i].Status = mark
			updated, _ := json.MarshalIndent(tasks, "", " ")
			os.WriteFile("tasks.json", updated, 0644)
			return tasks
		}
	}
	fmt.Println("No such id for any task. Try again.")
	return tasks
}

func main() {

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
				tasks = add(strings.Trim(task, ""), id, tasks)
			}
		case "update":
			id, description, found := strings.Cut(task, " ")
			idconverted, _:= strconv.Atoi(id)
			if found{
				tasks = update(tasks, idconverted, strings.Trim(description, `"`))
			} else{
				fmt.Println("Update what?! Try again.")
			}
		case "mark-in-progress":
			id, _ := strconv.Atoi(task)
			tasks = mark(tasks, id, "in-progress")
		case "mark-done":
			id, _ := strconv.Atoi(task)
			tasks = mark(tasks, id, "done")
		case "delete":
			id, _ := strconv.Atoi(task)
			tasks = delete(id, tasks)
		case "list":
			fmt.Println("list")
			if !found{
				list("all")
			}else{
				list(task)
			}
		}
	}
}