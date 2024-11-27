package repo

import (
	"testing"
	"time"
)

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCreate(t *testing.T) {

	testTable := Events{
		Event{UserID: 1, Time: time.Now().Add(time.Hour), Description: "User 1 TimeNow + 1 Hour"},
		Event{UserID: 1, Time: time.Now().Add(2 * time.Hour), Description: "User 1 TimeNow + 2 Hour"},
		Event{UserID: 1, Time: time.Now(), Description: "User 1 TimeNow"},
		Event{UserID: 1, Time: time.Now().Add(38 * time.Hour), Description: "User 1 TimeNow + 38 Hours"},
		Event{UserID: 2, Time: time.Now().Add(24 * time.Hour), Description: "User 2 TimeNow + 24 Hours"},
		Event{UserID: 2, Time: time.Now(), Description: "User 2 TimeNow"},
		Event{UserID: 3, Time: time.Now().Add(24 * time.Hour), Description: "User 3 TimeNow + 24 Hours"},
	}

	expected := []struct {
		UserID       int
		Descriptions []string
	}{
		{
			UserID:       1,
			Descriptions: []string{"User 1 TimeNow", "User 1 TimeNow + 1 Hour", "User 1 TimeNow + 2 Hour", "User 1 TimeNow + 38 Hours"},
		},
		{
			UserID:       2,
			Descriptions: []string{"User 2 TimeNow", "User 2 TimeNow + 24 Hours"},
		},
		{
			UserID:       3,
			Descriptions: []string{"User 3 TimeNow + 24 Hours"},
		},
	}

	storage := NewCash()

	var err error
	for _, test := range testTable {
		err, storage.Cash[test.UserID] = storage.Create(&test)
		if err != nil {
			t.Error(err.Error())
		}
	}

	for i := 0; i < len(expected); i++ {
		var localDescriptioins []string
		for _, ev := range storage.Cash[expected[i].UserID] {
			localDescriptioins = append(localDescriptioins, ev.Description)
		}
		if !isEqual(localDescriptioins, expected[i].Descriptions) {
			t.Errorf("Not Equal:\nhave - %v\nwant - %v", localDescriptioins, expected[i].Descriptions)
		}
	}

	//fmt.Println(storage)
}

func TestGet(t *testing.T) {

	testTableCreate := Events{
		Event{UserID: 1, Time: time.Now().Add(time.Hour), Description: "User 1 TimeNow + 1 Hour"},
		Event{UserID: 1, Time: time.Now().Add(2 * time.Hour), Description: "User 1 TimeNow + 2 Hour"},
		Event{UserID: 1, Time: time.Now(), Description: "User 1 TimeNow"},
		Event{UserID: 1, Time: time.Now().Add(38 * time.Hour), Description: "User 1 TimeNow + 38 Hours"},
		Event{UserID: 2, Time: time.Now().Add(24 * time.Hour), Description: "User 2 TimeNow + 24 Hours"},
		Event{UserID: 2, Time: time.Now(), Description: "User 2 TimeNow"},
		Event{UserID: 3, Time: time.Now().Add(24 * time.Hour), Description: "User 3 TimeNow + 24 Hours"},
	}

	testTableGet := []struct {
		UserID int
		Start  time.Time
		Period time.Duration
	}{
		{
			UserID: 1,
			Start:  time.Now(),
			Period: 24 * time.Hour,
		},
		{
			UserID: 1,
			Start:  time.Now(),
			Period: 1 * time.Hour,
		},
		{
			UserID: 1,
			Start:  time.Now(),
			Period: 7 * 24 * time.Hour,
		},
		{
			UserID: 2,
			Start:  time.Now(),
			Period: 2 * time.Hour,
		},
		{
			UserID: 3,
			Start:  time.Now(),
			Period: 12414 * time.Hour,
		},
	}

	expected := []struct {
		UserID       int
		Descriptions []string
	}{
		{
			UserID:       1,
			Descriptions: []string{"User 1 TimeNow + 1 Hour", "User 1 TimeNow + 2 Hour"},
		},
		{
			UserID:       1,
			Descriptions: []string{"User 1 TimeNow + 1 Hour"},
		},
		{
			UserID:       1,
			Descriptions: []string{"User 1 TimeNow + 1 Hour", "User 1 TimeNow + 2 Hour", "User 1 TimeNow + 38 Hours"},
		},
		{
			UserID:       2,
			Descriptions: []string{},
		},
		{
			UserID:       3,
			Descriptions: []string{"User 3 TimeNow + 24 Hours"},
		},
	}

	storage := NewCash()

	var err error
	for _, test := range testTableCreate {
		err, storage.Cash[test.UserID] = storage.Create(&test)
		if err != nil {
			t.Error(err.Error())
		}
	}

	for i, test := range testTableGet {
		events := storage.Get(test.UserID, test.Start, test.Period)
		var locDescriptions []string
		for _, ev := range events {
			locDescriptions = append(locDescriptions, ev.Description)
		}
		if !isEqual(locDescriptions, expected[i].Descriptions) {
			t.Errorf("Not Equal:\nhave - %v\nwant - %v", locDescriptions, expected[i].Descriptions)
		}
	}

	//fmt.Println(storage)
}
