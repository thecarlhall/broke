// Package processors contains the default processors for ingesting data from
// an ODS file.
package processors

import (
	"fmt"
	"odf/ods"
	"strconv"
)

func FindProcessor(year int64) Processor {
	if year > 2008 {
		return NewProcessor{}
	} else {
		return OldProcessor{}
	}
}

/********************************************************************************
 * Processor interface
 *******************************************************************************/
// Interface Processor defines the methods for processing a spreadsheet.
type Processor interface {
	Name() string
	Print(ods.Table)
	Process() error
}

/********************************************************************************
 * BaseProcessor struct
 *******************************************************************************/
// Processor base structure for common methods.
type BaseProcessor struct {
}

// Print out the data in a table.
func (p BaseProcessor) Print(table ods.Table) {
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

/********************************************************************************
 * OldProcessor struct
 *******************************************************************************/
// OldProcessor processes rows in a spreadsheet that follow the older format.
type OldProcessor struct {
	BaseProcessor
}

func (p OldProcessor) Name() string {
	return "old"
}

func (p OldProcessor) Process() error {
	/*
        String biller = row.getCell(0).getStringCellValue();
        if ("Leftover".equals(biller)) {
            return;
        }

        Date date = row.getCell(1).getDateCellValue();
        double amount = 0;
        if (row.getCell(2).getCellType() == Cell.CELL_TYPE_NUMERIC) {
            amount = row.getCell(2).getNumericCellValue();
        }

        double credit = 0;
        double debit = 0;
        if (amount < 0 || !"Paycheck".equals(biller)) {
            credit = Math.abs(amount);
        } else {
            debit = Math.abs(amount);
        }
        System.out.println(biller + "; " + date + "; " + debit + "; " + credit);
	*/
	return nil
}

/********************************************************************************
 * NewProcessor struct
 *******************************************************************************/
// NewProcessor processes rows in a spreadsheet that follow the newer format.
type NewProcessor struct {
	BaseProcessor
}

func (p NewProcessor) Name() string {
	return "new"
}

func (p NewProcessor) Process() error {
	/*
        String biller = row.getCell(0).getStringCellValue();
        Date date = row.getCell(1).getDateCellValue();
        double debit = 0;
        if (row.getCell(2).getCellType() == Cell.CELL_TYPE_NUMERIC) {
            debit = row.getCell(2).getNumericCellValue();
        }
        double credit = 0;
        if (row.getCell(3).getCellType() == Cell.CELL_TYPE_NUMERIC) {
            credit = row.getCell(3).getNumericCellValue();
        }
        System.out.println(biller + "; " + date + "; " + debit + "; " + credit);
	*/
	return nil
}
