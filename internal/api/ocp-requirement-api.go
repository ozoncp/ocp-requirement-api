package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-requirement-api/internal/flusher"
	"github.com/ozoncp/ocp-requirement-api/internal/kafka"
	"github.com/ozoncp/ocp-requirement-api/internal/metrics"
	"github.com/ozoncp/ocp-requirement-api/internal/models"
	"github.com/ozoncp/ocp-requirement-api/internal/repository"
	requirementApiV1Description "github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type requirementApiV1 struct {
	requirementApiV1Description.UnimplementedOcpRequirementApiServer
	repo     repository.RequirementsRepo
	flush    flusher.Flusher
	producer kafka.Producer
}

func NewRequirementApiV1(repo repository.RequirementsRepo, flush flusher.Flusher, kafkaProducer kafka.Producer) requirementApiV1Description.OcpRequirementApiServer {
	return &requirementApiV1{
		repo:     repo,
		flush:    flush,
		producer: kafkaProducer,
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

	metrics.CreateCounterInc()
	err = r.producer.Send(
		kafka.CreateMessage(
			kafka.Create.String(),
			kafka.EventContent{
				Id:     resultId,
				UserId: req.UserId,
				Text:   req.Text,
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Requirement %d has been created.", resultId)

	return resp, nil
}

func (r *requirementApiV1) MultiCreateRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.MultiCreateRequirementV1Request,
) (*requirementApiV1Description.MultiCreateRequirementV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateRequirementV1")
	ctx = opentracing.ContextWithSpan(context.Background(), span)
	defer span.Finish()

	requirements := make([]models.Requirement, 0, len(req.Requirements))
	for _, rq := range req.Requirements {
		requirement := models.Requirement{
			UserId: rq.UserId,
			Text:   rq.Text,
		}
		requirements = append(requirements, requirement)
	}
	failedRequirements := r.flush.FlushWithContext(requirements, ctx)

	result := make([]*requirementApiV1Description.Requirement, 0, len(requirements))
	for _, r := range failedRequirements {
		result = append(
			result,
			&requirementApiV1Description.Requirement{
				Id:     r.Id,
				UserId: r.UserId,
				Text:   r.Text,
			})
	}
	resp := &requirementApiV1Description.MultiCreateRequirementV1Response{Requirements: result}

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

func (r *requirementApiV1) UpdateRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.UpdateRequirementV1Request,
) (*requirementApiV1Description.UpdateRequirementV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	requirement := models.Requirement{
		Id:     req.Requirements.Id,
		UserId: req.Requirements.UserId,
		Text:   req.Requirements.Text,
	}
	updated, err := r.repo.UpdateEntity(requirement)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if updated == false {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	resp := &requirementApiV1Description.UpdateRequirementV1Response{Updated: updated}

	metrics.UpdateCounterInc()
	err = r.producer.Send(
		kafka.CreateMessage(
			kafka.Update.String(),
			kafka.EventContent{
				Id:     req.Requirements.Id,
				UserId: req.Requirements.UserId,
				Text:   req.Requirements.Text,
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Printf("Requirement %d has been updated.", req.Requirements.Id)

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
	metrics.RemoveCounterInc()

	err = r.producer.Send(
		kafka.CreateMessage(
			kafka.Remove.String(),
			kafka.EventContent{
				Id: req.RequirementId,
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
