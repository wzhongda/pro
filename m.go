package main

import (
	"fmt"
	"time"
)

func GetTimeSlice(st, et int64) []string {
	fmt.Println(st, et)
	date := make([]string, 0)
	for i := st; i <= et; i += 86400 {
		tm := time.Unix(i, 0)
		date = append(date, tm.Format("2006-01-02"))
	}
	return date
}

func main() {
	/*timeFormat := "2006-01-02 15:04:05"
	dateNow := time.Now()
	PredateNow := time.Now().AddDate(0, 0, -6)
	startTime := time.Date(PredateNow.Year(), PredateNow.Month(), PredateNow.Day(), 0, 0, 0, 0, PredateNow.Location()).Format(timeFormat)
	endTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 23, 59, 59, 0, dateNow.Location()).Format(timeFormat)
	fmt.Println(startTime, "----", endTime)*/
	var low_risk_7 [7]int64
	var low_risk_30 [30]int64
	result := make(map[string]map[int64]int64)
	//map[2022-06-30:map[1:21] 2022-07-01:map[1:2] 2022-07-04:map[1:2]]
	level := make(map[int64]int64)
	level[1] = 21
	level1 := make(map[int64]int64)
	level1[1] = 2
	level2 := make(map[int64]int64)
	level2[1] = 2
	result["2022-06-30"] = level
	result["2022-07-01"] = level1
	result["2022-07-04"] = level2

	TimeSlice := GetTimeSlice(1656518400, 1657123199)
	TimeLen := len(TimeSlice)
	for index, item := range TimeSlice {
		for i := int64(0); i <= 3; i++ {
			if _, ok := result[item]; ok {
				v, _ := result[item][i]
				if i == 0 {
					if TimeLen == 7 {
						low_risk_7[index] = v
					} else {
						low_risk_30[index] = v
					}
				}
			}
		}
	}
	fmt.Printf("result %+v", low_risk_7)
}
