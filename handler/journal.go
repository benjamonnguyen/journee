package handler

import (
	"benjinguyen.me/journee/journal"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func UpsertJournal(e *core.RequestEvent) error {
	var request journal.UpsertJournalRequest
	if err := e.BindBody(&request); err != nil {
		return e.BadRequestError("", err)
	}
	e.App.Logger().Debug("upserting journal", "request", request)

	record, _ := e.App.FindFirstRecordByFilter(
		"journals",
		"date = {:date} &&  user = {:user}",
		dbx.Params{"date": request.Date, "user": e.Auth.Id},
	)

	if record == nil {
		collection, err := e.App.FindCollectionByNameOrId("journals")
		if err != nil {
			return err
		}

		record = core.NewRecord(collection)
		record.Set("date", request.Date)
		record.Set("user", e.Auth.Id)
	}

	// update
	record.Set("content", request.Content)
	record.Set("emotion_id", int(request.EmotionID))
	if request.EnergyLevel == nil {
		record.Set("energy_level", -1)
	} else {
		record.Set("energy_level", *request.EnergyLevel)
	}
	e.App.Logger().Debug("saving journal record", "record", record)
	if err := e.App.Save(record); err != nil {
		return err
	}

	return nil
}

func GetJournal(e *core.RequestEvent) error {
	date := e.Request.PathValue("date")
	record, _ := e.App.FindFirstRecordByFilter(
		"journals",
		"date = {:date} &&  user = {:user}",
		dbx.Params{"date": date, "user": e.Auth.Id},
	)
	if record == nil {
		return e.NotFoundError("", nil)
	}

	return e.JSON(200, journal.GetJournalResponse{
		Date:        date,
		Content:     record.GetString("content"),
		EmotionID:   journal.EmotionID(record.GetInt("emotion_id")),
		EnergyLevel: record.GetInt("energy_level"),
	})
}
