package repository

import "github.com/ozoncp/ocp-requirement-api/internal/models"

type Repo interface {
	AddEntities(entities []models.Requirement) error
	ListEntities(limit, offset uint64) ([]models.Requirement, error)
	DescribeEntity(entityId uint64) (*models.Requirement, error)
}
