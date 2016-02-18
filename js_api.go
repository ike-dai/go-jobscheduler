package main

import (
	"flag"
	"github.com/ike-dai/go-jobscheduler/jobscheduler"
	"os"
)

var joc_url string

func parseOption() {
	flag.StringVar(&joc_url, "u", "http://localhost:4444/", "Please input JOC URL")
	flag.StringVar(&joc_url, "url", "http://localhost:4444/", "Please input JOC URL")

	flag.Parse()
}

func StartJob(client *jobscheduler.Client, job, at string) {
	params := &jobscheduler.StartJobInput{Job: job, At: at}
	client.StartJob(params)
}

func initClient() *jobscheduler.Client {
	return jobscheduler.NewClient(joc_url)
}

// Usage(start job command): js_api start_job "job name" -url http://localhost:4444
// Usage(start job_chain command): js_api start_job_chain "job_chain_name"
func main() {

	switch os.Args[1] {
	case "start_job":
		job := os.Args[2]
		os.Args = os.Args[2:]
		parseOption()
		client := initClient()
		StartJob(client, job, "")
	case "start_job_chain":
	case "order":
	case "schedule":
	case "process_class":
	default:
	}
}
