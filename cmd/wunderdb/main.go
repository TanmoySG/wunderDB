package main

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/collections"
	"github.com/TanmoySG/wunderDB/internal/data"
	dbs "github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/stats"
	// "github.com/TanmoySG/wunderDB/pkg/schema"
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
	db, err := d.GetDatabase("db-one")
	if err != nil {
		fmt.Println(err)
	}

	c := collections.UseDatabase(*db)
	// err = c.CreateCollection("coll-two", model.Schema{"type": "object"}, model.Metadata{}, model.Access{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	collection, err := c.GetCollection("coll-two")
	if err != nil {
		fmt.Println(err)
	}

	dt := data.UseCollection(*collection)
	dte := "valnd"

	// s, _ := schema.UseSchema(collection.Schema)
	// isValid, _ := s.Validate(dte)

	// filter := map[string]interface{}{
	// 	"key":   "field",
	// 	"value": "value",
	// }

	filter2 := map[string]interface{}{
		"key":   "field",
		"value": "valnd",
	}

	// filter3 := map[string]interface{}{
	// 	"key":   "field",
	// 	"value": "vag",
	// }

	// if isValid {
	// dt.Add(dte)
	// datab, err := dt.Get(filter3)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for id, datan := range datab {
	// 	fmt.Printf("%s , %v\n", id, datan.DataMap())

	// }

	err = dt.Update(dte, filter2)
	if err != nil {
		fmt.Println(err)
	}

	// err = dt.Delete(dte, filter)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// } else {
	// 	fmt.Print("not valid")
	// }

	err = wf.UnloadDatabases(w.Databases)
	if err != nil {
		fmt.Println(err)
	}

	memAlloc := stats.GetAllocatedMemory()
	fmt.Printf("Allocated Memory: %s", memAlloc)

}
