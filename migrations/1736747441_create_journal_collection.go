package migrations

import (
	"benjinguyen.me/journee/util"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection := core.NewBaseCollection("journals")

		// set rules
		collection.ViewRule = util.StrPtr("@request.auth.id != ''")
		collection.CreateRule = util.StrPtr(
			"@request.auth.id != '' && @request.body.user = @request.auth.id",
		)
		collection.UpdateRule = util.StrPtr(`
			@request.auth.id != '' &&
			user = @request.auth.id &&
			(@request.body.user:isset = false || @request.body.user = @request.auth.id)
		`)

		// add fields
		collection.Fields.Add(&core.EditorField{
			Name:     "content",
			Required: true,
		})

		usersCollection, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		collection.Fields.Add(&core.RelationField{
			Name:          "user",
			Required:      true,
			CascadeDelete: true,
			CollectionId:  usersCollection.Id,
		})

		collection.Fields.Add(&core.TextField{
			Name:     "date",
			Required: true,
		})

		// add autodate/timestamp fields (created/updated)
		collection.Fields.Add(&core.AutodateField{
			Name:     "created",
			OnCreate: true,
		})
		collection.Fields.Add(&core.AutodateField{
			Name:     "updated",
			OnCreate: true,
			OnUpdate: true,
		})

		// add index
		collection.AddIndex("idx_user", false, "user", "")

		// validate and persist
		// (use SaveNoValidate to skip fields validation)
		err = app.Save(collection)
		if err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
