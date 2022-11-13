package main

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
)

func main() {

	w := wfs.NewWFileSystem("wfs/")
	namespaces, err := w.LoadNamespaces()
	if err != nil {
		fmt.Print(err)
	}

	db := model.WDBFileSystem{
		Namespaces: namespaces,
	}

	jsonString, err := json.Marshal(db)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}
