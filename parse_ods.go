package main

import (
	"fmt"
	"odf/ods"
	"os"
	"broke/processors"
	"strconv"
)

const (
	odsPath = "/home/chall/Downloads/bills.ods"
)

func main() {
	var doc ods.Doc

	f, err := ods.Open(odsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()

	if err := f.ParseContent(&doc); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Dump the first table one line per row, writing
	// tab separated, quoted fields.
	for _, table := range doc.Table {
		year, err := strconv.ParseInt(table.Name, 0, 0)
		if err != nil {
			continue
		}

		var processor processors.Processor = processors.FindProcessor(year)
		println(fmt.Sprintf("\n***** Sheet: %s [%s] *****", table.Name, processor.Name()))
		processor.Print(table)
	}
}
