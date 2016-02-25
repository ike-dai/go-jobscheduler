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
	ok, err := client.ModifyHotFolder(params)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	t.Log(ok)
	time.Sleep(time.Second * 10) // for waiting JobScheduler process
}

func TestStartJob(t *testing.T) {
	params := &jobscheduler.StartJobInput{Job: test_job}
	_, err := client.StartJob(params)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
}

func TestShowJobs(t *testing.T) {
	jobs, err := client.ShowJobs()
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}

	for _, job := range jobs {
		if job.Job == test_job_name {
			return
		}
	}
	t.Errorf("Not found %s \n", test_job)
}

func TestShowJob(t *testing.T) {
	job, err := client.ShowJob(test_job)
	if err != nil {
		t.Errorf("Not found %s \n", test_job)
	}
	t.Log(job)
}

func TestShowJobConf(t *testing.T) {
	job_conf, err := client.ShowJobConf(test_job)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	if job_conf.Name != test_job_name {
		t.Errorf("Not get correct Job: [expected: %s, actual: %s]\n", test_job_name, job_conf.Name)
	}
	t.Log(job_conf)
}

func TestShowHistory(t *testing.T) {
	params := &jobscheduler.ShowHistoryInput{Job: test_job}
	history_entry, err := client.ShowHistory(params)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	if len(history_entry) == 0 {
		t.Errorf("No history entry at job: %s \n", test_job)
	}
	t.Log(history_entry)
}

func TestStopJob(t *testing.T) {
	_, err := client.StopJob(test_job)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}

	time.Sleep(time.Second * 10) // for waiting JobScheduler process

	job, err := client.ShowJob(test_job)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	if job.State != "stopped" {
		t.Errorf("Not much state: [expect: %s, actual: %s] \n", "stopped", job.State)
	}
	t.Log(job)
}

func TestUnStopJob(t *testing.T) {
	_, err := client.UnStopJob(test_job)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	time.Sleep(time.Second * 10) // for waiting JobScheduler process

	job, err := client.ShowJob(test_job)
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	if job.State != "pending" {
		t.Errorf("Not much state: [expect: %s, actual: %s] \n", "pending", job.State)
	}
	t.Log(job)
}

func TestUpdateJob(t *testing.T) {
	job_conf, err := client.ShowJobConf("test/test_job")
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
	job_conf.Title = "Changed"

	_, err2 := client.UpdateJob(job_conf, "test/test_job")
	if err2 != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err2.Code, err2.Text)
	}
}
func TestRemoveJob(t *testing.T) {
	_, err := client.RemoveJob("test/test_job")
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}
}
