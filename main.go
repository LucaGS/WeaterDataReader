package main

import (
	Month_Package "ExslReaderv2/Month"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	//Einlesen von xlsx Datei
	AllMonths := make([]Month_Package.Month, 13)
	totalCount := 0.0
	rowCount := 0.0
	f, err := excelize.OpenFile("dataexport_20240118T091822.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	for index1, row := range rows {
		if index1 <= 10 {
			continue
		}
		tempString := ""
		tempFloat := 0.0
		dateString := ""
		count++
		for index, value := range row {
			if index == 0 {
				dateString = value
			}
			if index == 1 {
				tempString = row[index]
				fmt.Printf("temp String: %v ", tempString)
				tempFloat, err = strconv.ParseFloat(tempString, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				totalCount += tempFloat

				rowCount = float64(index1)
			}
			parsedTime, err := time.Parse("2006-01-02T15:04:05", dateString)
			if int(parsedTime.Year()) == 2024 {
				break
			}
			currentMonth := parsedTime.Month()
			currentDay := int(parsedTime.Day())
			currentHour := int(parsedTime.Hour())
			//Wenn noch kein Name vorhanden ist Dann erstelle neuen Monat noch testen
			if len(AllMonths[int(currentMonth)-1].Name) == 0 {
				fmt.Printf("Es ist ein neuer Monat: %v\n", currentMonth)
				AllMonths[int(currentMonth)-1] = *Month_Package.NewMonth(currentMonth.String())
			}

			AllMonths[int(currentMonth)-1].AllTemps[currentDay-1][currentHour] = tempFloat
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	fmt.Printf("%v", rowCount/12)
	for i := 0; i < 12; i++ {
		AllMonths[i].CalcMonthlyAvg()
		AllMonths[i].CalcDays()
		fmt.Println(AllMonths[i].Days)
		fmt.Printf("\nMonat: %v ,Tage %v durchschnitt %v\n", AllMonths[i].Name, AllMonths[i].Days, AllMonths[i].AvgTemp)
	}
	//Monate als json Speichern
	for i := 0; i < len(AllMonths); i++ {
		jsonData, err := json.MarshalIndent(AllMonths[i], "", "    ")
		if err != nil {
			fmt.Println("Fehler beim json Kodieren")
			return
		}
		filename := AllMonths[i].Name
		err = saveToFile(filename, jsonData)
		if err != nil {
			fmt.Println("Fehler beim Speichern")
			return

		}
		fmt.Println("JSON-Daten erfolgreich in", filename, "gespeichert.")
	}

}
func saveToFile(dateiname string, daten []byte) error {
	datei, err := os.Create(dateiname)
	if err != nil {
		return err
	}
	defer datei.Close()

	_, err = datei.Write(daten)
	if err != nil {
		return err
	}

	return nil

}
