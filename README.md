# cj-converter
A cli tool for converting format from csv or json to another type in Golang.


<summary><h2 style="display: inline-block">Table of Contents</h2></summary>
<ol>
  <li><a href="#overview">Overview</a></li>
  <li><a href="#requirements">Requirements</a></li>
  <li><a href="#usage">Usage</a></li>
  <li><a href="#tests">Tests</a></li>
  <li><a href="#contact">Contact</a></li>
</ol>


## Overview
---
This tool will take csv file path, read data from csv format, convert it to JSON format, and save into a assign path. The JSON file is contains an array of JSON objects, one CSV line corresponds to one JSON object.


## Requirements
---
1. Provide a command line tool (CLI tool) to read in the CSV file and convert it to a JSON file that contains an array of JSON objects - one CSV line corresponds to one JSON object.
2. (Hint: you need to convert to data objects (e.g. Employee) first and then serialize to JSON format).
Consider that may need more requirements in the future (flexibly read csv header and convert)

It should contain the go.mod file and unittest.

CSV format example:
```
ID,FirstName,LastName,Email,Description,Role,Phone
1,Marc,Smith,marc@glasnostic.com,Writer of Java,Dev,541-754-3010 
2,John,Young,john@glasnostic.com,Interested in MHW,HR,541-75-3010 
3,Peter,Scott,peter@glasnostic.com,amateur boxer,Dev,541-754-3010
```

## Usage
---
``` bash
# compiles the go package
go build csv2json

# pick a csv file path and assign the destination for json file.
./csv2json -csv sample_data/sample.csv -json sample_data/sample.json
```


## Tests
---
``` bash
go test -v
```


## Contact
---
email: adelberteng@gmail.com
