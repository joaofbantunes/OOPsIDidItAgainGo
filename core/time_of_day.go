package core

type TimeOfDay struct {
	hour   int
	minute int
}

func NewTimeOfDay(hour int, minute int) TimeOfDay {
	return TimeOfDay{
		hour:   hour,
		minute: minute,
	}
}
