package jobscheduler_test

import (
	"github.com/ike-dai/go-jobscheduler/jobscheduler"
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

func TestShowSchedule(t *testing.T) {
	schedule, err := client.ShowSchedule("test/sample_schedule")
	if err != nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", err.Code, err.Text)
	}

	if schedule == nil {
		t.Errorf("Not found any schedule \n")
	}
}

func TestAddSchedule(t *testing.T) {
	period1 := &jobscheduler.Period{SingleStart: "10:00"}
	period2 := &jobscheduler.Period{SingleStart: "12:00"}
	periods := []*jobscheduler.Period{period1, period2}
	schedule := &jobscheduler.ScheduleConf{
		Name:   "test_schedule",
		Period: periods,
	}
	answer := client.AddSchedule(schedule, "test")
	if answer.Ok == nil {
		t.Errorf("Got Error: [code: %s, text: %s] \n", answer.Error.Code, answer.Error.Text)
	}
}

func TestUpdateSchedule(t *testing.T) {
	schedule, _ := client.ShowSchedule("test/test_schedule")
	schedule_conf := jobscheduler.ConvertToScheduleConf(schedule)
	schedule_conf.Period[0] = &jobscheduler.Period{SingleStart: "11:11"}
	client.UpdateSchedule(schedule_conf, "test")
	after_schedule, _ := client.ShowSchedule("test/test_schedule")

	if after_schedule.Period[0].SingleStart != "11:11" {
		t.Errorf("Not much: [expected: 11:11, actual: %s] \n", after_schedule.Period[0].SingleStart)
	}
}
