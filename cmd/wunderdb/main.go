package main

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
)

func main() {

	w := wfs.NewWFileSystem("wfs")
	namespaces, err := w.LoadNamespaces()
	if err != nil {
		fmt.Print(err)
	}

	db := model.WDBFileSystem{
		Namespaces: namespaces,
	}

	data := db.Namespaces["ns1"].Databases["db-one"].Collections["coll-one"].Data
	data["field0"] = "val2w"

	db.Namespaces["ns1"].Databases["db-one"].Collections["coll-one"].Data = data

	err = w.UnloadNamespaces(db.Namespaces)
	if err != nil {
		fmt.Println(err)
	}

}
