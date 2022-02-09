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

func (j *JSONConverter) WriteTo(jsonData []map[string]string, csvPath string) error {
	headerKeys := []string{}
	for k := range jsonData[0] {
		headerKeys = append(headerKeys, []string{k}...)
	}

	csvOutputFile, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer csvOutputFile.Close()

	writer := csv.NewWriter(csvOutputFile)
	defer writer.Flush()

	err = writer.Write(headerKeys)
	if err != nil {
		return err
	}

	for _, row := range jsonData {
		var rowData []string
		for _, k := range headerKeys {
			rowData = append(rowData, row[k])
		}
		if err := writer.Write(rowData); err != nil {
			return err
		}
	}
	return nil
}

func (j *JSONConverter) WriteWithOrder(jsonData []map[string]string, csvPath string, headerKeys []string) error {
	csvOutputFile, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer csvOutputFile.Close()
	writer := csv.NewWriter(csvOutputFile)
	defer writer.Flush()

	err = writer.Write(headerKeys)
	if err != nil {
		return err
	}

	for _, row := range jsonData {
		var rowData []string
		for _, k := range headerKeys {
			rowData = append(rowData, row[k])
		}
		err = writer.Write(rowData)
		if err != nil {
			return err
		}
	}
	return nil
}
