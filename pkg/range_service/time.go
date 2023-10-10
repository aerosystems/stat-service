package RangeService

import "time"

const defaultTimeDuration = 60 * 60 * 24 // 24 hours

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func NewTimeRange(
	start time.Time,
	end time.Time,
) *TimeRange {
	return &TimeRange{
		start,
		end,
	}
}

func GetTimeRangeFromQuery(query map[string][]string) (*TimeRange, error) {
	start := time.Now().Add(-defaultTimeDuration * time.Second)
	end := time.Now()

	var err error
	if len(query["startTime"]) != 0 {
		start, err = time.Parse(time.RFC3339, query["startTime"][0])
		if err != nil {
			return nil, err
		}
	}
	if len(query["endTime"]) != 0 {
		end, err = time.Parse(time.RFC3339, query["endTime"][0])
		if err != nil {
			return nil, err
		}
	}

	return NewTimeRange(start, end), nil
}
