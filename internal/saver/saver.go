package saver

import (
	"github.com/ozoncp/ocp-requirement-api/internal/flusher"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"time"
)

type Saver interface {
	Save(requirement models.Requirement)
	Close()
	Init()
}

type saver struct {
	buffer    []models.Requirement
	flusher   flusher.Flusher
	done      chan struct{}
	preBuffer chan models.Requirement
}

func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	return &saver{
		buffer:    make([]models.Requirement, 0, capacity),
		flusher:   flusher,
		done:      make(chan struct{}),
		preBuffer: make(chan models.Requirement),
	}
}

func (s *saver) Save(requirement models.Requirement) {
	s.preBuffer <- requirement
}

func (s *saver) Close() {
	s.done <- struct{}{}
}

func (s *saver) Init() {
	ticker := time.NewTicker(time.Second)
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
}

func (s *saver) flush() {
	notFlushedRequirements := s.flusher.Flush(s.buffer)
	s.buffer = append(make([]models.Requirement, 0, cap(s.buffer)), notFlushedRequirements...)
}
