package main

import (
	"fmt"
	ns "github.com/TanmoySG/wunderDB/internal/namespaces"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
)

func main() {

	w := wfs.NewWFileSystem("wfs/")
	namespaces, err := w.LoadNamespaces()
	if err != nil {
		fmt.Print(err)
	}

	db := model.WDB{
		Namespaces: namespaces,
	}

	nss := ns.Namespaces(db.Namespaces)
	err = nss.DeleteNamespace("fso")
	if err != nil {
		fmt.Println(err)
	}
	err = nss.CreateNewNamespace("fs0o", model.Metadata{}, model.Access{})
	if err != nil {
		fmt.Println(err)
	}

	err = w.UnloadNamespaces(db.Namespaces)
	if err != nil {
		fmt.Println(err)
	}

}
