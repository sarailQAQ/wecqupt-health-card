package util

func GetKey(day, hour int) string {
	if ! (0 < day && day < 32 && 0 < hour && hour < 25) {
		return ""
	}
	return dateCode[day - 1] + hourCode[hour - 1]
}

