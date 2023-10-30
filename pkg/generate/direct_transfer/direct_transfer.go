package direct_transfer

import (
	"fmt"
	"math/rand"
	"time"
)

func NewNoReff() string {
	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, _ := time.Parse(layOut, timeStampString)
	hour, minute, second := timeStamp.Clock()

	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()

	num := rand.Intn(4)

	no_reff := "DT" + fmt.Sprint(year) + fmt.Sprint(int(month)) + fmt.Sprint(day) + fmt.Sprint(hour) + fmt.Sprint(minute) + fmt.Sprint(second) + fmt.Sprint(num)

	return no_reff
}
