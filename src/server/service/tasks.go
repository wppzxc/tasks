package service

import "github.com/wppzxc/tasks/src/database"

func GetCurrentTasks() (*database.Tasks, error) {
	task, err := database.GetCurrentTasks()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetTaskByID(taskID string) (*database.Tasks, error) {
	task, err := database.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func CreateTask(task *database.Tasks) (*database.Tasks, error) {
	newTask, err := database.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}

func UpdateTask(task *database.Tasks) (*database.Tasks, error) {
	newTask, err := database.UpdateTask(task)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}
