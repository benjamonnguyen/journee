package migrations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		if err := renameDateOldUp(app); err != nil {
			return err
		}

		collection, err := app.FindCollectionByNameOrId("journals")
		if err != nil {
			return err
		}
		collection.Fields.Add(&core.DateField{
			Name:     "date",
			Required: true,
		})
		app.Save(collection)

		records, err := app.FindAllRecords("journals")
		if err != nil {
			return err
		}
		for _, record := range records {
			date, _ := time.Parse(time.DateOnly, record.GetString("date_old"))
			fmt.Println(record.Id, record.GetString("date_old"), date)

			record.Set("date", date)

			// new min == 0
			if record.GetInt("energy_level") < 0 {
				record.Set("energy_level", 0)
			}

			app.Save(record)
		}

		collection.Fields.RemoveByName("date_old")
		app.Save(collection)

		return nil
	}, func(app core.App) error {
		if err := renameDateOldDown(app); err != nil {
			return err
		}

		return nil
	})
}

func renameDateOldUp(app core.App) error {

	collection, err := app.FindCollectionByNameOrId("pbc_1407978968")
	if err != nil {
		return err
	}

	// update collection data
	if err := json.Unmarshal([]byte(`{
		"indexes": []
	}`), &collection); err != nil {
		return err
	}

	// update field
	if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
		"autogeneratePattern": "",
		"hidden": false,
		"id": "text2862495610",
		"max": 0,
		"min": 0,
		"name": "date_old",
		"pattern": "",
		"presentable": false,
		"primaryKey": false,
		"required": false,
		"system": false,
		"type": "text"
	}`)); err != nil {
		return err
	}

	return app.Save(collection)
}

func renameDateOldDown(app core.App) error {
	collection, err := app.FindCollectionByNameOrId("pbc_1407978968")
	if err != nil {
		return err
	}

	// update collection data
	if err := json.Unmarshal([]byte(`{
		"indexes": [
			"CREATE UNIQUE INDEX `+"`"+`idx_zO2oifMmsb`+"`"+` ON `+"`"+`journals`+"`"+` (\n  `+"`"+`user`+"`"+`,\n  `+"`"+`date`+"`"+`\n)"
		]
	}`), &collection); err != nil {
		return err
	}

	// update field
	if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
		"autogeneratePattern": "",
		"hidden": false,
		"id": "text2862495610",
		"max": 0,
		"min": 0,
		"name": "date",
		"pattern": "",
		"presentable": false,
		"primaryKey": false,
		"required": false,
		"system": false,
		"type": "text"
	}`)); err != nil {
		return err
	}

	return app.Save(collection)
}
