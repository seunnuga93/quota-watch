package main

import "context"

// mockProvider returns sample data so quota-watch runs with no credentials.
type mockProvider struct{}

func (mockProvider) Quotas(ctx context.Context) ([]Quota, error) {
	return []Quota{
		{Resource: "eks-nodes (eu-west-1)", Used: 9, Limit: 10},
		{Resource: "vpc-elastic-ips", Used: 3, Limit: 5},
		{Resource: "rds-instances", Used: 4, Limit: 40},
	}, nil
}
