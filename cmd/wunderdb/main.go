package main

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/model"
)

func main() {

	db := model.WDBFileSystem{
		Namespaces: map[model.Identifier]model.Namespace{
			"gns": {
				Metadata: model.Metadata{},
				Path: model.Paths{},
				Databases: map[model.Identifier]model.Database{
					"db-one": {
						Metadata: model.Metadata{},
						Path:     model.Paths{},
						Collections: map[model.Identifier]model.Collection{
							"coll-one": {
								Path:     model.Paths{},
								Metadata: model.Metadata{},
								Data:     model.Data{},
								Schema:   model.Schema{},
							},
						},
					},
				},
			},
		},
		Users: map[model.Identifier]model.User{},
	}

	// jsonString, err := json.Marshal(db.Namespaces["gns"].Databases["db-one"].Collections["coll-one"])
	jsonString, err := json.Marshal(db)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}
