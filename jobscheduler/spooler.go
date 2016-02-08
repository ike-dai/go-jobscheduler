package jobscheduler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type Spooler struct {
	XMLName xml.Name `xml:"spooler"`
	Answer  *Answer  `xml:"answer"`
}

type Answer struct {
	XMLName xml.Name `xml:"answer"`
	Ok      *Ok      `xml:"ok,omitempty"`
	Error   *Error   `xml:"ERROR,omitempty"`
	State   *State   `xml:"state,omitempty"`
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
	XMLName    xml.Name `xml:"state"`
	ConfigFile string   `xml:"config_file,attr"`
	Db         string   `xml:"db,attr"`
	DbWaiting  string   `xml:"db_waiting,attr"`
	Id         string   `xml:"id,attr"`
	LogFile    string   `xml:"log_file,attr"`
	Loop       string   `xml:"loop,attr"`
	Jobs       *Jobs    `xml:"jobs,omitempty"`
	//JobChains      JobChains      `xml:"job_chains,omitempty"`
	//ProcessClasses ProcessClasses `xml:"process_classes,omitempty"`
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
	buf, _ := xml.MarshalIndent(*params, "", " ")
	req, _ := http.NewRequest("POST", c.Url, strings.NewReader(string(buf)))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("[ERROR]: Cannot access JobScheduler API: %s \n", err)
	}
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}

type ModifyHotFolderInput struct {
	XMLName xml.Name `xml:"modify_hot_folder"`
	Folder  string   `xml:"folder,attr"`
	Job     *JobConf `xml:"job,omitempty"`
	//	JobChain JobChainConf `xml:"job_chain,omitempty"`
}

func (c *Client) ModifyHotFolder(params *ModifyHotFolderInput) *Answer {
	buf, err := xml.MarshalIndent(*params, "", " ")
	if err != nil {
		fmt.Printf("[ERROR]: Parse error: %s \n", err)
	}
	req, _ := http.NewRequest("POST", c.Url, strings.NewReader(string(buf)))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("[ERROR]: Cannot access JobScheduler API: %s \n", err)
	}
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}
