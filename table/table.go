package table

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/rishabdhar12/commands-hive.git/models"
)

func PrintTable(list []models.Command) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{

			{Text: "ID"},
			{Text: "Category"},
			{Text: "Name"},
			{Text: "Paltform"},
			{Text: "Description"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, row := range list {
		id := row.ID
		category := row.Category
		name := row.Name
		paltform := row.Paltform
		description := row.Description

		cells = append(cells, []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: id},
			{Text: category},
			{Align: simpletable.AlignRight, Text: name},
			{Align: simpletable.AlignRight, Text: paltform},
			{Align: simpletable.AlignRight, Text: description},
		})

		table.Body = &simpletable.Body{Cells: cells}
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
