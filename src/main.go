package main

import (
	"github.com/wppzxc/tasks/src/cmd"
)

func main() {
	command := cmd.NewTasksServerCommand()
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
