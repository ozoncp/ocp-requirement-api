package saver

import (
	"github.com/ozoncp/ocp-requirement-api/internal/flusher"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"sync"
	"time"
)

type Saver interface {
	Save(requirement models.Requirement)
	Close()
	Init()
}

type saver struct {
	buffer  []models.Requirement
	flusher flusher.Flusher
	mu      sync.Mutex
	done    chan bool
}

func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	return &saver{
		buffer:  make([]models.Requirement, 0, capacity),
		flusher: flusher,
		done:    make(chan bool),
	}
}

func (s *saver) Save(requirement models.Requirement) {
	s.mu.Lock()
	s.buffer = append(s.buffer, requirement)
	s.mu.Unlock()
}

func (s *saver) Close() {
	s.done <- true
	s.flush()
}

func (s *saver) Init() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-s.done:
				ticker.Stop()
				return
			case <-ticker.C:
				s.mu.Lock()
				s.flush()
				s.mu.Unlock()
			}
		}
	}()
}

func (s *saver) flush() {
	s.buffer = append(make([]models.Requirement, 0, cap(s.buffer)), s.flusher.Flush(s.buffer)...)
}
