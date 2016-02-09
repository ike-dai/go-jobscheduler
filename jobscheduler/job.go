package jobscheduler

import (
	"encoding/xml"
)

type Jobs struct {
	XMLName xml.Name `xml:"jobs"`
	Job     []Job    `xml:"job,omitempty"`
}

type Job struct {
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
	Source           *Source  `xml:"source,omitempty"`
}

type Source struct {
	XMLName xml.Name `xml:"source"`
	Job     *JobConf `xml:"job,omitempty"`
}

type JobConf struct {
	XMLName           xml.Name `xml:"job"`
	Enabled           string   `xml:"enabled,attr,omitempty"`
	ForceIdleTimeout  string   `xml:"force_idle_timeout,attr,omitempty"`
	IdleTimeout       string   `xml:"idle_timeout,attr,omitempty"`
	IgnoreSignals     string   `xml:"ignore_signals,attr,omitempty"`
	JavaOptions       string   `xml:"java_options,attr,omitempty"`
	MinTasks          string   `xml:"min_tasks,attr,omitempty"`
	Name              string   `xml:"name,attr,omitempty"`
	Order             string   `xml:"order,attr,omitempty"`
	Priority          string   `xml:"priority,attr,omitempty"`
	ProcessClass      string   `xml:"process_class,attr,omitempty"`
	Replace           string   `xml:"replace,attr,omitempty"`
	SpoolerId         string   `xml:"spooler_id,attr,omitempty"`
	StopOnError       string   `xml:"stop_on_error,attr,omitempty"`
	Tasks             string   `xml:"tasks,attr,omitempty"`
	Temporary         string   `xml:"temporary,attr,omitempty"`
	Timeout           string   `xml:"timeout,attr,omitempty"`
	Title             string   `xml:"title,attr,omitempty"`
	Visible           string   `xml:"visible,attr,omitempty"`
	WarnIfLongerThan  string   `xml:"warn_if_longer_than,attr,omitempty"`
	WarnIfShorterThan string   `xml:"warn_if_shorter_than,attr,omitempty"`
	Script            *Script  `xml:"script"`
}

type StartJobInput struct {
	XMLName     xml.Name `xml:"start_job"`
	After       int      `xml:"after,attr,omitempty"`
	At          string   `xml:"at,attr,omitempty"`
	Force       bool     `xml:"force,attr,omitempty"`
	Job         string   `xml:"job,attr"` // Job is required to execute start_job
	Name        string   `xml:"name,attr,omitempty"`
	Web_service string   `xml:"web_service,attr,omitempty"`
	//Environment *Environment `xml:"environment,omitempty"` //omitempty valid to struct pointer
	//Params      *Params      `xml:"params,omitempty"`
}
type ShowJobsInput struct {
	XMLName xml.Name `xml:"start_job"`
}

func (c *Client) StartJob(params *StartJobInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}

func (c *Client) ShowJobs() *Answer {
	params := &ShowStateInput{What: "job_chain_jobs"}
	return c.ShowState(params)
}

func (c *Client) ShowJobsWithSource() *Answer {
	params := &ShowStateInput{What: "source"}
	return c.ShowState(params)
}

func (c *Client) AddJob(job *JobConf, folder string) *Answer {
	params := &ModifyHotFolderInput{Folder: folder, Job: job}
	return c.ModifyHotFolder(params)
}

func (c *Client) ShowJob(path_name string) *JobConf {
	answer := c.ShowJobsWithSource()
	for _, job := range answer.State.Jobs.Job {
		if job.Path == path_name {
			job.Source.Job.Name = job.Name
			return job.Source.Job
		}
	}
	return nil
}
