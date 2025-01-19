package journal

type UpsertJournalRequest struct {
	Date        string
	Content     string
	EmotionID   EmotionID
	EnergyLevel *int
}

type GetJournalResponse struct {
	Date        string
	Content     string
	EmotionID   EmotionID
	EnergyLevel int
}
