package main

import (
	"fmt"
	"log"
	"math"
	"odf/ods"
	"regexp"
	"strconv"
	"strings"
)

const (
	odsPath = "/Users/carl/Downloads/bills.ods"
)

var (
	moneyCleanRegex = regexp.MustCompile(`[$,]`)
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
			log.Printf("*** %s\n", row[0])
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
			var err error
			debit, err = parseMoney(row[2])
			printErr(err)
		}
		if len(row) >= 4 {
			var err error
			credit, err = parseMoney(row[3])
			printErr(err)
		}
		log.Printf("%35s; %11s; %10.2f; %5.2f\n", biller, date, debit, credit)
	}
}

func ProcessOldFormat(table ods.Table) {
	for _, row := range table.Strings() {
		if row[0] == "Leftover" {
			continue
		}
		if strings.Index(row[0], "Billing Cycle") == 0 {
			log.Printf("*** %s\n", row[0])
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
			if printErr(err) {
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

		log.Printf("%35s; %11s; %10.2f; %5.2f\n", biller, date, debit, credit)
	}
}

// Parse a string into a float. This removes non-numeric characters.
func parseMoney(data string) (float64, error) {
	var amount float64
	var err error
	if len(data) > 0 {
		amount, err = strconv.ParseFloat(moneyCleanRegex.ReplaceAllString(data, ""), 32)
	}
	return amount, err
}

func printErr(err error) (bool) {
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return true
	}
	return false
}

func main() {
	var doc ods.Doc

	f, err := ods.Open(odsPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	if err := f.ParseContent(&doc); err != nil {
		log.Fatal(err)
		return
	}

	// Dump the first table one line per row, writing
	// tab separated, quoted fields.
	for _, table := range doc.Table {
		if year, err := strconv.Atoi(table.Name); err == nil {
			if year > 2008 {
				log.Printf("***** Sheet: %s [new] *****", table.Name)
					ProcessNewFormat(table)
			} else {
				log.Printf("***** Sheet: %s [old] *****", table.Name)
					ProcessOldFormat(table)
			}
		}
	}
}
