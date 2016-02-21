package jobscheduler_test

import (
	//"github.com/ike-dai/go-jobscheduler/jobscheduler"
	"testing"
)

func TestShowSchedules(t *testing.T) {
	schedules, err := client.ShowSchedules()
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}

	if schedules == nil {
		t.Errorf("Not found any schedule \n")
	}

	test_path := "/test/sample_schedule"
	for _, schedule := range schedules.Schedule {
		if schedule.Path == test_path {
			return
		}
	}
	t.Errorf("Not found schedule: %s \n", test_path)
}
