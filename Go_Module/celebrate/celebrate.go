package celebrate

import "time"

var message = "Happy New Year"

func GetYear() int {
	return time.Now().Year()
}

func GetMessage() string {
	return message
}
