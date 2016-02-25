package jobscheduler

import (
	"encoding/xml"
)

type Jobs struct {
	XMLName xml.Name `xml:"jobs"`
	Job     []*Job   `xml:"job,omitempty"`
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
	RunTime          *RunTime `xml:"run_time,omitempty"`
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
	RunTime           *RunTime `xml:"run_time,omitempty"`
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

type ModifyJobInput struct {
	XMLName xml.Name `xml:"modify_job"`
	Job     string   `xml:"job,attr"`
	Cmd     string   `xml:"cmd,attr"`
}

type ShowJobsInput struct {
	XMLName xml.Name `xml:"start_job"`
}

type ShowHistoryInput struct {
	XMLName xml.Name `xml:"show_history"`
	Job     string   `xml:"job,attr,omitempty"`
	Id      string   `xml:"id,attr,omitempty"`
	Next    string   `xml:"next,attr,omitempty"`
	Prev    string   `xml:"prev,attr,omitempty"`
	What    string   `xml:"what,attr,omitempty"`
}

func (c *Client) StartJob(params *StartJobInput) (*Ok, *Error) {
	resp, err := c.CallApi(params)
	if err != nil {
		return nil, err
	}
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer.Ok, spooler.Answer.Error
}

func (c *Client) ShowJobs() ([]*Job, *Error) {
	params := &ShowStateInput{What: "job_chain_jobs"}
	state, err := c.ShowState(params)
	return state.Jobs.Job, err
}

func (c *Client) ShowJob(job_name string) (*Job, *Error) {
	params := &ShowStateInput{What: "job_chain_jobs"}
	state, err := c.ShowState(params)
	if err != nil {
		return nil, err
	}
	for _, job := range state.Jobs.Job {
		if job.Path == "/"+job_name {
			return job, nil
		}
	}
	return nil, nil
}

func (c *Client) ShowJobsWithSource() ([]*Job, *Error) {
	params := &ShowStateInput{What: "source"}
	state, err := c.ShowState(params)
	if err != nil {
		return nil, err
	}
	return state.Jobs.Job, err
}

func (c *Client) AddJob(job *JobConf, folder string) (*Ok, *Error) {
	params := &ModifyHotFolderInput{Folder: folder, Job: job}
	return c.ModifyHotFolder(params)
}

func (c *Client) ShowJobConf(job_name string) (*JobConf, *Error) {
	jobs, err := c.ShowJobsWithSource()
	if err != nil {
		return nil, err
	}
	for _, job := range jobs {
		if job.Path == "/"+job_name {
			job.Source.Job.Name = job.Name
			return job.Source.Job, err
		}
	}
	return nil, err
}

// UpdateJob:
func (c *Client) UpdateJob(job *JobConf, job_name string) (*Ok, *Error) {
	found_job, err := c.ShowJob(job_name)
	if err != nil {
		return nil, err
	}
	if found_job == nil {
		return nil, &Error{Code: "error", Text: "Not found update target Job"}
	}
	params := &ModifyHotFolderInput{Folder: getFolderName(job_name), Job: job}
	return c.ModifyHotFolder(params)
}

func (c *Client) ModifyJob(params *ModifyJobInput) (*Ok, *Error) {
	all_cmd := []string{"stop", "unstop", "start", "wake", "end", "suspend", "continue", "remove"}
	if contains(all_cmd, params.Cmd) {
		resp, err := c.CallApi(params)
		if err != nil {
			return nil, err
		}
		spooler := GetSpoolerFromResponseBody(resp)
		return spooler.Answer.Ok, nil
	}
	return nil, &Error{Code: "error", Text: "Not supported command " + params.Cmd}
}

func (c *Client) StopJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "stop"}
	return c.ModifyJob(params)
}

func (c *Client) UnStopJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "unstop"}
	return c.ModifyJob(params)
}

func (c *Client) SuspendJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "suspend"}
	return c.ModifyJob(params)
}

func (c *Client) WakeJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "wake"}
	return c.ModifyJob(params)
}

func (c *Client) EndJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "end"}
	return c.ModifyJob(params)
}

func (c *Client) ContinueJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "continue"}
	return c.ModifyJob(params)
}

func (c *Client) RemoveJob(job_name string) (*Ok, *Error) {
	params := &ModifyJobInput{Job: job_name, Cmd: "remove"}
	return c.ModifyJob(params)
}

func (c *Client) ShowHistory(params *ShowHistoryInput) ([]*HistoryEntry, *Error) {
	resp, err := c.CallApi(params)
	if err != nil {
		return nil, err
	}
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer.History.HistoryEntry, nil
}
