package operations

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/rishabdhar12/commands-hive.git/models"
	"github.com/rishabdhar12/commands-hive.git/table"
)

// Add adds a command to the list
func Add(list []models.Command) []models.Command {

	time := time.Now().UnixMilli()

	timeString := strconv.Itoa(int(time))

	id := timeString[len(timeString)-4:]

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the command : ")
	scanner.Scan()
	iName := scanner.Text()

	fmt.Printf("Enter the category : ")
	scanner.Scan()
	iCategory := scanner.Text()

	fmt.Printf("Enter the platform : ")
	scanner.Scan()
	iPlatform := scanner.Text()

	fmt.Printf("Enter the description : ")
	scanner.Scan()
	iDescription := scanner.Text()

	validate := ValidateInput(iName, iCategory, iPlatform, iDescription)

	if !validate {
		panic("Validation error")
	} else {

		item := models.Command{
			ID:          id,
			Name:        iName,
			Category:    iCategory,
			Platform:    iPlatform,
			Description: iDescription,
		}

		list = append(list, item)
		WriteToFile(list, item)
	}
	return list
}

// validate input
func ValidateInput(iCommand string, iCategory string, iPlatform string, iDescription string) bool {
	if iCommand == "" || iCategory == "" || iPlatform == "" || iDescription == "" {
		return false
	} else {
		return true
	}
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

// Delete deletes a command from the list
func DeleteCommand() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the ID of the command you want to delete : ")
	scanner.Scan()
	iID := scanner.Text()

	file, err := os.ReadFile("commands.json")
	if err == nil {

		var list []models.Command

		err = json.Unmarshal(file, &list)
		if err != nil {
			fmt.Println("File is empty")
			panic(err)
		}

		for i, item := range list {
			if item.ID == iID {
				list = append(list[:i], list[i+1:]...)
				break
			} else {
				fmt.Println("Command not found")
			}
		}

		updatedData, err := json.Marshal(list)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("commands.json", updatedData, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("File not found")
	}
}

// Search command by id
func SearchById() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the ID of the command you want to search : ")
	scanner.Scan()
	iID := scanner.Text()

	file, err := os.ReadFile("commands.json")
	if err == nil {

		var list []models.Command
		var list1 []models.Command

		err = json.Unmarshal(file, &list)
		if err != nil {
			fmt.Println("File is empty")
			panic(err)
		}

		for _, item := range list {
			if item.ID == iID {
				foundItem := item
				list1 = append(list1, foundItem)
				table.PrintTable(list1)
				break
			} else {
				fmt.Println("Command not found")
			}
		}
	} else {
		fmt.Println("File not found")
	}
}

// Print Help
func PrintHelp() {
	fmt.Println("Commands Hive is a CLI tool to manage your commands")
	fmt.Println("Usage: commands-hive -[arguments]")
	fmt.Println("Commands:")
	fmt.Println("a\t\t\t\tAdd a new command")
	fmt.Println("d\t\t\t\tDelete a command")
	fmt.Println("s\t\t\t\tSearch a command")
	fmt.Println("h\t\t\t\tPrint help")
	fmt.Println("v\t\t\t\tPrint version")
}

// Print Version
func PrintVersion() {
	fmt.Println("Version 1.0.0")
}
