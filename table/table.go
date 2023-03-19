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
			{Text: "Name"},
			{Text: "Category"},
			{Text: "Platform"},
			{Text: "Description"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, row := range list {
		id := row.ID
		name := row.Name
		category := row.Category
		platform := row.Platform
		description := row.Description

		cells = append(cells, []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: id},
			{Align: simpletable.AlignLeft, Text: name},
			{Align: simpletable.AlignLeft, Text: category},
			{Align: simpletable.AlignLeft, Text: platform},
			{Align: simpletable.AlignLeft, Text: description},
		})

		table.Body = &simpletable.Body{Cells: cells}
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
