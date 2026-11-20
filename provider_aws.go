package main

import "context"

// awsProvider is a placeholder for a real AWS Service Quotas implementation.
//
// TODO: implement Quotas() using github.com/aws/aws-sdk-go-v2/service/servicequotas
//   1. load config: cfg, _ := config.LoadDefaultConfig(ctx)
//   2. client := servicequotas.NewFromConfig(cfg)
//   3. list quotas per service code (e.g. "ec2", "eks") and map AppliedQuotaValue -> Limit
//   4. fetch current usage from CloudWatch / the relevant describe APIs -> Used
type awsProvider struct{}

func (awsProvider) Quotas(ctx context.Context) ([]Quota, error) {
	return nil, errNotImplemented
}

var errNotImplemented = errorString("aws provider not implemented yet — see TODO")

type errorString string

func (e errorString) Error() string { return string(e) }
