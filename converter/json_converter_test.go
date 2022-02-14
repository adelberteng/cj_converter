package converter

import (
	"fmt"
	"reflect"
	"testing"
)

func compareSlicesElementEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, i := range s1 {
		containElement := false
		for _, j := range s2 {
			if i == j {
				containElement = true
			}
		}
		if containElement == false {
			return false
		}
	}

	return true
}

func compareInteralSlice(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !compareSlicesElementEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

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
		{
			"normal",
			&JSONConverter{},
			args{"../sample/sample_for_test.json"},
			[]map[string]string{
				{"Uid": "1", "Name": "Albert", "Gender": "Male", "Age": "28"},
			},
			false,
		},
		{
			"no such file or directory",
			&JSONConverter{},
			args{"../nonexisted/sample_for_test.json"},
			[]map[string]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			got, err := j.Read(tt.args.json_path)
			if err != nil {
				fmt.Println(err)
			}
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

func TestJSONConverter_ConvertTo(t *testing.T) {
	type args struct {
		jsonData []map[string]string
	}
	tests := []struct {
		name string
		j    *JSONConverter
		args args
		want [][]string
	}{
		{
			"normal",
			&JSONConverter{},
			args{
				[]map[string]string{
					{"Uid": "1", "Name": "Albert", "Gender": "Male", "Age": "28"},
				},
			},
			[][]string{
				{"Uid", "Name", "Gender", "Age"},
				{"1", "Albert", "Male", "28"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			if got := j.ConvertTo(tt.args.jsonData); !compareInteralSlice(got, tt.want) {
				t.Errorf("JSONConverter.ConvertTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONConverter_ConvertWithOrder(t *testing.T) {
	type args struct {
		jsonData   []map[string]string
		headerKeys []string
	}
	tests := []struct {
		name string
		j    *JSONConverter
		args args
		want [][]string
	}{
		{
			"normal",
			&JSONConverter{},
			args{
				[]map[string]string{
					{"Uid": "1", "Name": "Albert", "Gender": "Male", "Age": "28"},
				},
				[]string{"Uid", "Name", "Gender", "Age"},
			},
			[][]string{
				{"Uid", "Name", "Gender", "Age"},
				{"1", "Albert", "Male", "28"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			if got := j.ConvertWithOrder(tt.args.jsonData, tt.args.headerKeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONConverter.ConvertWithOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONConverter_WriteTo(t *testing.T) {
	type args struct {
		csvData [][]string
		csvPath string
	}
	tests := []struct {
		name    string
		j       *JSONConverter
		args    args
		wantErr bool
	}{
		{
			"normal",
			&JSONConverter{},
			args{
				[][]string{
					{"Uid", "Name", "Gender", "Age"},
					{"1", "Albert", "Male", "28"},
				},
				"../sample/sample_for_test.csv",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONConverter{}
			if err := j.WriteTo(tt.args.csvData, tt.args.csvPath); (err != nil) != tt.wantErr {
				t.Errorf("JSONConverter.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
