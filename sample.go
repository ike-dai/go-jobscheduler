package main

import (
	"./jobscheduler"
	"flag"
	"fmt"
)

func main() {
	var joc_url string
	flag.StringVar(&joc_url, "u", "http://localhost:4444/", "Please input JOC URL")
	flag.StringVar(&joc_url, "url", "http://localhost:4444/", "Please input JOC URL")

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
		fmt.Println("@@@@@@@@@@@@Start job@@@@@@@@@@@@@@@@")
		params := &jobscheduler.StartJobInput{Job: "support_test/sample_1_local_standalone_job"}
		fmt.Println(*params)
		answer := client.StartJob(params)
		fmt.Printf("--%T--\n", *answer)
		fmt.Println(answer.Ok.Task.Job)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Show State@@@@@@@@@@@@@@@@")
		params2 := &jobscheduler.ShowStateInput{What: "source"}
		answer = client.ShowState(params2)
		for _, job := range answer.State.Jobs.Job {
			fmt.Println(job.Job)
		}
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Show Jobs@@@@@@@@@@@@@@@@")
		answer = client.ShowJobs()
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Show Job@@@@@@@@@@@@@@@@")
		job := client.ShowJob("/support_test/hogehoge")
		fmt.Println(job)
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
	/*
		fmt.Println("@@@@@@@@@@@@Modify Job@@@@@@@@@@@@@@@@")
		params := &jobscheduler.ModifyJobInput{Job: "support_test/hogehoge", Cmd: "suspend"}
		answer := client.ModifyJob(params)
		fmt.Println(answer)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Stop Job@@@@@@@@@@@@@@@@")
		answer := client.StopJob("support_test/hogehoge")
		fmt.Println(answer)
	*/

	fmt.Println("@@@@@@@@@@@@UnStop Job@@@@@@@@@@@@@@@@")
	answer := client.UnStopJob("support_test/hogehoge")
	fmt.Println(answer)

	/*
		fmt.Println("@@@@@@@@@@@@Suspend Job@@@@@@@@@@@@@@@@")
		answer := client.SuspendJob("support_test/hogehoge")
		fmt.Println(answer)
	*/
}
