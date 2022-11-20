package main

import (
	"fmt"

	dbs "github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
)

func main() {

	wf := wfs.NewWFileSystem("wfs/")
	databases, err := wf.LoadDatabases()
	if err != nil {
		fmt.Print(err)
	}

	w := model.WDB{
		Databases: databases,
	}

	d := dbs.WithWDB(w)
	err = d.CreateNewDatabase("whatdb", model.Metadata{}, model.Access{})
	if err != nil {
		fmt.Println(err)
	}

	err = wf.UnloadDatabases(w.Databases)
	if err != nil {
		fmt.Println(err)
	}

}
