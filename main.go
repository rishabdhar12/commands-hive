package main

import (
	"flag"

	"github.com/rishabdhar12/commands-hive.git/models"
	"github.com/rishabdhar12/commands-hive.git/operations"
)

func main() {
	// flags
	AddFlag := flag.Bool("a", false, "add a command")
	ListFlag := flag.Bool("l", false, "list all commands")
	DeleteFlag := flag.Bool("d", false, "delete a command by id")
	SearchFlag := flag.Bool("s", false, "search a command by id")
	HelpFlag := flag.Bool("h", false, "help")
	VersionFlag := flag.Bool("v", false, "print version")

	var list []models.Command

	flag.Parse()

	if *AddFlag {
		operations.Add(list)
	} else if *ListFlag {
		operations.ReadFromFile()
	} else if *DeleteFlag {
		operations.DeleteCommand()
	} else if *SearchFlag {
		operations.SearchById()
	} else if *HelpFlag {
		operations.PrintHelp()
	} else if *VersionFlag {
		operations.PrintVersion()
	} else {
		operations.PrintHelp()
	}

}

