package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func decodeJSON(m map[string]interface{}) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			for _, value := range decodeJSON(vv) {
				values = append(values, value)
			}
		case string:
			values = append(values, vv)
		case float64:
			values = append(values, strconv.FormatFloat(vv, 'f', -1, 64))
		case []interface{}:
			// Arrays aren't currently handled, since you haven't indicated that we should
			// and it's non-trivial to do so.
		case bool:
			values = append(values, strconv.FormatBool(vv))
		case nil:
			values = append(values, "nil")
		}
	}
	return values
}

func convertJSON(inputfile string, outfile string) error {
	data, err := ioutil.ReadFile(inputfile)
	if err != nil {
		return err
	}
	var d interface{}
	err = json.Unmarshal(data, &d)
	if err != nil {
		log.Fatal("Failed to unmarshal")
		return err
	}
	values := decodeJSON(d.(map[string]interface{}))
	fmt.Println(values)

	f, err := os.Create(outfile)
	if err != nil {
		log.Fatal("Failed to create outputfile")
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.Write(values); err != nil {
		log.Fatal("Failed to write to file")
		return err
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal("Failed to flush outputfile.csv")
		return err
	}
	return nil
}

func main() {

	err := convertJSON("./orgs.json", "orgs.csv")
	if err != nil {
		log.Fatal("failed to open orgs json file")
	}

	err = convertJSON("./spaces.json", "spaces.csv")
	if err != nil {
		log.Fatal("failed to open orgs json file")
	}
}
