package api

import (
	"context"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"github.com/ozoncp/ocp-requirement-api/internal/repository"
	requirementApiV1Description "github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type requirementApiV1 struct {
	requirementApiV1Description.UnimplementedOcpRequirementApiServer
	repo repository.RequirementsRepo
}

func NewRequirementApiV1(repo repository.RequirementsRepo) requirementApiV1Description.OcpRequirementApiServer {
	return &requirementApiV1{
		repo: repo,
	}
}

func (r *requirementApiV1) CreateRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.CreateRequirementV1Request,
) (*requirementApiV1Description.CreateRequirementV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	requirement := models.Requirement{
		UserId: req.UserId,
		Text:   req.Text,
	}

	resultId, err := r.repo.AddEntity(requirement)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &requirementApiV1Description.CreateRequirementV1Response{RequirementId: resultId}

	log.Printf("Requirement %d has been created.", resultId)

	return resp, nil
}

func (r *requirementApiV1) DescribeRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.DescribeRequirementV1Request,
) (*requirementApiV1Description.DescribeRequirementV1Response, error) {
	requirement, err := r.repo.DescribeEntity(req.RequirementId)

	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if requirement == nil {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	result := requirementApiV1Description.Requirement{
		Id:     requirement.Id,
		UserId: requirement.UserId,
		Text:   requirement.Text,
	}
	resp := &requirementApiV1Description.DescribeRequirementV1Response{Requirement: &result}

	return resp, nil
}

func (r *requirementApiV1) ListRequirementsV1(
	ctx context.Context,
	req *requirementApiV1Description.ListRequirementsV1Request,
) (*requirementApiV1Description.ListRequirementsV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	requirements, err := r.repo.ListEntities(req.Limit, req.Offset)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*requirementApiV1Description.Requirement, 0, len(requirements))
	for _, r := range requirements {
		result = append(
			result,
			&requirementApiV1Description.Requirement{
				Id:     r.Id,
				UserId: r.UserId,
				Text:   r.Text,
			})
	}
	resp := &requirementApiV1Description.ListRequirementsV1Response{Requirements: result}
	return resp, nil
}

func (r *requirementApiV1) RemoveRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.RemoveRequirementV1Request,
) (*requirementApiV1Description.RemoveRequirementV1Response, error) {
	found, err := r.repo.RemoveEntity(req.RequirementId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}
	if found == false {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	resp := &requirementApiV1Description.RemoveRequirementV1Response{Found: found}

	return resp, nil
}
