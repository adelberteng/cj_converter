package converter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCSVConverter_Read(t *testing.T) {
	type args struct {
		csv_path string
	}
	tests := []struct {
		name    string
		c       *CSVConverter
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			"normal",
			&CSVConverter{},
			args{"../sample/sample_for_test.csv"},
			[][]string{
				{"Uid", "Name", "Gender", "Age"},
				{"1", "Albert", "Male", "28"},
			},
			false,
		},
		{
			"the csv path is not exist",
			&CSVConverter{},
			args{"../sample/sample_not_exist.csv"},
			[][]string{},
			true,
		},
		{
			"invalid csv path",
			&CSVConverter{},
			args{"../sample/sample_for_test.json"},
			[][]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVConverter{}
			got, err := c.Read(tt.args.csv_path)
			if err != nil {
				fmt.Println("testcase err info: ", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVConverter.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSVConverter.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSVConverter_ConvertTo(t *testing.T) {
	type args struct {
		csv_data [][]string
	}
	tests := []struct {
		name    string
		c       *CSVConverter
		args    args
		want    []map[string]string
		wantErr bool
	}{
		{
			"normal",
			&CSVConverter{},
			args{
				[][]string{
					{"Uid", "Name", "Gender", "Age"},
					{"1", "Albert", "Male", "28"},
				},
			},
			[]map[string]string{
				{"Uid":"1", "Name":"Albert", "Gender":"Male", "Age":"28"},
			},
			false,
		},
		{
			"csv have no content",
			&CSVConverter{},
			args{
				[][]string{},
			},
			[]map[string]string{},
			true,
		},
		{
			"csv rows less than 1",
			&CSVConverter{},
			args{
				[][]string{
					{"Uid", "Name", "Gender", "Age"},
				},
			},
			[]map[string]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVConverter{}
			got, err := c.ConvertTo(tt.args.csv_data)
			if err != nil {
				fmt.Println("testcase err info: ", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVConverter.ConvertTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSVConverter.ConvertTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSVConverter_WriteTo(t *testing.T) {
	type args struct {
		json_data interface{}
		json_path string
	}
	tests := []struct {
		name    string
		c       *CSVConverter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVConverter{}
			if err := c.WriteTo(tt.args.json_data, tt.args.json_path); (err != nil) != tt.wantErr {
				t.Errorf("CSVConverter.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
