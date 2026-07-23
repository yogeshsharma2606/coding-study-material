package intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	got := Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	want := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
	got = Merge([][]int{{1, 4}, {4, 5}})
	if !reflect.DeepEqual(got, [][]int{{1, 5}}) {
		t.Errorf("got %v", got)
	}
}

func TestInsert(t *testing.T) {
	got := Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	want := [][]int{{1, 5}, {6, 9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
	got = Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	want = [][]int{{1, 2}, {3, 10}, {12, 16}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestEraseOverlapIntervals(t *testing.T) {
	if got := EraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := EraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestCanAttendMeetings(t *testing.T) {
	if CanAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}) {
		t.Error("should not attend all")
	}
	if !CanAttendMeetings([][]int{{7, 10}, {2, 4}}) {
		t.Error("should attend all")
	}
}

func TestMinMeetingRooms(t *testing.T) {
	if got := MinMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := MinMeetingRooms([][]int{{7, 10}, {2, 4}}); got != 1 {
		t.Errorf("got %d", got)
	}
}

func TestIntervalIntersection(t *testing.T) {
	a := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	b := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	got := IntervalIntersection(a, b)
	want := [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestFindMinArrowShots(t *testing.T) {
	if got := FindMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := FindMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}); got != 4 {
		t.Errorf("got %d", got)
	}
}
