package converter

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type CSVConverter struct{}

func (c *CSVConverter) Read(csvPath string) ([][]string, error) {
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		return [][]string{}, errors.New("the csv path is not exist!")
	} else if file_extension := filepath.Ext(csvPath); file_extension != ".csv" {
		return [][]string{}, errors.New("please provide valid csv path!")
	}

	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
		return [][]string{}, errors.New("error happened while opening file.")
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return [][]string{}, errors.New("error happened while reading csv.")
	}

	return csvData, nil
}

func (c *CSVConverter) ConvertTo(csvData [][]string) ([]map[string]string, error) {
	if len(csvData) == 0 {
		return []map[string]string{}, errors.New("csv have no content, please check csv content.")
	} else if len(csvData) == 1 {
		return []map[string]string{}, errors.New("csv rows less than 1, please check csv content.")
	}

	var jsonData []map[string]string

	csv_header := csvData[0]
	for _, row := range csvData[1:] {
		rowMap := make(map[string]string)
		for idx, col := range row {
			rowMap[csv_header[idx]] = col
		}
		jsonData = append(jsonData, rowMap)
	}

	return jsonData, nil
}

func (c *CSVConverter) WriteTo(jsonData interface{}, jsonPath string) error {
	if fileExtension := filepath.Ext(jsonPath); fileExtension != ".json" {
		return errors.New("please provide valid json path!")
	}

	j, _ := json.Marshal(jsonData)
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer jsonFile.Close()

	_, err = jsonFile.WriteString(string(j))
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
