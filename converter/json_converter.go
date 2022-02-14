package converter

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
)

type JSONConverter struct{}

func (j *JSONConverter) Read(jsonPath string) ([]map[string]string, error) {
	var jsonData []map[string]string

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return []map[string]string{}, err
	}
	defer jsonFile.Close()

	jsonByteData, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(jsonByteData), &jsonData)

	return jsonData, nil
}

func (j *JSONConverter) ConvertTo(jsonData []map[string]string) [][]string {
	var csvData [][]string

	headerKeys := []string{}
	for k := range jsonData[0] {
		headerKeys = append(headerKeys, k)
	}

	csvData = append(csvData, headerKeys)
	for _, row := range jsonData {
		var rowData []string
		for _, k := range headerKeys {
			rowData = append(rowData, row[k])
		}
		csvData = append(csvData, rowData)
	}

	return csvData
}

// ConvertWithOrder be used when you need specific order with header or only need specific columns.
func (j *JSONConverter) ConvertWithOrder(jsonData []map[string]string, headerKeys []string) [][]string {
	var csvData [][]string

	csvData = append(csvData, headerKeys)
	for _, row := range jsonData {
		var rowData []string
		for _, k := range headerKeys {
			rowData = append(rowData, row[k])
		}
		csvData = append(csvData, rowData)
	}

	return csvData
}

func (j *JSONConverter) WriteTo(csvData [][]string, csvPath string) error {
	csvFile, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, row := range csvData {
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
