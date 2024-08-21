package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := "tasks.json"

	// 1.name 2.defaultvalue 3.description
	addCmd := flag.String("add", "", "Add a new task")
	updateId := flag.Int("updateId", 0, "Id of the task to update")
	updateTitle := flag.String("updateString", "", "Title to be updated")
	updateStatus := flag.String("updateStatus", "", "Status to be updated")
	deleteCmd := flag.Int("delete", 0, "Delete a task")
	listCmd := flag.Bool("listall", false, "List all tasks")
	inProgress := flag.Bool("inprogress", false, "List in progress tasks")
	todoTask := flag.Bool("todotask", false, "List all todo tasks")

	// Parsing the flags
	flag.Parse()

	// loading task
	tasks, err := LoadTask(filename)
	if err != nil {
		fmt.Println("Error loading tasks")
		return
	}

	switch {
	case *addCmd != "":
		tasks = AddTask(tasks, *addCmd)

	case *deleteCmd != 0:
		tasks = DeleteTask(tasks, *deleteCmd)

	case *listCmd:
		ListTask(tasks)

	case *inProgress:
		InprogressTask(tasks)

	case *todoTask:
		TodoTask(tasks)

	case *updateId != 0:
		if *updateTitle != "" && *updateStatus != "" {
			updatedTask, err := UpdateTask(tasks, *updateId, *updateTitle, *updateStatus)
			if err != nil {
				fmt.Println("Failed to Update Task")
			}
			fmt.Println("Successfully Update Task")
			fmt.Println(updatedTask)
		} else {
			fmt.Println("Add title and status")
		}

	default:
		fmt.Println("No valid command provided")
	}

	// saving tasks
	err = SaveTask(filename, tasks)
	if err != nil {
		fmt.Println("Error saving tasks", err)
	}
	fmt.Println("Task saved successfully")

}
