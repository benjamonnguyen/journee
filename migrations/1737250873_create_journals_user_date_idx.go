package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1407978968")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_zO2oifMmsb` + "`" + ` ON ` + "`" + `journals` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `date` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1407978968")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_user` + "`" + ` ON ` + "`" + `journals` + "`" + ` (user)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
