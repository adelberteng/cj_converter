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
This tool will take csv or json file path, read data from specify format, convert it to another format, and save into a assign path.

## Requirements
---


CSV format example:
```
Uid,Name,Gender,Age
1,Albert,Male,28
2,Alice,Female,20
3,Bob,Male,30
4,Charlie,Male,40
```

## Usage
---
``` bash
# compiles the go package
go build -o cj_converter

# choose from-type as csv, pick a csv file path and assign the destination for json file.
./cj_converter -from csv -csv sample/sample_origin.csv -json sample/sample.json

# json to csv
./cj_converter -from json -json sample/sample.json -csv sample/sample.csv 

```


## Tests
---
``` bash
go test -v
```


## Contact
---
email: adelberteng@gmail.com
