package journal

type energyLevel int

const (
	Dead energyLevel = iota + 1
	DozingOff
	Exhausted
	Tired
	Neutral
	OK
	Good
	Great
	Amazing
)

var energyThresholds = []energyThreshold{
	{
		ID:        Dead,
		Emoji:     "💀",
		Threshold: 11,
	},
	{
		ID:        DozingOff,
		Emoji:     "😪",
		Threshold: 22,
	},
	{
		ID:        Exhausted,
		Emoji:     "🥱",
		Threshold: 33,
	},
	{
		ID:        Tired,
		Emoji:     "😮‍💨",
		Threshold: 44,
	},
	{
		ID:        Neutral,
		Emoji:     "😑",
		Threshold: 55,
	},
	{
		ID:        OK,
		Emoji:     "😌",
		Threshold: 66,
	},
	{
		ID:        Good,
		Emoji:     "🙂",
		Threshold: 77,
	},
	{
		ID:        Great,
		Emoji:     "😊",
		Threshold: 88,
	},
	{
		ID:        Amazing,
		Emoji:     "🤩",
		Threshold: 100,
	},
}

type energyThreshold struct {
	ID        energyLevel
	Emoji     string
	Threshold int
}
