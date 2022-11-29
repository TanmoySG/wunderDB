package main

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/fsLoader"
	"github.com/TanmoySG/wunderDB/model"
	wdbClient "github.com/TanmoySG/wunderDB/wdb"
)

func main() {
	fmt.Printf("Hello World!")

	fs := fsLoader.NewWFileSystem("wfs/")

	loadedDatabase, _ := fs.LoadDatabases()
	db := databases.WithWDB(loadedDatabase)
	wdbc := wdbClient.NewWdbClient(db)

	err := wdbc.AddDatabase("freaked", model.Metadata{})
	if err != nil {
		fmt.Printf("err %s", err)
	}

	err = fs.UnloadDatabases(db)
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
}
