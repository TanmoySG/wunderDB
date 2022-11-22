package main

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/collections"
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
	db, err := d.GetDatabase("whatdb")
	if err != nil {
		fmt.Println(err)
	}

	c := collections.UseDatabase(*db)
	err = c.DeleteCollection("v091")
	if err != nil {
		fmt.Println(err)
	}

	err = wf.UnloadDatabases(w.Databases)
	if err != nil {
		fmt.Println(err)
	}

}
