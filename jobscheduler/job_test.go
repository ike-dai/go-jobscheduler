package jobscheduler_test

import (
	"github.com/ike-dai/go-jobscheduler/jobscheduler"
	"os"
	"testing"
	"time"
)

var client *jobscheduler.Client
var test_job_dir string
var test_job_name string
var test_job string
var test_job_name_1 string
var test_job_name_2 string
var test_job_1 string
var test_job_2 string
var test_job_chain string

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	client = jobscheduler.NewClient(os.Getenv("JS_URL"))
	test_job_dir = "test"
	test_job_name = "test_job"
	test_job = test_job_dir + "/" + test_job_name
	test_job_name_1 = "test_job_1"
	test_job_name_2 = "test_job_2"
	test_job_1 = test_job_dir + "/" + test_job_name_1
	test_job_2 = test_job_dir + "/" + test_job_name_2
	test_job_chain = "test_job_chain"
}

func TestAddJob(t *testing.T) {
	script := &jobscheduler.Script{Language: "shell", Script: "echo test_job"}
	job := &jobscheduler.JobConf{Name: test_job_name, Script: script, Order: "no"}
	params := &jobscheduler.ModifyHotFolderInput{Folder: test_job_dir, Job: job}
	answer := client.ModifyHotFolder(params)
	if answer.Ok == nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", answer.Error.Code, answer.Error.Text)
	}
	time.Sleep(time.Second * 10) // for waiting JobScheduler process
}

func TestStartJob(t *testing.T) {
	params := &jobscheduler.StartJobInput{Job: test_job}
	answer := client.StartJob(params)
	if answer.Ok == nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", answer.Error.Code, answer.Error.Text)
	}
}

func TestShowJobs(t *testing.T) {
	answer := client.ShowJobs()
	if answer.State == nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", answer.Error.Code, answer.Error.Text)
	}

	for _, job := range answer.State.Jobs.Job {
		if job.Job == test_job_name {
			return
		}
	}
	t.Errorf("Not found %s/%s \n", test_job_dir, test_job_name)
}

func TestShowJob(t *testing.T) {
	job := client.ShowJob("/" + test_job)
	if job.Name != test_job_name {
		t.Errorf("Not get correct Job: [expected: %s, actual: %s]\n", test_job_name, job.Name)
	}
}

func TestRemoveJob(t *testing.T) {
	answer := client.RemoveJob("test/test_job")
	if answer.Ok == nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", answer.Error.Code, answer.Error.Text)
	}
}
