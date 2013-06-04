package main

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	//"database/sql"
	//_ "github.com/lib/pq"
	"odf/ods"

	"broke/logger"
)

const (
	ODS_PATH = "/Users/carl/Downloads/bills.ods"
	DEBUG    = true
)

var (
	moneyCleanRegex = regexp.MustCompile(`[$,]`)
	newSkipTerms    = []string{"Allowance", "Auto Insurance", "Auto Tags", "Balance", "Bills", "Income", "Savings"}
	oldSkipTerms    = []string{"Carry over", "Leftover"}
)

func ProcessNewFormat(table ods.Table) {
	for _, row := range table.Strings() {
		if len(row) < 2 || row[0] == "" || shouldSkip(newSkipTerms, row[0]) {
			continue
		}
		if strings.Index(row[0], "Billing Cycle") == 0 {
			logger.Info(row[0])
			continue
		}
		biller := row[0]
		date := row[1]
		if date == "" {
			continue
		}
		var debit, credit float64
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

		if credit != 0 || debit != 0 {
			if DEBUG {
				logger.Debugf("%35s; %11s; %10.2f; %5.2f", biller, date, debit, credit)
			}
		} else {
				logger.Debugf("********** %s [%s] had no values", biller, date)
		}
	}
}

func ProcessOldFormat(table ods.Table) {
	for _, row := range table.Strings() {
		if shouldSkip(oldSkipTerms, row[0]) {
			continue
		}
		if strings.Index(row[0], "Billing Cycle") == 0 {
			logger.Infof("*** %s", row[0])
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
		var credit, debit float64
		if amount < 0 || biller != "Paycheck" {
			credit = math.Abs(amount)
		} else {
			debit = math.Abs(amount)
		}

		if credit != 0 || debit != 0 {
			if DEBUG {
				logger.Debugf("%35s; %11s; %10.2f; %5.2f", biller, date, debit, credit)
			}
		} else {
			if DEBUG {
				logger.Debugf("%s [%s] had no values", biller, date)
			}
		}
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

func printErr(err error) bool {
	if err != nil {
		logger.Error(err)
		return true
	}
	return false
}

func shouldSkip(skips []string, val string) (bool) {
	return sort.SearchStrings(skips, strings.Trim(val, " ")) < 0
}

func main() {
	var doc ods.Doc

	f, err := ods.Open(ODS_PATH)
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer f.Close()

	if err := f.ParseContent(&doc); err != nil {
		logger.Fatal(err)
		return
	}

	// Dump the first table one line per row, writing
	// tab separated, quoted fields.
	for _, table := range doc.Table {
		if year, err := strconv.Atoi(table.Name); err == nil {
			if year >= 2008 {
				logger.Infof("***** Sheet: %s [new] *****", table.Name)
				ProcessNewFormat(table)
			} else {
				logger.Infof("***** Sheet: %s [old] *****", table.Name)
				ProcessOldFormat(table)
			}
		}
	}
}
