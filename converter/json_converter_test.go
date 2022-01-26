package converter

import (
	"reflect"
	"testing"
)

func TestJSONConverter_Read(t *testing.T) {
	type args struct {
		json_path string
	}
	tests := []struct {
		name    string
		j       *JSONConverter
		args    args
		want    []map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			got, err := j.Read(tt.args.json_path)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONConverter.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONConverter.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONConverter_WriteTo(t *testing.T) {
	type args struct {
		json_data []map[string]string
		csv_path  string
	}
	tests := []struct {
		name    string
		j       *JSONConverter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			if err := j.WriteTo(tt.args.json_data, tt.args.csv_path); (err != nil) != tt.wantErr {
				t.Errorf("JSONConverter.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONConverter_WriteWithOrder(t *testing.T) {
	type args struct {
		json_data   []map[string]string
		csv_path    string
		header_keys []string
	}
	tests := []struct {
		name    string
		j       *JSONConverter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			if err := j.WriteWithOrder(tt.args.json_data, tt.args.csv_path, tt.args.header_keys); (err != nil) != tt.wantErr {
				t.Errorf("JSONConverter.WriteWithOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
