package problem021

import "sort"

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
	startTime := make([]int, 0, len(lectureSchedules))
	endTime := make([]int, 0, len(lectureSchedules))
	for _, lectureSchedule := range lectureSchedules {
		startTime = append(startTime, lectureSchedule.startsAt)
		endTime = append(endTime, lectureSchedule.endsAt)
	}
	sort.Ints(startTime)
	sort.Ints(endTime)

	events := make([]event, 0, len(startTime)+len(endTime))
	for startIndex, endIndex := 0, 0; startIndex < len(startTime) || endIndex < len(endTime); {
		switch {
		case startIndex < len(startTime) && endIndex < len(endTime):
			if startTime[startIndex] <= endTime[endIndex] {
				events = append(events, event{
					category: START,
					time:     startTime[startIndex],
				})
				startIndex++
			} else {
				events = append(events, event{
					category: END,
					time:     endTime[endIndex],
				})
				endIndex++
			}
		case startIndex < len(startTime):
			events = append(events, event{
				category: START,
				time:     startTime[startIndex],
			})
			startIndex++
		case endIndex < len(endTime):
			events = append(events, event{
				category: END,
				time:     endTime[endIndex],
			})
			endIndex++
		}
	}

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
