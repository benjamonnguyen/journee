package handler

import (
	"strconv"
	"time"

	"benjinguyen.me/journee/journal"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func UpsertJournal(e *core.RequestEvent) error {
	var body string
	if err := e.BindBody(&body); err != nil {
		return e.BadRequestError("", err)
	}
	date := e.Request.PathValue("date")
	e.App.Logger().Debug("upserting journal", "body", body, "date", date)

	if _, err := time.Parse(time.DateOnly, date); err != nil {
		return e.BadRequestError("invalid date argument", err)
	}

	record, _ := e.App.FindFirstRecordByFilter(
		"journals",
		"date = {:date} &&  user = {:user}",
		dbx.Params{"date": date, "user": e.Auth.Id},
	)

	// create
	if record == nil {
		collection, err := e.App.FindCollectionByNameOrId("journals")
		if err != nil {
			return err
		}

		record = core.NewRecord(collection)
		record.Set("date", date)
		record.Set("user", e.Auth.Id)
	}

	// update
	el := e.Request.PathValue("el")
	switch el {
	case "content":
		record.Set(el, body)
	case "energyLevel":
		n, err := strconv.Atoi(body)
		if err != nil {
			return e.BadRequestError("", err)
		}
		record.Set(el, n)
	case "emotionID":
		n, err := strconv.Atoi(body)
		if err != nil {
			return e.BadRequestError("", err)
		}
		record.Set(el, n)
	default:
		return e.BadRequestError("invalid form element: "+el, nil)
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
