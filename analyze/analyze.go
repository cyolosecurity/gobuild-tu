package analyze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Entry struct {
	Package   string    `json:"Package"`
	TimeStart time.Time `json:"TimeStart"`
	TimesDone time.Time `json:"TimeDone"`
}

func JSONFile(graphFile string) {
	fmt.Println("Analyzing file:", graphFile)
	data, err := ioutil.ReadFile(graphFile)
	if err != nil {
		fmt.Println("could not read file:", err)
		return
	}
	JSON(data)
}

func compileTime(e Entry) time.Duration {
	return e.TimesDone.Sub(e.TimeStart)
}

func JSON(jsonData []byte) {
	var entries []Entry
	err := json.Unmarshal(jsonData, &entries)
	if err != nil {
		fmt.Println("JSON:", err)
		return
	}

	sort.Slice(entries, func(i, j int) bool {
		return compileTime(entries[i]) > compileTime(entries[j])
	})

	for _, e := range entries {
		fmt.Println(e.Package, compileTime(e))
	}
}
