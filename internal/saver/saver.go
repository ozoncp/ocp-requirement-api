package saver

import (
	"errors"
	"github.com/ozoncp/ocp-requirement-api/internal/flusher"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"time"
)

type Saver interface {
	Save(requirement models.Requirement) error
	Close() error
	Init() error
}

type saver struct {
	buffer    []models.Requirement
	flusher   flusher.Flusher
	done      chan struct{}
	preBuffer chan models.Requirement
}

func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	return &saver{
		buffer:  make([]models.Requirement, 0, capacity),
		flusher: flusher,
	}
}

func (s *saver) Save(requirement models.Requirement) error {
	if err := s.checkReady(); err != nil {
		return err
	}
	s.preBuffer <- requirement

	return nil
}

func (s *saver) Close() error {
	if err := s.checkReady(); err != nil {
		return err
	}
	s.done <- struct{}{}
	close(s.preBuffer)
	close(s.done)
	s.preBuffer = nil
	s.done = nil

	return nil
}

func (s *saver) Init() error {
	if err := s.checkAlreadyInit(); err != nil {
		return err
	}

	ticker := time.NewTicker(time.Second)
	s.done = make(chan struct{})
	s.preBuffer = make(chan models.Requirement)

	go func() {
		for {
			select {
			case requirement := <-s.preBuffer:
				s.buffer = append(s.buffer, requirement)
			case <-s.done:
				ticker.Stop()
				s.flush()
				return
			case <-ticker.C:
				s.flush()
			}
		}
	}()

	return nil
}

func (s *saver) checkReady() error {
	if s.preBuffer == nil || s.done == nil {
		return errors.New("call init first")
	}

	return nil
}

func (s *saver) checkAlreadyInit() error {
	if s.preBuffer != nil && s.done != nil {
		return errors.New("call close first")
	}

	return nil
}

func (s *saver) flush() {
	notFlushedRequirements := s.flusher.Flush(s.buffer)
	s.buffer = append(make([]models.Requirement, 0, cap(s.buffer)), notFlushedRequirements...)
}
