package operations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rishabdhar12/commands-hive.git/models"
	"github.com/rishabdhar12/commands-hive.git/table"
)

// Add adds a command to the list
func Add(list []models.Command, item models.Command) []models.Command {
	// TODO: Take inputs from user and validate
	list = append(list, item)
	return list
}

// WriteToFile writes the list to a file
func WriteToFile(list []models.Command, item models.Command) {

	file, err := os.ReadFile("commands.json")
	if err == nil {

		err = json.Unmarshal(file, &list)
		if err != nil {
			panic(err)
		}

		list = append(list, item)

		updatedData, err := json.Marshal(list)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("commands.json", updatedData, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("JSON1 data updated successfully!")
	} else {

		_, err := os.Create("commands.json")

		if err != nil {
			panic(err)
		}

		updatedData, err := json.Marshal(list)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("commands.json", updatedData, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("JSON data updated successfully!")
	}
}

// ReadFromFile reads the list from a file
func ReadFromFile() {
	file, err := os.ReadFile("commands.json")
	if err == nil {

		var list []models.Command

		err = json.Unmarshal(file, &list)
		if err != nil {
			panic(err)
		}

		table.PrintTable(list)
	} else {
		fmt.Println("No data found")
	}
}
