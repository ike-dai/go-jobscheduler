package jobscheduler

import (
	"encoding/xml"
)

type RunTime struct {
	XMLName     xml.Name   `xml:"run_time"`
	Begin       string     `xml:"begin,attr,omitempty"`
	End         string     `xml:"end,attr,omitempty"`
	Repeat      string     `xml:"repeat,attr,omitempty"`
	SingleStart string     `xml:"single_start,attr,omitempty"`
	Once        string     `xml:"once,attr,omitempty"`
	WhenHoliday string     `xml:"when_holiday,attr,omitempty"`
	Schedule    string     `xml:"schedule,attr,omitepmty"`
	TimeZone    string     `xml:"time_zone,attr,omittempty"`
	LetRun      string     `xml:"let_run,attr,omitempty"`
	Period      []*Period  `xml:"period,omitempty"`
	Date        []*Date    `xml:"date,omitempty"`
	Weekdays    *Weekdays  `xml:"weekdays,omitempty"`
	Monthdays   *Monthdays `xml:"monthdays,omitempty"`
	Ultimos     *Ultimos   `xml:"ultimos,omittempty"`
	Month       []*Month   `xml:"month,omitempty"`
	Holidays    *Holidays  `xml:"holidays,omitempty"`
	At          []*At      `xml:"at,omitempty"`
}

type At struct {
	XMLName xml.Name `xml:"at"`
	At      string   `xml:"at,attr"` //yyyy-mm-dd HH:MM:SS
}

type Schedules struct {
	XMLName  xml.Name    `xml:"schedules"`
	Schedule []*Schedule `xml:"schedule,omitempty"`
}

type ScheduleConf struct {
	XMLName   xml.Name   `xml:"schedule"`
	Name      string     `xml:"name,attr,omitempty"`
	Period    []*Period  `xml:"period,omitempty"`
	Date      []*Date    `xml:"date,omitempty"`
	Weekdays  *Weekdays  `xml:"weekdays,omitempty"`
	Monthdays *Monthdays `xml:"monthdays,omitempty"`
	Ultimos   *Ultimos   `xml:"ultimos,omittempty"`
	Month     []*Month   `xml:"month,omitempty"`
	Holidays  *Holidays  `xml:"holidays,omitempty"`
}
type Schedule struct {
	XMLName   xml.Name   `xml:"schedule"`
	Active    string     `xml:"active,attr,omitempty"`
	Name      string     `xml:"name,attr,omitempty"`
	Path      string     `xml:"path,attr,omitempty"`
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
	Date    string   `xml:"date,omitempty"`
}

type Weekdays struct {
	XMLName xml.Name `xml:"weekdays"`
	Day     []*Day   `xml:"day,omitempty"`
}

type Day struct {
	XMLName xml.Name  `xml:"day"`
	Day     string    `xml:"day,omitempty"` //in Weekdays: 1(Monday),2(Tuesday)... in Monthdays: 1(1st),2(2nd),3(3rd)... in Ultimos: 0(lastday),1(1st day),2(2nd days)....
	Period  []*Period `xml:"period,omitempty"`
}

//Specific days of month counting from firstday of the month
type Monthdays struct {
	XMLName xml.Name   `xml:"monthdays"`
	Day     []*Day     `xml:"day,omitempty"`
	Weekday []*Weekday `xml:"weekday,omitempty"`
}

//Specific weekday
type Weekday struct {
	XMLName xml.Name  `xml:"weekday"`
	Day     string    `xml:"day,omitempty"`       //sunday,monday,tueseday...
	Which   string    `xml:"which,attr,omiempty"` //1(1st week),2(2nd week)...
	Period  []*Period `xml:"period,omitempty"`
}

//Specific day of month counting from lastday of the month
type Ultimos struct {
	XMLName xml.Name `xml:"ultimos"`
	Day     []*Day   `xml:"day,omitempty"`
}

//Specific Month
type Month struct {
	XMLName   xml.Name   `xml:"month"`
	Month     string     `xml:"month,attr"`
	Period    []*Period  `xml:"period,omitempty"`
	Weekdays  *Weekdays  `xml:"weekdays,omitempty"`
	Monthdays *Monthdays `xml:"monthdays,omitempty"`
	Ultimos   *Ultimos   `xml:"ultimos,omittempty"`
}
type Holidays struct {
	XMLName xml.Name   `xml:"holidays"`
	Holiday []*Holiday `xml:"holiday,omitempty"`
	Include []*Include `xml:"include,omitempty"`
}

type Holiday struct {
	XMLName xml.Name `xml:"holiday"`
	Date    string   `xml:"date,attr,omitempty"` //yyyy-mm-dd
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	File    string   `xml:"file,attr,omitempty"`
}

func (c *Client) ShowSchedules() (*Schedules, *Error) {
	params := &ShowStateInput{What: "schedules"}
	answer := c.ShowState(params)
	if answer.Error != nil {
		return nil, answer.Error
	}
	if answer.State != nil {
		return answer.State.Schedules, nil
	}
	return nil, nil
}

func (c *Client) AddSchedule(schedule *ScheduleConf, folder string) *Answer {
	params := &ModifyHotFolderInput{Folder: folder, Schedule: schedule}
	return c.ModifyHotFolder(params)
}
