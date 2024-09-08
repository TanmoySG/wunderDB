package steps

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/internal/config"
)

const databasesBasePathFormat = "%s/databases/databases_persisted.json"

func MigrateDataToRecords(c config.Config) error {
	dbPath := fmt.Sprintf(databasesBasePathFormat, c.PersistantStoragePath)

	dbs, err := loadPersistentDataWithOldModel(dbPath)
	if err != nil {
		return fmt.Errorf("error loading wfs: %s", err)
	}

	for dbKey, db := range dbs {
		if _, ok := db["collections"]; !ok {
			return fmt.Errorf("no collections")
		}

		for cKey, collection := range db["collections"] {
			if _, ok := collection.(map[string]interface{})["data"]; !ok {
				fmt.Println("no data found")
				continue
			}

			rec := dbs[dbKey]["collections"][cKey].(map[string]interface{})["data"]
			dbs[dbKey]["collections"][cKey].(map[string]interface{})["records"] = rec
			delete(dbs[dbKey]["collections"][cKey].(map[string]interface{}), "data")
		}
	}

	return unloadPersistentDataWithNewModel(dbs, dbPath)
}

func loadPersistentDataWithOldModel(dbPath string) (map[string]map[string]map[string]interface{}, error) {
	databases := make(map[string]map[string]map[string]interface{})
	persitedDatabasesBytes, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("error reading database file: %s", err)
	}

	err = json.Unmarshal(persitedDatabasesBytes, &databases)
	if err != nil {
		return nil, fmt.Errorf("error marshaling database file: %s", err)
	}

	return databases, nil
}

func unloadPersistentDataWithNewModel(dbs map[string]map[string]map[string]interface{}, dbPath string) error {
	dbsBytes, err := json.Marshal(dbs)
	if err != nil {
		return fmt.Errorf("error marshaling database file: %s", err)
	}

	err = os.WriteFile(dbPath, dbsBytes, 0740)
	if err != nil {
		return err
	}

	return nil
}
