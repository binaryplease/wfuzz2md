package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// WfuzzResult holds the output form wfuzz's json
type WfuzzResult []struct {
	Chars    int           `json:"chars"`
	Code     int           `json:"code"`
	Payload  string        `json:"payload"`
	Lines    int           `json:"lines"`
	Location string        `json:"location"`
	Method   string        `json:"method"`
	PostData []interface{} `json:"post_data"`
	Server   string        `json:"server"`
	URL      string        `json:"url"`
	Words    int           `json:"words"`
}

func main() {
	var results WfuzzResult

	err := json.NewDecoder(os.Stdin).Decode(&results)
	if err != nil {
		log.Fatal("Could not parse JSON. Did you run wfuzz with \"-o json\" ?")
	}

	fmt.Println("| Code | Chars | URL |")
	fmt.Println("|---|---|---|")
	for _, r := range results {
		fmt.Println("|", r.Code, "|", r.Chars, "\t| ", r.URL, " |")
	}
}
