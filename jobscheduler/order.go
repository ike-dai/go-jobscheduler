package jobscheduler

import (
	"encoding/xml"
)

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

func (c *Client) AddOrder(params *AddOrderInput) *Answer {
	resp := c.CallApi(params)
	spooler := GetSpoolerFromResponseBody(resp)
	return spooler.Answer
}
