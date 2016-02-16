package main

import (
	"flag"
	"fmt"
	"github.com/ike-dai/go-jobscheduler/jobscheduler"
)

func main() {
	var joc_url string
	flag.StringVar(&joc_url, "u", "http://localhost:4444/", "Please input JOC URL")
	flag.StringVar(&joc_url, "url", "http://localhost:4444/", "Please input JOC URL")
	flag.Parse()

	fmt.Println(joc_url)

	client := jobscheduler.NewClient(joc_url)
	/*
		fmt.Println("@@@@@@@@@@@@Modify Hot Folder@@@@@@@@@@@@@@@@")
		script := &jobscheduler.Script{Language: "shell", Script: "echo test"}
		job := &jobscheduler.JobConf{Name: "hogehoge", Script: script}
		params := &jobscheduler.ModifyHotFolderInput{Folder: "support_test", Job: job}
		answer := client.ModifyHotFolder(params)
		if answer.Ok != nil {
			fmt.Println("OK!!!")
		}
		if answer.Error != nil {
			fmt.Println("ERROR!!!")
		}
		fmt.Println(answer.Ok)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Update Job@@@@@@@@@@@@@@@@")
		job := client.ShowJob("/support_test/hogehoge")
		script := &jobscheduler.Script{Language: "shell", Script: "echo hoge hoge hoge test"}
		job.Script = script
		answer := client.UpdateJob(job, "support_test")
		//params := &jobscheduler.ModifyHotFolderInput{Folder: "support_test", Job: job}
		//answer := client.ModifyHotFolder(params)
		fmt.Println(answer)
	*/
	fmt.Println("@@@@@@@@@@@@Show Order History@@@@@@@@@@@@@@@@")
	params := &jobscheduler.ShowOrderHistoryInput{JobChain: "test/test_job_chain"}
	answer := client.ShowOrderHistory(params)
	for _, order := range answer.JobChain.OrderHistory.Order {

		fmt.Println(order.StartTime)
		fmt.Println(order.EndTime)
		fmt.Println(order.Id)
	}

}
