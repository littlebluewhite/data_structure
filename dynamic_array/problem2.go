package main

import (
	"fmt"
	"math"
	"sort"
)

//Liz wants everyone at her compony to all hang out for a game of Minecraft, but everyone has told her
//they have TO MANY MEETING.
//She knows that there must be some time throughout the day that they are all free for some quality time.
//Write a function that given an array of everyone's meetings, returns a list of merged meetings so that
//Liz can when everyone is free for games.

//example
//meetings = [
//{StartTime: 4, EndTime: 8},
//{StartTime: 3, EndTime: 5},
//{StartTime: 0, EndTime: 1},
//{StartTime: 10, EndTime: 12},
//{StartTime: 9, EndTime: 10},
//]
//
//mergedMeetings = [
//{StartTime: 0, EndTIme: 1},
//{StartTime: 3, EndTIme: 8},
//{StartTime: 9, EndTIme: 12},
//]

type meeting struct {
	StartTime int
	EndTime   int
}

func main() {
	meetings := []meeting{
		{StartTime: 4, EndTime: 8},
		{StartTime: 3, EndTime: 5},
		{StartTime: 0, EndTime: 1},
		{StartTime: 10, EndTime: 12},
		{StartTime: 9, EndTime: 10},
	}
	fmt.Println(mergeMeetings(meetings))
}

func mergeMeetings(meetings []meeting) (result []meeting) {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].StartTime < meetings[j].StartTime
	})
	for i, m := range meetings {
		if i == 0 {
			result = append(result, m)
		} else {
			if m.StartTime <= result[len(result)-1].EndTime {
				max := int(math.Max(float64(m.EndTime), float64(result[len(result)-1].EndTime)))
				result[len(result)-1].EndTime = max
			} else {
				result = append(result, m)
			}
		}
	}
	return
}
