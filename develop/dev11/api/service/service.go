package service

import (
	"errors"
	"time"

	"github.com/Nestik55/develop/dev11/api/service/repo"
)

type Service struct {
	DB *repo.Cash
}

func NewService() *Service {
	return &Service{DB: repo.NewCash()}
}

func (s *Service) Create(ev *repo.Event) error {
	var err error
	err, s.DB.Cash[ev.UserID] = s.DB.Create(ev)
	return err
}

func (s *Service) Update(newEv *repo.Event, oldEv *repo.Event) error {
	var err error
	err, s.DB.Cash[oldEv.UserID] = s.DB.Update(newEv, oldEv)
	return err
}

func (s *Service) Delete(ev *repo.Event) error {
	var err error
	err, s.DB.Cash[ev.UserID] = s.DB.Delete(ev)
	return err
}

func (s *Service) Get(UserID int, start time.Time, period string) (error, repo.Events) {
	if period == "m" {
		return nil, s.DB.Get(UserID, start, 30*24*time.Hour)
	}
	if period == "w" {
		return nil, s.DB.Get(UserID, start, 7*24*time.Hour)
	}
	if period == "d" {
		return nil, s.DB.Get(UserID, start, 24*time.Hour)
	}
	return errors.New("service: [get] Incorrect period of time"), repo.Events{}
}
