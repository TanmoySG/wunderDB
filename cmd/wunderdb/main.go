package main

import (
	"fmt"

	dbs "github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
)

func main() {

	w := wfs.NewWFileSystem("wfs/")
	databases, err := w.LoadDatabases()
	if err != nil {
		fmt.Print(err)
	}

	db := model.WDB{
		Databases: databases,
	}

	d := dbs.UseDatabases(db)
	err = d.CreateNewDatabase("whatdb", model.Metadata{}, model.Access{})
	if err != nil {
		fmt.Println(err)
	}

	err = w.UnloadDatabases(db.Databases)
	if err != nil {
		fmt.Println(err)
	}

}
