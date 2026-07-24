package main

import (
	"container/heap"
)

// Event represents arrival event
type Event struct {
	flightId     string
	timestamp    int
	gateDuration int
	assign       chan GateInfo
}

// GateInfo represents assignment result
type GateInfo struct {
	gate      int
	timestamp int
}

// --------------------
// Min heap for gates
// --------------------

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// --------------------
// Occupied gate heap
// --------------------

type OccupiedGate struct {
	releaseTime int
	gate        int
}

type OccupiedHeap []OccupiedGate

func (h OccupiedHeap) Len() int { return len(h) }

func (h OccupiedHeap) Less(i, j int) bool {
	return h[i].releaseTime < h[j].releaseTime
}

func (h OccupiedHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *OccupiedHeap) Push(x interface{}) {
	*h = append(*h, x.(OccupiedGate))
}

func (h *OccupiedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// --------------------
// Main function
// --------------------

func gateAssignment(nGates int, events <-chan Event) {

	// available gates
	available := &IntHeap{}
	heap.Init(available)

	for i := 0; i < nGates; i++ {
		heap.Push(available, i)
	}

	// occupied gates
	occupied := &OccupiedHeap{}
	heap.Init(occupied)

	// waiting queue
	waitQueue := []Event{}

	currentTime := 0

	for event := range events {

		currentTime = event.timestamp

		// free gates that are done
		for occupied.Len() > 0 && (*occupied)[0].releaseTime <= currentTime {

			freed := heap.Pop(occupied).(OccupiedGate)
			heap.Push(available, freed.gate)
		}

		// add arriving flight to queue
		waitQueue = append(waitQueue, event)

		// assign gates to waiting flights
		for len(waitQueue) > 0 && available.Len() > 0 {

			flight := waitQueue[0]
			waitQueue = waitQueue[1:]

			gate := heap.Pop(available).(int)

			assignTime := currentTime

			// send assignment
			flight.assign <- GateInfo{
				gate:      gate,
				timestamp: assignTime,
			}

			// mark gate occupied
			release := assignTime + flight.gateDuration

			heap.Push(occupied, OccupiedGate{
				releaseTime: release,
				gate:        gate,
			})
		}

		// if no gates available, process future releases
		for len(waitQueue) > 0 && occupied.Len() > 0 {

			nextFree := heap.Pop(occupied).(OccupiedGate)

			currentTime = nextFree.releaseTime

			heap.Push(available, nextFree.gate)

			flight := waitQueue[0]
			waitQueue = waitQueue[1:]

			gate := heap.Pop(available).(int)

			assignTime := currentTime

			flight.assign <- GateInfo{
				gate:      gate,
				timestamp: assignTime,
			}

			release := assignTime + flight.gateDuration

			heap.Push(occupied, OccupiedGate{
				releaseTime: release,
				gate:        gate,
			})
		}
	}
}
