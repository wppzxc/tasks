package main

import (
	"github.com/wppzxc/tasks/src/cmd"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	command := cmd.NewTasksServerCommand()
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
