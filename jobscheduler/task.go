package jobscheduler

import "encoding/xml"

type Task struct {
	XMLName     xml.Name `xml:"task"`
	Enqueued    string   `xml:"enqueued,attr"`
	Force_start string   `xml:"force_start,attr"`
	Id          string   `xml:"id,attr"`
	Job         string   `xml:"job,attr"`
	Logfile     string   `xml:"log_file,attr"`
	Name        string   `xml:"name,attr"`
	StartAt     string   `xml:"start_at,attr"`
	State       string   `xml:"state,attr"`
}
