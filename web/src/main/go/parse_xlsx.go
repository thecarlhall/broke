package broke

import (
  "github.com/tealeg/xlsx"
  "fmt"
)

const (
  xlsxPath = "/home/chall/Downloads/bills.xlsx"
)

func main() {
  var xlFile *xlsx.File
  var error error

  xlFile, error = xlsx.OpenFile(xlsxPath)
  if error != nil {
    fmt.Printf("BOOM!")
  }

  //for sheetIndex := 0; sheetIndex < len(xlFile.Sheets); sheetIndex++ {
  //  var sheet = xlFile.Sheets[sheetIndex]
  for _, sheet := range xlFile.Sheets {
    for _, row := range sheet.Rows {
      for _, cell := range row.Cells {
        if cell.String() != "" {
          fmt.Printf("%s", cell.String())
        }
      }
    }
  }
}
