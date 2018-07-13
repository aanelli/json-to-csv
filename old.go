package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := jsonToCSV("./orgs.json", "org")
	if err != nil {
		fmt.Println("error: ", err)
	}
	err = jsonToCSV("./spaces.json", "space")
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func jsonToCSV(filename string, orgOrSpace string) error {
	// reading data from JSON File
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	if orgOrSpace == "org" {
		var jsonData []org
	} else if orgOrSpace == "space" {
		var jsonData []space
	} else {
		var jsonData []interface{}
		return errors.New("json function called without an org or space in function call")
	}

	// Unmarshal JSON data
	err = json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return err
	}
	// Create a csv file
	f, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	for _, obj := range jsonData {
		var record []string
		record = append(record, obj.Name)
		record = append(record, obj.GUID)
		if orgOrSpace == "space" {
			record = append(record, obj.OrganizationGUID)
		}

		record = append(record, obj.Apps)
		record = append(record, obj.EmpID)
		record = append(record, "\n")
		w.Write(record)
	}
	w.Flush()
	return nil
}
