package converter

import (
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVConverter{}
			got, err := c.Read(tt.args.csv_path)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CSVConverter{}
			got, err := c.ConvertTo(tt.args.csv_data)
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
