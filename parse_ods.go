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
			fmt.Printf("\n*** %s\n", row[0])
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
		println(row[2])
			debit, _ = parseMoney(row[2])
		}
		if len(row) >= 4 {
			credit, _ = parseMoney(row[3])
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
		var err error
		if len(row) >= 3 {
			amount, err = parseMoney(row[2])
			println(amount)
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

func parseMoney(data string) (float64, error) {
	amount, err := strconv.ParseFloat(strings.Replace(data, "$", "", 1), 32)
	return amount, err
}

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
	//for _, table := range doc.Table {
		table := doc.Table[0]
		year, err := strconv.Atoi(table.Name)

		//if err != nil {
		//	continue
		//}

		//Print(table)
		if year > 2008 {
			println(fmt.Sprintf("\n***** Sheet: %s [new] *****", table.Name))
			ProcessNewFormat(table)
		} else {
			println(fmt.Sprintf("\n***** Sheet: %s [old] *****", table.Name))
			ProcessOldFormat(table)
		}
	//}
}
