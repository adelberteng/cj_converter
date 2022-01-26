package converter

import (
	"os"
	"io/ioutil"
	"encoding/csv"
	"encoding/json"
)

type JSONConverter struct {}


func (j *JSONConverter) Read(json_path string) ([]map[string]string, error) {
	var json_data []map[string]string

	json_file, err := os.Open(json_path)
	if err != nil {
		return []map[string]string{}, err
	}
	defer json_file.Close()

	json_bytedata, _ := ioutil.ReadAll(json_file)
	json.Unmarshal([]byte(json_bytedata), &json_data)

	return json_data, nil
}

func (j *JSONConverter) WriteTo(json_data []map[string]string, csv_path string) error {
	header_keys := []string{}
	for k := range json_data[0] {
		header_keys = append(header_keys, []string{k}...)
	}

	csv_output_file, err := os.Create(csv_path)
	if err != nil {
		return err
	}
	defer csv_output_file.Close()

	writer := csv.NewWriter(csv_output_file)
	defer writer.Flush()

	if err := writer.Write(header_keys); err != nil {
		return err
	}

	for _, row := range json_data {
		var row_data []string
		for _, k := range header_keys {
			row_data = append(row_data, row[k])
		}
		if err := writer.Write(row_data); err != nil {
			return err
		}
	}
	return nil
}

func (j *JSONConverter) WriteWithOrder(json_data []map[string]string, csv_path string, header_keys []string) error {
	csv_output_file, err := os.Create(csv_path)
	if err != nil {
		return err
	}
	defer csv_output_file.Close()
	writer := csv.NewWriter(csv_output_file)
	defer writer.Flush()

	if err := writer.Write(header_keys); err != nil {
		return err
	}

	for _, row := range json_data {
		var row_data []string
		for _, k := range header_keys {
			row_data = append(row_data, row[k])
		}
		if err := writer.Write(row_data); err != nil {
			return err
		}
	}
	return nil
}