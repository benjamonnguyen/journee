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
	EmotionHappy
	EmotionJoyous
	EmotionEstatic
)

var emotionOptions = []emotionOption{
	{
		ID:    EmotionAngry,
		Emoji: "😡",
	},
	{
		ID:    EmotionAnxious,
		Emoji: "😰",
	},
	{
		ID:    EmotionDepressed,
		Emoji: "😞",
	},
	{
		ID:    EmotionSad,
		Emoji: "😔",
	},
	{
		ID:    EmotionGrumpy,
		Emoji: "😒",
	},
	{
		ID:    EmotionNeutral,
		Emoji: "😑",
	},
	{
		ID:    EmotionHappy,
		Emoji: "🙂",
	},
	{
		ID:    EmotionJoyous,
		Emoji: "😄",
	},
	{
		ID:    EmotionEstatic,
		Emoji: "🤩",
	},
}

type emotionOption struct {
	ID    EmotionID
	Emoji string
}

func (o emotionOption) getID() string {
	return fmt.Sprintf("%d", o.ID)
}
