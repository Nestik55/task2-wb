package repo

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type Cash struct {
	Cash map[int]Events
	mu   sync.RWMutex
}

func NewCash() *Cash {
	return &Cash{Cash: map[int]Events{}, mu: sync.RWMutex{}}
}

func (c *Cash) Create(event *Event) (error, Events) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := 0; i < len(c.Cash[event.UserID]); i++ {
		if event.Time == c.Cash[event.UserID][i].Time {
			return errors.New("repo: [create] There is already an event for this time"), c.Cash[event.UserID]
		}
	}

	c.Cash[event.UserID] = append(c.Cash[event.UserID], *event)

	sort.Slice(c.Cash[event.UserID], func(i, j int) bool {
		return c.Cash[event.UserID][i].Time.Before(c.Cash[event.UserID][j].Time)
	})

	return nil, c.Cash[event.UserID]
}

func (c *Cash) Update(newEvent *Event, oldEvent *Event) (error, Events) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := 0; i < len(c.Cash[oldEvent.UserID]); i++ {
		if c.Cash[oldEvent.UserID][i].Time == oldEvent.Time {
			c.Cash[oldEvent.UserID][i] = *newEvent

			sort.Slice(c.Cash[oldEvent.UserID], func(i, j int) bool {
				return c.Cash[oldEvent.UserID][i].Time.Before(c.Cash[oldEvent.UserID][j].Time)
			})

			return nil, c.Cash[oldEvent.UserID]
		}
	}

	return errors.New("repo: [update] oldEvent not found"), c.Cash[oldEvent.UserID]
}

func (c *Cash) Delete(event *Event) (error, Events) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := 0; i < len(c.Cash[event.UserID]); i++ {
		if event.Time == c.Cash[event.UserID][i].Time {
			c.Cash[event.UserID] = append(c.Cash[event.UserID][:i], c.Cash[event.UserID][i+1:]...)
			return nil, c.Cash[event.UserID]
		}
	}

	return errors.New("repo: [delete] Event not found"), c.Cash[event.UserID]
}

func (c *Cash) Get(userID int, start time.Time, period time.Duration) Events {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var result []Event

	end := start.Add(period)
	for i := 0; i < len(c.Cash[userID]); i++ {
		if c.Cash[userID][i].Time.After(start) && c.Cash[userID][i].Time.Before(end) {
			result = append(result, c.Cash[userID][i])
		}
	}

	return result
}
