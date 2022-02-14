package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adelberteng/cj_converter/converter"
	"github.com/adelberteng/cj_converter/utils"
)

var logger = utils.GetLogger()

func main() {
	startTime := time.Now()

	var (
		from = flag.String("from", "", "from")
		csvPath = flag.String("csv", "", "csvPath")
		jsonPath = flag.String("json", "", "jsonPath")
	)
	flag.Parse()

	if *from == "csv" {
		c := converter.CSVConverter{}
		csvData, err := c.Read(*csvPath)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}

		jsonData, err := c.ConvertTo(csvData)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}

		err = c.WriteTo(jsonData, *jsonPath)
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}

		logger.Info("Data write in json file successfully!")
	} else if *from == "json" {
		j := converter.JSONConverter{}
		jsonData, err := j.Read(*jsonPath)
		if err != nil {
			logger.Error(err)
		}

		csvData := j.ConvertTo(jsonData)

		err = j.WriteTo(csvData, *csvPath)
		if err != nil {
			logger.Error(err)
		}

		logger.Info("Data write in csv file successfully!")
	} else if *from == "" {
		logger.Error("Please add '-from json' or '-from csv' after cmd!")
	} else {
		logger.Error("wrong 'from' value.")
	}

	fmt.Println("Program running time: " + time.Now().Sub(startTime).String())
}
