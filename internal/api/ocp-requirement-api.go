package api

import (
	"context"
	requirementApiV1Description "github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api"
	"github.com/rs/zerolog/log"
)

type requirementApiV1 struct {
	requirementApiV1Description.UnimplementedOcpRequirementApiServer
}

func NewRequirementApiV1() requirementApiV1Description.OcpRequirementApiServer {
	return &requirementApiV1{}
}

func (r *requirementApiV1) CreateRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.CreateRequirementV1Request,
) (*requirementApiV1Description.CreateRequirementV1Response, error) {
	log.Printf("CreateRequirementV1")
	resp := &requirementApiV1Description.CreateRequirementV1Response{}
	return resp, nil
}

func (r *requirementApiV1) DescribeRequirementV1(
	ctx context.Context,
	req *requirementApiV1Description.DescribeRequirementV1Request,
) (*requirementApiV1Description.DescribeRequirementV1Response, error) {
	log.Printf("DescribeRequirementV1")
	resp := &requirementApiV1Description.DescribeRequirementV1Response{}
	return resp, nil
}

func (r *requirementApiV1) ListRequirementsV1(
	ctx context.Context,
	req *requirementApiV1Description.ListRequirementsV1Request,
) (*requirementApiV1Description.ListRequirementsV1Response, error) {
	log.Printf("ListRequirementsV1")
	resp := &requirementApiV1Description.ListRequirementsV1Response{}
	return resp, nil
}

func (r *requirementApiV1) RemoveRequirementsV1(
	ctx context.Context,
	req *requirementApiV1Description.RemoveRequirementV1Request,
) (*requirementApiV1Description.RemoveRequirementV1Response, error) {
	log.Printf("RemoveRequirementsV1")
	resp := &requirementApiV1Description.RemoveRequirementV1Response{}
	return resp, nil
}
