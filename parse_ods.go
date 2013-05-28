package main

import (
	"fmt"
	"log"
	"math"
	"odf/ods"
	"os"
	"strconv"
	"strings"
)

const (
	odsPath = "/Users/carl/Downloads/bills.ods"
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
		log.Fatal(err)
		//fmt.Fprintln(os.Stderr, err)
		return
	}

	// Dump the first table one line per row, writing
	// tab separated, quoted fields.
	for _, table := range doc.Table {
		year, err := strconv.Atoi(table.Name)

		if err != nil {
			continue
		}

		//Print(table)
		if year > 2008 {
			println(fmt.Sprintf("\n***** Sheet: %s [new] *****", table.Name))
			ProcessNewFormat(table)
		} else {
			println(fmt.Sprintf("\n***** Sheet: %s [old] *****", table.Name))
			ProcessOldFormat(table)
		}
	}
}

// Print out the data in a table.
func Print(table ods.Table) {
	for _, row := range table.Strings() {
		if len(row) == 0 || row[0] == "" {
			continue
		}
		sep := ""
		for _, field := range row {
			fmt.Print(sep, strconv.Quote(field))
			sep = "\t"
		}
		fmt.Print("\n")
	}
}

func ProcessNewFormat(table ods.Table) {
	for _, row := range table.Strings() {
		if len(row) < 2 || row[0] == "" || strings.Index(row[0], "Balance") == 0 {
			continue
		}
		if strings.Index(row[0], "Billing Cycle") == 0 {
			println(row[0])
			continue
		}
		biller := row[0]
		date := row[1]
		if date == "" {
			continue
		}
		debit := 0.0
		credit := 0.0
		if len(row) >= 3 {
			debit, _ = strconv.ParseFloat(row[2], 32)
		}
		if len(row) >= 4 {
			debit, _ = strconv.ParseFloat(row[3], 32)
		}
		fmt.Printf("%35s; %11s; %10.2f; %5.2f\n", biller, date, debit, credit)
	}
}

func ProcessOldFormat(table ods.Table) {
	for _, row := range table.Strings() {
		//println(strings.Index(row[0], "Billing Cycle"))
		if len(row) > 2 || row[0] == "" || row[0] == "Leftover" {
			continue
		}
		biller := row[0]
		date := row[1]
		if date == "" {
			continue
		}
		var amount float64
		if len(row) >= 3 {
			var err error
			amount, err = strconv.ParseFloat(strings.Replace(row[2], "$", "", 1), 32)
			if err != nil {
				log.Fatal(err)
				continue
			}
		}
		credit := 0.0
		debit := 0.0
		if amount < 0 || biller != "Paycheck" {
			credit = math.Abs(amount)
		} else {
			debit = math.Abs(amount)
		}

		fmt.Printf("%35s; %11s; %10.2f; %5.2f\n", biller, date, debit, credit)
	}
}
