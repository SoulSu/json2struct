package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var jsonFile *string
var structName *string

func main() {

	jsonFile = flag.String("jf", "", "input file for json file path")
	structName = flag.String("sn", "Hi", "struct name ")

	flag.Parse()

	if *jsonFile == "" {
		fmt.Println("input file can not empty")
		os.Exit(1)
	}

	_, err := os.Stat(*jsonFile)
	checkErr(err)

	file, err := os.OpenFile(*jsonFile, os.O_RDONLY, 0666)
	checkErr(err)

	fileContentBytes, err := ioutil.ReadAll(file)
	checkErr(err)

	var jsonInter interface{}

	err = json.Unmarshal(fileContentBytes, &jsonInter)
	checkErr(err)

	jsonMap := jsonInter.(map[string]interface{})

	outputString := "package output\n\ntype " + *structName + " struct{\n"
	for k, v := range jsonMap {

		switch v.(type) {
		case int:
			outputString += fmt.Sprintf("%s int `json:\"%s\"`\n", k, k)
		case int64:
			outputString += fmt.Sprintf("%s int64 `json:\"%s\"`\n", k, k)
		case string:
			outputString += fmt.Sprintf("%s string `json:\"%s\"`\n", k, k)
		case float64:
			outputString += fmt.Sprintf("%s float64 `json:\"%s\"`\n", k, k)
		case float32:
			outputString += fmt.Sprintf("%s float32 `json:\"%s\"`\n", k, k)
		case bool:
			outputString += fmt.Sprintf("%s bool `json:\"%s\"`\n", k, k)
		default:
			fmt.Printf("no type booooooooooom!!!!")
			os.Exit(3)
		}
	}
	outputString = strings.TrimRight(outputString, ",\n")

	outputString += "\n}"

	outputFile, err := os.Create("output.go")
	checkErr(err)

	_, err = outputFile.WriteString(fmt.Sprintf(outputString))
	checkErr(err)

	// format
	// cmd := exec.Command("gofmt.exe -w output.go")
	// err = cmd.Run()
	// checkErr(err)

	fmt.Printf(outputString)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
