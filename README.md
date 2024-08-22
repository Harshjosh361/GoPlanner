# GoPlanner

**Go Planner** is a simple command-line application for managing tasks. This tool allows you to add, update, delete, and list tasks from the command line. It's a great way to practice working with the filesystem, handling user inputs, and building CLI applications in Go.

## Features

- **Add Tasks**: Add new tasks with a title.
- **Update Tasks**: Update the status of existing tasks (e.g., to "in-progress" or "done").
- **Delete Tasks**: Remove tasks by their ID.
- **List Tasks**:
  - All tasks
  - Tasks that are done
  - Tasks that are not done
  - Tasks that are in progress

## Usage

1.**Add a Task**:
  ```bash
   ./myapp -add "Task Title"
```
2.**Update Task**:
```bash
./myapp -update TASK_ID -status "in-progress"
```
3.**Delete task**
```bash
./myapp -delete TASK_ID
```
4.**List all task**
```bash
./myapp -listall
```
5.**List in-Progress task**
```bash
./myapp -inprogress
```
6.**List todo task**
```bash
./myapp -todo
```
https://roadmap.sh/projects/task-tracker





