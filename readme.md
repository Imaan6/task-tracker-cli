Task tracker is a simple cli to track your tasks:

You *run* the program by: ```go run main.go```

To *add* a task: ``` add task1 ```

To *Update* a task ```update 1 "Go to the gym" ```

To *Delete* a task ```delete 1```

To *mark* as *in-progress* or as *done* : ```mark-in-progess 1``` || ```mark-done```

To *list* all tasks: ```list```

To *list* all tasks that are *done*: ```list done```
To *list* all tasks that are *not done*: ```list todo```
to *list* all tasks that are *in progress*: ```list in-progress```

N.B: The tasks are stored in a json file in the same repo called tasks.json
