package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// чтение аргументов
	var datafile *string

	datafile = flag.String("datafile", "./indications.json", `Path to datafile. Default: "./indications.json"`)

	flag.Parse()

	log.Printf("Datafile: %s\n", *datafile)

	// открытие файла
	f, err := os.Open(*datafile)
	if err != nil {
		log.Printf("ERROR os.Open, %s\n", err)
		return
	}

	// преобразование в срез
	data, err := parseFile(f)
	if err != nil {
		log.Printf("ERROR parseFile, %s\n", err)
		return
	}

	// сортировка данных
	err = sortByDate(data)
	if err != nil {
		log.Printf("ERROR sortByDate, %s\n", err)
		return
	}

	// форматированный вывод
	print(data)
}

// Indication структура показания
type Indication struct {
	Indicator string    `json:"indicator"`
	Value     int       `json:"value"`
	Date      time.Time `json:"date"`
}

// преобразование файла в срез показаний приборов
func parseFile(file *os.File) (data []Indication, err error) {
	var (
		b []byte
	)

	// чтение содержимого файла
	b, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// преобразование из json в срез структур
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return
}

// функция сортировки данных по дате. Заменяет ссылку входного среза на отсортированный срез
func sortByDate(data []Indication) (err error) {

	// sortedData := make([]Indication, len(data))
	// copy(sortedData, data)

	var sortedData []Indication
	sortedData = data

	// timestamp := item.Date.Unix() // числовое представление даты (секунды с 1 января 1970 года). поле для сортировки

	for i := len(sortedData) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if sortedData[j].Date.Unix() > sortedData[j+1].Date.Unix() {
				sortedData[j], sortedData[j+1] = sortedData[j+1], sortedData[j]
			}
		}
	}

	// for _, item := range data {
	// 	log.Printf("data.Date: %v", item.Date)
	// }
	// for _, item := range sortedData {
	// 	log.Printf("sortedData.Date: %v", item.Date)
	// }

	data = sortedData

	return
}

// функция форматированного вывода среза показаний
func print(data []Indication) {
	var (
		res string
	)

	for _, i := range data {
		res = fmt.Sprintf("%s", res)
		res = fmt.Sprintf("value: %v", i.Value) // поле "значение" показания прибора
		res = fmt.Sprintf("indicator: %s", i.Indicator)
		res = fmt.Sprintf("date: %v", i.Date)
		res = fmt.Sprintf("\n")
	}

	// for _, i := range data {
	// 	res = fmt.Sprintf("indicator: %s, value: %d, date: %v\n", i.Indicator, i.Value, i.Date)
	// 	log.Printf(res)
	// }

	log.Printf(res)
}
