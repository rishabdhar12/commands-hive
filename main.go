package main

import (
	"flag"
	"fmt"

	"github.com/rishabdhar12/commands-hive.git/models"
	"github.com/rishabdhar12/commands-hive.git/operations"
)

func main() {
	// flags
	AddFlag := flag.Bool("a", false, "add a command")
	ListFlag := flag.Bool("l", false, "list all commands")

	var list []models.Command

	item := models.Command{
		ID:          "1",
		Name:        "ls",
		Category:    "file",
		Paltform:    "linux",
		Description: "list files",
	}

	flag.Parse()

	if *AddFlag {
		list = operations.Add(list, item)
		fmt.Println(list)
		operations.WriteToFile(list, item)
	} else if *ListFlag {
		operations.ReadFromFile()
	}

}
