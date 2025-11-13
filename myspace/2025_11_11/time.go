package time

type Time int


func fromSeconds(t int) Time {
	return Time(t)
}

func fromMinutes(t int) Time {
	return Time(t * 60)
}

func fromHours(t int) Time {
	return Time(t * 3600)
}

func (t Time) seconds() int {
	return int(t)
}

func (t Time) minutes() int {
	return int(t/60)
}

func (t Time) hours() int {
	return int(t / 3600)
}
