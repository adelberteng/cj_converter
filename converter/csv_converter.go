package converter

import (
	"os"
	"errors"
	"path/filepath"
	"encoding/csv"
	"encoding/json"
	"log"
)


type CSVConverter struct {}

func (c *CSVConverter) Read(csv_path string) ([][]string, error) {
	if _, err := os.Stat(csv_path); os.IsNotExist(err) {
		return [][]string{}, errors.New("the csv path is not exist!")
	} else if file_extension := filepath.Ext(csv_path); file_extension != ".csv" {
		return [][]string{}, errors.New("please provide valid csv path!")
	}

	csv_file, err := os.Open(csv_path)
	if err != nil {
		log.Fatal(err)
		return [][]string{}, errors.New("error happened while opening file.")
	}
	defer csv_file.Close()

	csvReader := csv.NewReader(csv_file)
	csv_data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return [][]string{}, errors.New("error happened while reading csv.")
	}

	return csv_data, nil
}

func (c *CSVConverter) ConvertTo(csv_data [][]string) ([]map[string]string, error) {
	if len(csv_data) == 0 {
		return []map[string]string{}, errors.New("csv have no content, please check csv content.")
	} else if len(csv_data) == 1 {
		return []map[string]string{}, errors.New("csv rows less than 1, please check csv content.")
	}

	var json_data []map[string]string

	csv_header := csv_data[0]
	for _, row := range csv_data[1:] {
		row_map := make(map[string]string)
		for idx, col := range row {
			row_map[csv_header[idx]] = col
		}
		json_data = append(json_data, row_map)
	}

	return json_data, nil
}

func (c *CSVConverter) WriteTo(json_data interface{}, json_path string) error {
	if file_extension := filepath.Ext(json_path); file_extension != ".json" {
		return errors.New("please provide valid json path!")
	}

	j, _ := json.Marshal(json_data)
	json_file, err := os.Create(json_path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer json_file.Close()
	
	_, err = json_file.WriteString(string(j))
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}