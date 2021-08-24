package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
)

type Repo interface {
	AddEntities(entities []models.Requirement) error
	AddEntity(entity models.Requirement) (uint64, error)
	ListEntities(limit, offset uint64) ([]models.Requirement, error)
	DescribeEntity(entityId uint64) (*models.Requirement, error)
	RemoveEntity(entityId uint64) (bool, error)
}

type RequirementsRepo struct {
	db sqlx.DB
}

func NewRequirementsRepo(db sqlx.DB) RequirementsRepo {
	return RequirementsRepo{
		db: db,
	}
}

func (r RequirementsRepo) AddEntities(entities []models.Requirement) error {
	query := "INSERT INTO requirement(user_id, text) VALUES(:user_id, :text)"
	if _, err := r.db.NamedExec(query, entities); err != nil {
		return err
	}

	return nil
}

func (r RequirementsRepo) AddEntity(entity models.Requirement) (uint64, error) {
	var resultId uint64
	query := "INSERT INTO requirement(user_id, text) VALUES($1, $2) RETURNING id"
	err := r.db.QueryRowx(query, entity.UserId, entity.Text).Scan(&resultId)
	if err != nil {
		return resultId, err
	}
	return resultId, nil
}

func (r RequirementsRepo) ListEntities(limit, offset uint64) ([]models.Requirement, error) {
	var entities []models.Requirement
	query := "SELECT * FROM requirement ORDER BY id LIMIT $1 OFFSET $2"
	if err := r.db.Select(&entities, query, limit, offset); err != nil {
		return nil, err
	}
	return entities, nil
}

func (r RequirementsRepo) DescribeEntity(entityId uint64) (*models.Requirement, error) {
	entity := models.Requirement{}
	query := "SELECT * FROM requirement r WHERE r.id = $1"
	err := r.db.QueryRowx(query, entityId).StructScan(&entity)
	switch err {
	case nil:
		return &entity, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (r RequirementsRepo) RemoveEntity(entityId uint64) (bool, error) {
	query := "DELETE FROM requirement r WHERE r.id = $1"
	result, err := r.db.Exec(query, entityId)
	if err != nil {
		return false, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if count != 1 {
		return false, nil
	}
	return true, nil
}
