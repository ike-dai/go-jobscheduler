package jobscheduler_test

import (
	"github.com/ike-dai/go-jobscheduler/jobscheduler"
	"testing"
	"time"
)

func TestAddJobChain(t *testing.T) {
	script1 := &jobscheduler.Script{Language: "shell", Script: "echo Job1"}
	job1 := &jobscheduler.JobConf{Name: test_job_name_1, Script: script1, Order: "yes"}
	params1 := &jobscheduler.ModifyHotFolderInput{Folder: test_job_dir, Job: job1}
	client.ModifyHotFolder(params1)

	script2 := &jobscheduler.Script{Language: "shell", Script: "echo Job2"}
	job2 := &jobscheduler.JobConf{Name: test_job_name_2, Script: script2, Order: "yes"}
	params2 := &jobscheduler.ModifyHotFolderInput{Folder: test_job_dir, Job: job2}
	client.ModifyHotFolder(params2)

	job_chain_node1 := &jobscheduler.JobChainNodeConf{State: "step1", Job: test_job_name_1, NextState: "step2", ErrorState: "error"}
	job_chain_node2 := &jobscheduler.JobChainNodeConf{State: "step2", Job: test_job_name_2, NextState: "success", ErrorState: "error"}
	error_node := &jobscheduler.JobChainNodeConf{State: "error"}
	success_node := &jobscheduler.JobChainNodeConf{State: "success"}
	node_list := []*jobscheduler.JobChainNodeConf{job_chain_node1, job_chain_node2, success_node, error_node}
	job_chain := &jobscheduler.JobChainConf{Name: test_job_chain, JobChainNode: node_list}
	params3 := &jobscheduler.ModifyHotFolderInput{Folder: test_job_dir, JobChain: job_chain}
	answer, err := client.ModifyHotFolder(params3)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	t.Log(answer)
	time.Sleep(time.Second * 15) // for waiting JobScheduler process
}

func TestStartJobChain(t *testing.T) {
	params := &jobscheduler.AddOrderInput{Id: "test", JobChain: test_job_dir + "/" + test_job_chain, State: "step1"}
	_, err := client.StartJobChain(params)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
}

func TestShowJobChain(t *testing.T) {
	params := &jobscheduler.ShowJobChainInput{JobChain: test_job_dir + "/" + test_job_chain}
	_, err := client.ShowJobChain(params)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
}

func TestRemoveJobChain(t *testing.T) {
	params := &jobscheduler.RemoveJobChainInput{JobChain: test_job_dir + "/" + test_job_chain}
	_, err := client.RemoveJobChain(params)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
}
