package handler

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func UpsertJournal(e *core.RequestEvent) error {
	var form struct {
		Date    string
		Content string
	}
	if err := e.BindBody(&form); err != nil {
		return e.BadRequestError("", err)
	}
	e.App.Logger().Debug("upserting journal", "form", form)

	record, _ := e.App.FindFirstRecordByFilter(
		"journals",
		"date = {:date}",
		dbx.Params{"date": form.Date},
	)

	if record == nil {
		collection, err := e.App.FindCollectionByNameOrId("journals")
		if err != nil {
			return err
		}

		e.App.Logger().Debug("collection", "id", collection.Id)
		e.App.Logger().Debug("auth", "auth", e.Auth)
		e.App.Logger().Debug("user", "id", e.Auth.Id)

		record = core.NewRecord(collection)
		record.Set("date", form.Date)
		record.Set("user", e.Auth.Id)
	}

	record.Set("content", form.Content)
	if err := e.App.Save(record); err != nil {
		return err
	}

	return nil
}
