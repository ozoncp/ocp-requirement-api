package flusher

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"github.com/ozoncp/ocp-requirement-api/internal/repository"
	"github.com/ozoncp/ocp-requirement-api/internal/utils"
	"log"
	"unsafe"
)

type Flusher interface {
	Flush(entities []models.Requirement) []models.Requirement
	FlushWithContext(entities []models.Requirement, ctx context.Context) []models.Requirement
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

	if err != nil {
		log.Print(err)
		return requirements
	}

	notFlushedRequirements := make([]models.Requirement, 0)

	for _, requirements := range requirementsBulks {
		if err := f.repository.AddEntities(requirements); err != nil {
			log.Print(err)
			notFlushedRequirements = append(notFlushedRequirements, requirements...)
		}
	}

	return notFlushedRequirements
}

func (f flusher) FlushWithContext(requirements []models.Requirement, ctx context.Context) []models.Requirement {
	requirementsBulks, err := utils.SplitRequirementsToBulks(requirements, f.chunkSize)

	if err != nil {
		log.Print(err)
		return requirements
	}

	notFlushedRequirements := make([]models.Requirement, 0)

	for _, requirements := range requirementsBulks {
		span, _ := opentracing.StartSpanFromContext(
			ctx,
			"bulk added",
		)
		span.SetTag("size", unsafe.Sizeof(requirements))
		if err := f.repository.AddEntities(requirements); err != nil {
			log.Print(err)
			notFlushedRequirements = append(notFlushedRequirements, requirements...)
		}
		span.Finish()
	}

	return notFlushedRequirements
}
