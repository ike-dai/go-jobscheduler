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
	XMLName           xml.Name       `xml:"job_chain"`
	Name              string         `xml:"name,attr"`
	Orders            string         `xml:"orders,attr"`
	OrdersRecoverable string         `xml:"orders_recoverable,attr"`
	Path              string         `xml:"path,attr"`
	RunningOrders     string         `xml:"running_orders,attr"`
	State             string         `xml:"state,attr"`
	FileBased         *FileBased     `xml:"file_based"`
	JobChainNodes     []JobChainNode `xml:"job_chain_node"`
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

type OrderQueue struct {
	XMLName       xml.Name `xml:"order_queue"`
	Length        string   `xml:"length,attr"`
	NextStartTime string   `xml:"next_start_time,attr"`
}

func (c *Client) ShowJobChains() *Answer {
	params := &ShowStateInput{What: "job_chains"}
	return c.ShowState(params)
}
