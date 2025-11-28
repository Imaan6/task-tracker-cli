package main

import "time"

/*
lower case fields = private fields that can only be accessed from inside the package
upper case struct fields ( like the task example) = public fields that can be accessed from anywhere
*/

type task struct{
	Id int 
	Description string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}
