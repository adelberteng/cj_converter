package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	config "gopkg.in/ini.v1"

	"github.com/adelberteng/cj_converter/converter"
	goLogger "github.com/adelberteng/go_logger"
)


func main() {
	start_time := time.Now()

	cfg, err := config.Load("conf/config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logDir := cfg.Section("dev").Key("log_dir").String()
	logName := cfg.Section("dev").Key("log_file_name").String()

	os.MkdirAll(logDir, 0766)
	logFile, err := os.OpenFile(logDir+"/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil && err == os.ErrNotExist {
		os.Create(logDir + "/" + logName)
	} else if err != nil {
		log.Fatalf("log file open error : %v", err)
	}
	defer logFile.Close()
	logger := goLogger.CreateLogger(logFile, "debug", "")

	var from = flag.String("from", "", "from")
	var csvPath = flag.String("csv", "", "csvPath")
	var jsonPath = flag.String("json", "", "jsonPath")
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

		err = j.WriteTo(jsonData, *csvPath)
		if err != nil {
			fmt.Println(err)
		}
	} else if *from == "" {
		logger.Error("Please add '-from json' or '-from csv' after cmd!")
	} else {
		logger.Error("wrong 'from' value.")
	}

	fmt.Println("Program running time: " + time.Now().Sub(start_time).String())
}
