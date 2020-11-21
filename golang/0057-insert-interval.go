package main

func main() {

}

func insert(intervals []Interval, newInterval Interval) []Interval {
	var result []Interval
	added := false
	for _, interval := range intervals {
		if interval.End < newInterval.Start {
			result = append(result, interval)
		} else if interval.Start > newInterval.End {
			if !added {
				added = true
				result = append(result, newInterval)
			}
			result = append(result, interval)
		} else {
			newInterval.Start = min(newInterval.Start, interval.Start)
			newInterval.End = max(newInterval.End, interval.End)
		}
	}
	if !added {
		added = true
		result = append(result, newInterval)
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
