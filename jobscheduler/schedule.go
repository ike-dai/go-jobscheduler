package jobscheduler

import (
	"encoding/xml"
)

type Schedules struct {
	XMLName  xml.Name    `xml:"schedules"`
	Schedule []*Schedule `xml:"schedule,omitempty"`
}

type Schedule struct {
	XMLName   xml.Name   `xml:"schedule"`
	Period    []*Period  `xml:"period,omitempty"`
	Date      []*Date    `xml:"date,omitempty"`
	Weekdays  *Weekdays  `xml:"weekdays,omitempty"`
	Monthdays *Monthdays `xml:"monthdays,omitempty"`
	Ultimos   *Ultimos   `xml:"ultimos,omittempty"`
	Month     []*Month   `xml:"month,omitempty"`
	Holidays  *Holidays  `xml:"holidays,omitempty"`
}

type Period struct {
	XMLName        xml.Name `xml:"period"`
	SingleStart    string   `xml:"single_start,attr,omitempty"`
	LetRun         string   `xml:"let_run,attr,omitempty"`
	Begin          string   `xml:"begin,attr,omitempty"`
	End            string   `xml:"end,attr,omitempty"`
	Repeat         string   `xml:"repeat,attr,omitempty"`
	AbsoluteRepeat string   `xml:"absolute_repeat,attr,omitempty"`
	WhenHoliday    string   `xml:"when_holiday,attr,omitempty"`
}

type Date struct {
	XMLName xml.Name `xml:"date"`
	Date    string   `xml:"date,attr,omitempty"`
}

type Weekdays struct {
	XMLName xml.Name `xml:"weekdays"`
	Day     *Day     `xml:"day,omitempty"`
}

type Day struct {
	XMLName xml.Name  `xml:"day"`
	Day     string    `xml:"day,attr,omitempty"`
	Period  []*Period `xml:"period,omitempty"`
}

type Monthdays struct {
}
type Holidays struct {
}
