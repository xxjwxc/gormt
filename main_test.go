package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/gormt/data/view/model"
)

func TestDomain(t *testing.T) {
	strjson := tools.ReadFile("test.txt")
	var str string
	for _, v := range strjson {
		str += v
	}
	var pkg model.DBInfo
	json.Unmarshal([]byte(str), &pkg)
	// out, _ := json.Marshal(pkg)
	// tools.WriteFile("test.txt", []string{string(out)}, true)

	list := model.Generate(pkg)
	fmt.Println(list)
}
