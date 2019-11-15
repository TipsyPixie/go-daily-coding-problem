package problem021

import (
	"sort"
)

type Schedule struct {
	startsAt int
	endsAt   int
}

func NewLecture(startsAt int, endsAt int) Schedule {
	return Schedule{
		startsAt: startsAt,
		endsAt:   endsAt,
	}
}

const (
	START = 0
	END   = 1
)

type event struct {
	category int
	time     int
}

func Run(lectureSchedules []Schedule) int {
	events := make([]event, 0, len(lectureSchedules)*2)
	for _, lectureSchedule := range lectureSchedules {
		events = append(events,
			event{
				category: START,
				time:     lectureSchedule.startsAt,
			}, event{
				category: END,
				time:     lectureSchedule.endsAt,
			})
	}
	sort.Slice(events, func(i, j int) bool { return events[i].time <= events[j].time })

	classroomCount, maxClassroomCount := 0, 0
	for _, event := range events {
		if event.category == START {
			classroomCount++
			if classroomCount > maxClassroomCount {
				maxClassroomCount = classroomCount
			}
		} else {
			classroomCount--
		}
	}
	return maxClassroomCount
}
