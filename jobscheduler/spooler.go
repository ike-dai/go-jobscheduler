package jobscheduler

import (
	"encoding/xml"
)

type Spooler struct {
	XMLName xml.Name `xml:"spooler"`
	Answer  *Answer  `xml:"answer"`
}

type Answer struct {
	XMLName  xml.Name  `xml:"answer"`
	Ok       *Ok       `xml:"ok,omitempty"`
	Error    *Error    `xml:"ERROR,omitempty"`
	State    *State    `xml:"state,omitempty"`
	JobChain *JobChain `xml:"job_chain,omitempty"`
	History  *History  `xml:"history,omitempty"`
}

type Ok struct {
	XMLName xml.Name `xml:"ok"`
	Task    *Task    `xml:"task,omitempty"`
}
type Error struct {
	XMLName xml.Name `xml:"ERROR"`
	Code    string   `xml:"code,attr"`
	Text    string   `xml:"text,attr"`
	Time    string   `xml:"time,attr"`
}

type State struct {
	XMLName    xml.Name   `xml:"state"`
	ConfigFile string     `xml:"config_file,attr"`
	Db         string     `xml:"db,attr"`
	DbWaiting  string     `xml:"db_waiting,attr"`
	Id         string     `xml:"id,attr"`
	LogFile    string     `xml:"log_file,attr"`
	Loop       string     `xml:"loop,attr"`
	Jobs       *Jobs      `xml:"jobs,omitempty"`
	JobChains  *JobChains `xml:"job_chains,omitempty"`
	Schedules  *Schedules `xml:"schedules,omitempty"`
	//ProcessClasses ProcessClasses `xml:"process_classes,omitempty"`
}

type History struct {
	XMLName      xml.Name        `xml:"history"`
	HistoryEntry []*HistoryEntry `xml:"history.entry,omitempty"`
}

type HistoryEntry struct {
	XMLName   xml.Name `xml:"history.entry"`
	Cause     string   `xml:"cause,attr,omitempty"`
	EndTime   string   `xml:"end_time,attr,omitempty"`
	Error     string   `xml:"error,attr,omitempty"`
	ExitCode  string   `xml:"exit_code,attr,omitempty"`
	Id        string   `xml:"id,attr,omitempty"`
	JobName   string   `xml:"job_name,attr,omitempty"`
	Pid       string   `xml:"pid,attr,omitempty"`
	SpoolerId string   `xml:"spooler_id,attr,omitempty"`
	StartTime string   `xml:"start_time,attr,omitempty"`
	Steps     string   `xml:"steps,attr,omitempty"`
	Task      string   `xml:"task,attr,omitempty"`
}

type Script struct {
	XMLName       xml.Name `xml:"script"`
	ComClass      string   `xml:"com_class,attr,omitempty"`
	FileName      string   `xml:"file_name,attr,omitempty"`
	JavaClass     string   `xml:"java_class,attr,omitempty"`
	JavaClassPath string   `xml:"java_class_path,attr,omitempty"`
	Language      string   `xml:"language,attr"`
	Script        string   `xml:",chardata"`
}

type ShowStateInput struct {
	XMLName   xml.Name `xml:"show_state"`
	MaxOrders int64    `xml:"max_orders,attr,omitempty"`
	What      string   `xml:"what,attr,omitempty"`
}

func (c *Client) ShowState(params *ShowStateInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}

type ModifyHotFolderInput struct {
	XMLName  xml.Name      `xml:"modify_hot_folder"`
	Folder   string        `xml:"folder,attr"`
	Job      *JobConf      `xml:"job,omitempty"`
	JobChain *JobChainConf `xml:"job_chain,omitempty"`
	Schedule *ScheduleConf `xml:"schedule,omitempty"`
}

func (c *Client) ModifyHotFolder(params *ModifyHotFolderInput) (*Answer, *Error) {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer, spooler.Answer.Error
}
