package jobscheduler

import (
	//	"encoding/xml"
	//	"fmt"
	//	"log"
	//"io/ioutil"
	"net/http"
	//"strings"
)

type Client struct {
	Url string
	*http.Client
}

func NewClient(url string) *Client {
	return &Client{url, http.DefaultClient}
}

/*
func (c *Client) Jobs() Jobs {
	command := ShowState{What: "job_chain_jobs"}
	spooler := command.Execute(c)
	return spooler.Answer.State.Jobs
}

//--------------this struct for API response----------------
type Spooler struct {
	XMLName xml.Name `xml:"spooler"`
	Answer  Answer   `xml:"answer"`
}

type Answer struct {
	XMLName xml.Name `xml:"answer"`
	Ok      Ok       `xml:"ok"`
	Error   Error    `xml:"ERROR"`
	State   State    `xml:"state"`
}
type Ok struct {
	XMLName xml.Name `xml:"ok"`
	Task    Task     `xml:"task"`
}
type Error struct {
	XMLName xml.Name `xml:"ERROR"`
	Code    string   `xml:"code,attr"`
	Text    string   `xml:"text,attr"`
	Time    string   `xml:"time,attr"`
}

type State struct {
	XMLName xml.Name `xml:"state"`
	Jobs    Jobs     `xml:"jobs"`
}

type Jobs struct {
	XMLName xml.Name `xml:"jobs"`
	Job     []Job    `xml:"job"`
}

type Job struct {
	*Client
	XMLName          xml.Name `xml:"job"`
	AllSteps         string   `xml:"all_steps,attr"`
	AllTasks         string   `xml:"all_tasks,attr"`
	Enabled          string   `xml:"enabled,attr"`
	InPeriod         string   `xml:"in_period,attr"`
	Job              string   `xml:"job,attr"`
	JobChainPriority string   `xml:"job_chain_priority,attr"`
	LogFile          string   `xml:"log_file,attr"`
	Name             string   `xml:"name,attr"`
	Order            string   `xml:"order,attr"`
	Path             string   `xml:"path,attr"`
	State            string   `xml:"state,attr"`
	Tasks            string   `xml:"tasks,attr"`
}

func (job *Job) Start() {
	log.Printf("[INFO]: %s", job.Url)
}

type Task struct {
	XMLName     xml.Name `xml:"task"`
	Enqueued    string   `xml:"enqueued,attr"`
	Force_start string   `xml:"force_start,attr"`
	Id          string   `xml:"id,attr"`
	Job         string   `xml:"job,attr"`
}

type Environment struct {
	XMLName  xml.Name   `xml:"environment,omitempty"`
	Variable []Variable `xml:"variable,omitempty"`
}

type Variable struct {
	XMLName xml.Name `xml:"variable"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:"value,attr,omitempty"`
}

type Params struct {
	XMLName xml.Name `xml:"params"`
	Param   []Param  `xml:"param,omitempty"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:"value,attr,omitempty"`
}

//----------------

func make_start_job_xml_body(job_name string, after int, at string) string {
	fmt.Println(after)
	fmt.Println(at)
	command := StartJob{Job: job_name, After: after}
	buf, _ := xml.MarshalIndent(command, "", " ")
	fmt.Println(string(buf))
	return string(buf)
}
*/
