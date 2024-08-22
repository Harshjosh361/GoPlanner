package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func AddTask(tasks []Task, title string) []Task {
	task := Task{
		Id:        len(tasks) + 1,
		Title:     title,
		Status:    "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tasks = append(tasks, task)
	return tasks
}

func UpdateTask(tasks []Task, id int, title string, status string) ([]Task, error) {
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Title = title
			tasks[i].Status = status
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("Task not updated")
}

func DeleteTask(tasks []Task, id int) []Task {
	var newTask []Task
	for i := range tasks {
		if tasks[i].Id == id {
			newTask = append(tasks[:i], tasks[i+1:]...)
			return newTask
		}
	}
	return newTask
}

func ListTask(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("%d: %s [%s]\n", task.Id, task.Title, task.Status)
	}
}

func CompletedTask(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "done" {
			fmt.Printf("%d: %s [%s]\n", task.Id, task.Title, task.Status)
		}
	}
}

func InprogressTask(tasks []Task) {

	for _, task := range tasks {
		if task.Status == "in progress" {
			fmt.Printf("%d: %s [%s]\n", task.Id, task.Title, task.Status)
		}
	}
}

func TodoTask(tasks []Task) {
	for _, task := range tasks {
		if task.Status == "todo" {
			fmt.Printf("%d: %s [%s]\n", task.Id, task.Title, task.Status)
		}
	}
}
