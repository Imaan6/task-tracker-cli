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
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func marshallAndWrite(tasks []Task) {
	updated, err := json.MarshalIndent(tasks, "", " ")
	if err != nil{
		fmt.Println("error during marshall indent of json:", err)
	}
	if err := os.WriteFile("tasks.json", updated, 0644); err != nil{
		fmt.Println("Error occured during WriteFile:", err)
	}
}

//function used to delete tasks
func delete(id int, tasks []Task)[]Task{

	for i, v := range(tasks){
		if v.Id == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			marshallAndWrite(tasks)
			return tasks
		}
	}

	fmt.Println("No such tasks with this id. Try again.")
	return tasks
}

//function used to add tasks
func add(desc string, id int, tasks []Task) []Task {

	task := Task{
		Id: id,
		Description: desc,
		Status: "todo",
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, task)

	marshallAndWrite(tasks)
	
	return tasks
}

//function used to list tasks
func list(cmd string){
	var tasks []Task
	var status string 

	readAndUnmarshall(&tasks)

	switch cmd{
	case "done":
		status = "done"
	case "todo":
		status = "todo"
	case "in-progress":
		status = "in-progress"
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

//function used to update tasks
func update(tasks []Task, id int, desc string)[]Task{
	
	for i, v := range tasks{
		if (v.Id == id){
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
			marshallAndWrite(tasks)
			return tasks
		}
	}
	
	fmt.Println("No such id for any task. Try again.")
	return tasks
}

//function used to assign a mark to tasks
func mark(tasks []Task, id int, mark string)[]Task{
	for i, v := range tasks{
		if v.Id == id {
			tasks[i].Status = mark
			marshallAndWrite(tasks)
			return tasks
		}
	}
	fmt.Println("No such id for any task. Try again.")
	return tasks
}

func readAndUnmarshall(tasks *[]Task){
	data, err := os.ReadFile("tasks.json")
	if err != nil{
		fmt.Println("Error encountered when trying to read from file:", err)
	}

	if err := json.Unmarshal(data, &tasks); err != nil{
		fmt.Println("Error while unmarshalling:", err)
	}
}

func main() {

	var tasks []Task
	
	scanner := bufio.NewScanner(os.Stdin)
	
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err){
		marshallAndWrite(tasks)
	}
	
	id := 0
	
	for scanner.Scan() {
		readAndUnmarshall(&tasks)

		length := len(tasks)
		if length != 0{
			id = tasks[length - 1].Id
		}

		command := scanner.Text()
		cmd, task, found := strings.Cut(command, " ")

		switch cmd {
		case "add":
			id++;
			if found{
				tasks = add(strings.Trim(task, `"`), id, tasks)
			}
		case "update":
			id, description, found := strings.Cut(task, " ")
			idconverted, err := strconv.Atoi(id)
			if err != nil{
				fmt.Println("Error encountered during atoi conversion:", err)
			}
			if found {
				tasks = update(tasks, idconverted, strings.Trim(description, `"`))
			} else {
				fmt.Println("Update what?! Try again.")
			}
		case "mark-in-progress":
			id, err := strconv.Atoi(task)
			if err != nil{
				fmt.Println("Error encountered during atoi conversion:", err)
			}
			tasks = mark(tasks, id, "in-progress")
		case "mark-done":
			id, err := strconv.Atoi(task)
			if err != nil{
				fmt.Println("Error encountered during atoi conversion:", err)
			}
			tasks = mark(tasks, id, "done")
		case "delete":
			id, err := strconv.Atoi(task)
			if err != nil{
				fmt.Println("Error encountered during atoi conversion:", err)
			}
			tasks = delete(id, tasks)
		case "list":
			if !found{
				list("all")
			}else{
				list(task)
			}
		}
	}
}