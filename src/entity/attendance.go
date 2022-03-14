package entity

import "time"

type Attendance struct {
	Name           string    `json:"name"`
	AttendanceTime time.Time `json:"attendance_time"`
}
