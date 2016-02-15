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
		answer := client.ShowJobs()
		for _, job := range answer.State.Jobs.Job {
			fmt.Println(job.Job)
		}

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
		answer := client.StopJob("test/standalone_job")
		fmt.Println(answer)
	*/

	/*
		fmt.Println("@@@@@@@@@@@@UnStop Job@@@@@@@@@@@@@@@@")
		answer := client.UnStopJob("support_test/hogehoge")
		fmt.Println(answer)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Suspend Job@@@@@@@@@@@@@@@@")
		answer := client.SuspendJob("support_test/hogehoge")
		fmt.Println(answer)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Remove Job@@@@@@@@@@@@@@@@")
		answer := client.RemoveJob("support_test/hogehoge")
		fmt.Println(answer)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Show State@@@@@@@@@@@@@@@@")
		//params := &jobscheduler.ShowStateInput{What: "job_chain_jobs"}
		answer := client.ShowJobChains()
		for _, job_chain := range answer.State.JobChains.JobChain {
			fmt.Println(job_chain.Name)
			for _, job_chain_node := range job_chain.JobChainNodes {
				fmt.Printf("---%s\n", job_chain_node.State)
			}
		}
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Add Order@@@@@@@@@@@@@@@@")
		params := &jobscheduler.AddOrderInput{Id: "test", JobChain: "support_test/setback_test", State: "step2"}
		answer := client.StartJobChain(params)
		fmt.Println(answer)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Add JobChain@@@@@@@@@@@@@@@@")
		fmt.Println("Add Job1")
		script1 := &jobscheduler.Script{Language: "shell", Script: "echo Job1"}
		job1 := &jobscheduler.JobConf{Name: "job1", Script: script1, Order: "yes"}
		params1 := &jobscheduler.ModifyHotFolderInput{Folder: "test", Job: job1}
		client.ModifyHotFolder(params1)

		fmt.Println("Add Job2")
		script2 := &jobscheduler.Script{Language: "shell", Script: "echo Job2"}
		job2 := &jobscheduler.JobConf{Name: "job2", Script: script2, Order: "yes"}
		params2 := &jobscheduler.ModifyHotFolderInput{Folder: "test", Job: job2}
		client.ModifyHotFolder(params2)

		fmt.Println("Add JobChain1")
		job_chain_node1 := &jobscheduler.JobChainNodeConf{State: "step1", Job: "job1", NextState: "step2", ErrorState: "error"}
		job_chain_node2 := &jobscheduler.JobChainNodeConf{State: "step2", Job: "job2", NextState: "success", ErrorState: "error"}
		error_node := &jobscheduler.JobChainNodeConf{State: "error"}
		success_node := &jobscheduler.JobChainNodeConf{State: "success"}
		node_list := []*jobscheduler.JobChainNodeConf{job_chain_node1, job_chain_node2, success_node, error_node}
		job_chain := &jobscheduler.JobChainConf{Name: "job_chain1", JobChainNode: node_list}
		params3 := &jobscheduler.ModifyHotFolderInput{Folder: "test", JobChain: job_chain}
		answer := client.ModifyHotFolder(params3)
		fmt.Println(answer)
	*/
	/*
		fmt.Println("@@@@@@@@@@@@Remove JobChain@@@@@@@@@@@@@@@@")
		params := &jobscheduler.RemoveJobChainInput{JobChain: "test/job_chain1"}
		answer := client.RemoveJobChain(params)
		fmt.Println(answer)
	*/
	fmt.Println("@@@@@@@@@@@@Show JobChain@@@@@@@@@@@@@@@@")
	params := &jobscheduler.ShowJobChainInput{JobChain: "test/test_job_chain"}
	answer := client.ShowJobChain(params)
	fmt.Println(answer.JobChain.Name)

}
