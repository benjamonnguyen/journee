package journal

import "fmt"

type EmotionID int

const (
	EmotionAngry EmotionID = iota + 1
	EmotionAnxious
	EmotionDepressed
	EmotionSad
	EmotionGrumpy
	EmotionNeutral
	EmotionHopeful
	EmotionHappy
	EmotionJoyous
	EmotionEstatic
)

var emotionOptions = []emotionOption{
	{
		ID:    EmotionAngry,
		Name:  "Angry",
		Emoji: "😠",
	},
	{
		ID:    EmotionAnxious,
		Name:  "Anxious",
		Emoji: "😖",
	},
	{
		ID:    EmotionDepressed,
		Name:  "Depressed",
		Emoji: "😞",
	},
	{
		ID:    EmotionSad,
		Name:  "Sad",
		Emoji: "😔",
	},
	{
		ID:    EmotionGrumpy,
		Name:  "Grumpy",
		Emoji: "😒",
	},
	{
		ID:    EmotionNeutral,
		Name:  "Neutral",
		Emoji: "😑",
	},
	{
		ID:    EmotionHopeful,
		Name:  "Hopeful",
		Emoji: "😌",
	},
	{
		ID:    EmotionHappy,
		Name:  "Happy",
		Emoji: "🙂",
	},
	{
		ID:    EmotionJoyous,
		Name:  "Joyous",
		Emoji: "😄",
	},
	{
		ID:    EmotionEstatic,
		Name:  "Estatic",
		Emoji: "🤩",
	},
}

type emotionOption struct {
	ID    EmotionID
	Name  string
	Emoji string
}

func (o emotionOption) getID() string {
	return fmt.Sprintf("%d", o.ID)
}
