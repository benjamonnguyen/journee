package journal

type UpsertJournalRequest struct {
	Content     string    `json:"content,omitempty"`
	EmotionID   EmotionID `json:"emotionID,omitempty"`
	EnergyLevel int       `json:"energyLevel,omitempty"`
}

type GetJournalResponse struct {
	Date        string    `json:"date"`
	Content     string    `json:"content"`
	EmotionID   EmotionID `json:"emotionID"`
	EnergyLevel int       `json:"energyLevel"`
}
