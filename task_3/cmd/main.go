package main

import (
	"os"
	app "task_2/internal/app"
)

func main() {
	source := "file"
	if len(os.Args) > 1 {
		source = os.Args[1]
	}

	app := &app.App{}
	app.Run(source)
}