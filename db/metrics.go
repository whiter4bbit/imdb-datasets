package db

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type timerUpdate struct {
	name string
	t    time.Duration
}

type avgUpdate struct {
	name string
	v    int
}

type timerM struct {
	t time.Duration
	c int
}

type avgM struct {
	v int
	m int
	c int
}

var (
	Metrics = &metrics{}
)

func init() {
	//Metrics.start()
}

type metrics struct {
	timers   map[string]*timerM
	avg      map[string]*avgM
	timersCh chan *timerUpdate
	avgCh    chan *avgUpdate
}

func safeDiv(a, b int64) int64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func (m *metrics) start() {
	m.timers = make(map[string]*timerM)
	m.timersCh = make(chan *timerUpdate)
	m.avg = make(map[string]*avgM)
	m.avgCh = make(chan *avgUpdate)

	go func() {
		tickCh := time.Tick(10 * time.Second)

		for {
			select {
			case u := <-m.timersCh:
				existing := m.timers[u.name]
				if existing == nil {
					existing = &timerM{}
					m.timers[u.name] = existing
				}
				existing.t += u.t
				existing.c += 1
			case u := <-m.avgCh:
				existing := m.avg[u.name]
				if existing == nil {
					existing = &avgM{}
					m.avg[u.name] = existing
				}
				if existing.m < u.v {
					existing.m = u.v
				}

				existing.v += u.v
				existing.c += 1
			case <-tickCh:
				var lines []string
				for name, timer := range m.timers {
					lines = append(lines, fmt.Sprintf("%s %2.2fs (%dns/op) (count: %d)", name, timer.t.Seconds(), safeDiv(timer.t.Nanoseconds(), int64(timer.c)), timer.c))
					timer.t = 0
					timer.c = 0
				}
				for name, avg := range m.avg {
					lines = append(lines, fmt.Sprintf("%s %d (%d avg) (%d max)", name, avg.v, safeDiv(int64(avg.v), int64(avg.c)), avg.m))
					avg.c = 0
					avg.v = 0
					avg.m = 0
				}
				log.Printf(strings.Join(lines, ", "))
			}
		}
	}()
}
