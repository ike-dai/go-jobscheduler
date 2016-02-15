package jobscheduler

import (
	"encoding/xml"
)

type OrderHistory struct {
	XMLName xml.Name `xml:"order_history"`
	Order   []*Order `xml:"order,omitempty"`
}

type Order struct {
	XMLName      xml.Name `xml:"order"`
	Created      string   `xml:"created,attr,omitempty"`
	EndTime      string   `xml:"end_time,attr,omitempty"`
	HistoryId    string   `xml:"history_id,attr,omitempty"`
	Id           string   `xml:"id,attr,omitempty"`
	InitialState string   `xml:"initial_state,attr,omitempty"`
	JobChain     string   `xml:"job_chain,attr,omitempty"`
	Order        string   `xml:"order,attr,omitempty"`
	Path         string   `xml:"path,attr,omitempty"`
	Priority     string   `xml:"priority,attr,omitempty"`
	StartTime    string   `xml:"start_time,attr,omitempty"`
	State        string   `xml:"state,attr,omitempty"`
	StateText    string   `xml:"state_text,attr,omitempty"`
	Title        string   `xml:"title,attr,omitempty"`
}

type AddOrderInput struct {
	XMLName    xml.Name `xml:"add_order"`
	At         string   `xml:"at,attr,omitempty"`
	EndState   string   `xml:"end_state,attr,omitempty"`
	Id         string   `xml:"id,attr"`
	JobChain   string   `xml:"job_chain,attr"`
	Priority   string   `xml:"priority,attr,omitempty"`
	Replace    string   `xml:"replace,attr,omitempty"`
	State      string   `xml:"state,attr,omitempty"`
	Title      string   `xml:"title,attr,omitempty"`
	WebService string   `xml:"web_service,attr,omitempty"`
	//Params *Params `xml:"params,omitempty"`
	//RunTime *RunTime `xml:"run_time,omitempty"`
	//XmlPayload *XmlPayload `xml:"xml_payload,omitempty"`
}

type ShowOrderHistoryInput struct {
	JobChain        string
	MaxOrders       string
	MaxOrderHistory string
}

func (c *Client) AddOrder(params *AddOrderInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}

func (c *Client) ShowOrderHistory(params *ShowOrderHistoryInput) *Answer {
	show_job_chain_params := &ShowJobChainInput{JobChain: params.JobChain, MaxOrders: params.MaxOrders, MaxOrderHistory: params.MaxOrderHistory, What: "order_history"}
	return c.ShowJobChain(show_job_chain_params)
}
