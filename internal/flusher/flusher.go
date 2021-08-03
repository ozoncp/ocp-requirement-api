package flusher

import (
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"github.com/ozoncp/ocp-requirement-api/internal/repository"
	"github.com/ozoncp/ocp-requirement-api/internal/utils"
	"log"
)

type Flusher interface {
	Flush(entities []models.Requirement) []models.Requirement
}

type flusher struct {
	chunkSize  int
	repository repository.Repo
}

func NewFlusher(chunkSize int, repository repository.Repo) Flusher {
	return flusher{
		chunkSize:  chunkSize,
		repository: repository,
	}
}

func (f flusher) Flush(requirements []models.Requirement) []models.Requirement {
	requirementsBulks, err := utils.SplitRequirementsToBulks(requirements, f.chunkSize)

	flushedRequirements := make([]models.Requirement, 0)

	if err != nil {
		log.Print(err)
		return flushedRequirements
	}

	for _, requirements := range requirementsBulks {
		if err := f.repository.AddEntities(requirements); err != nil {
			log.Print(err)
			continue
		}
		flushedRequirements = append(flushedRequirements, requirements...)
	}

	return flushedRequirements
}
