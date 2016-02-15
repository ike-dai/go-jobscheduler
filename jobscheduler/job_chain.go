package jobscheduler

import (
	"encoding/xml"
)

type JobChains struct {
	XMLName  xml.Name   `xml:"job_chains"`
	Counter  string     `xml:"counter,attr"`
	JobChain []JobChain `xml:"job_chain,omitempty"`
}

type JobChain struct {
	XMLName           xml.Name        `xml:"job_chain"`
	Name              string          `xml:"name,attr"`
	Orders            string          `xml:"orders,attr"`
	OrdersRecoverable string          `xml:"orders_recoverable,attr"`
	Path              string          `xml:"path,attr"`
	RunningOrders     string          `xml:"running_orders,attr"`
	State             string          `xml:"state,attr"`
	FileBased         *FileBased      `xml:"file_based"`
	JobChainNodes     []*JobChainNode `xml:"job_chain_node"`
}

type FileBased struct {
	XMLName       xml.Name `xml:"file_based"`
	File          string   `xml:"file,attr"`
	LastWriteTime string   `xml:"last_write_time,attr"`
	State         string   `xml:"state,attr"`
}

type JobChainNode struct {
	XMLName    xml.Name    `xml:"job_chain_node"`
	ErrorState string      `xml:"error_state,attr"`
	Job        string      `xml:"job,attr"`
	NextState  string      `xml:"next_state,attr"`
	State      string      `xml:"state,attr"`
	OrderQueue *OrderQueue `xml:"order_queue"`
}

type JobChainConf struct {
	XMLName                  xml.Name            `xml:"job_chain"`
	Distributed              string              `xml:"distributed,attr,omitempty"`
	FileWatchingProcessClass string              `xml:"file_watching_process_class,attr,omitempty"`
	MaxOrders                string              `xml:"max_orders,attr,omitempty"`
	Name                     string              `xml:"name,attr"`
	OrderRecoverable         string              `xml:"order_recoverable,attr,omitempty"`
	ProcessClass             string              `xml:"process_class,attr,omitempty"`
	Title                    string              `xml:"title,attr,omitempty"`
	Visible                  string              `xml:"visible,attr,omitempty"`
	JobChainNode             []*JobChainNodeConf `xml:"state,omitempty"`
}

type JobChainNodeConf struct {
	XMLName       xml.Name       `xml:"job_chain_node"`
	ErrorState    string         `xml:"error_state,attr,omitempty"`
	Job           string         `xml:"job,attr,omitempty"`
	NextState     string         `xml:"next_state,attr,omitempty"`
	OnError       string         `xml:"on_error,attr,omitempty"`
	State         string         `xml:"state,attr,omitempty"`
	OnReturnCodes *OnReturnCodes `xml:"on_return_codes,omitempty"`
}

type OnReturnCodes struct {
	XMLName      xml.Name       `xml:"on_return_codes"`
	OnReturnCode []OnReturnCode `xml:"on_return_code,omitempty"`
}

type OnReturnCode struct {
	XMLName    xml.Name       `xml:"on_return_code"`
	ReturnCode string         `xml:"return_code,omitempty"`
	ToState    *ToState       `xml:"to_state,omitempty"`
	AddOrder   *AddOrderInput `xml:"add_order,omitempty"`
}

type ToState struct {
	XMLName xml.Name `xml:"to_state"`
	State   string   `xml:"state,attr,omitempty"`
}

type OrderQueue struct {
	XMLName       xml.Name `xml:"order_queue"`
	Length        string   `xml:"length,attr"`
	NextStartTime string   `xml:"next_start_time,attr"`
}

type RemoveJobChainInput struct {
	XMLName  xml.Name `xml:"remove_job_chain"`
	JobChain string   `xml:"job_chain,attr"`
}

type ShowJobChainInput struct {
	XMLName  xml.Name `xml:"show_job_chain"`
	JobChain string   `xml:"job_chain,attr"`
	What     string   `xml:"what,attr"`
}

func (c *Client) ShowJobChains() *Answer {
	params := &ShowStateInput{What: "job_chains"}
	return c.ShowState(params)
}

func (c *Client) ShowJobChain(params *ShowJobChainInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}

func (c *Client) StartJobChain(params *AddOrderInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}

func (c *Client) AddJobChain(job_chain *JobChainConf, folder string) *Answer {
	params := &ModifyHotFolderInput{Folder: folder, JobChain: job_chain}
	return c.ModifyHotFolder(params)
}

func (c *Client) RemoveJobChain(params *RemoveJobChainInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}
