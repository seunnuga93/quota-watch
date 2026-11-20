# quota-watch

A small Go CLI that checks **cloud service limits / quotas** and warns when any resource is close to its ceiling — so you find out *before* a traffic spike can't provision new nodes, not during the incident.

> I built this after a real production incident: a traffic spike tried to scale a cluster past a regional node quota, nodes couldn't be provisioned, and a user-facing service stalled while every dashboard looked "green". `quota-watch` is the proactive check that would have paged us a day earlier.

## What it does
- Reads a set of quota usages (current vs limit)
- Flags anything over a configurable utilisation threshold (default 80%)
- Prints a clear report and **exits non-zero** when something is over threshold — so it drops straight into a cron job or CI pipeline and can trigger an alert

## Usage
```bash
go run . --threshold 0.8
# or build it
make build
./bin/quota-watch --threshold 0.75
```

Example output:
```
RESOURCE                     USED   LIMIT   UTIL
eks-nodes (eu-west-1)        9      10      90%   <-- OVER THRESHOLD
vpc-elastic-ips              3      5       60%
rds-instances                4      40      10%

1 quota over 80% threshold.
```

## How it's wired
`quota-watch` is built around a small `Provider` interface:
```go
type Provider interface {
    Quotas(ctx context.Context) ([]Quota, error)
}
```
The repo ships with a **mock provider** so it runs with zero cloud credentials. To use it for real, implement `Provider` against the **AWS Service Quotas API** (or Azure/GCP equivalents) — the `TODO` in `provider_aws.go` marks exactly where.

## Why Go
It's a single static binary with no runtime deps — ideal for a cron job, a container sidecar, or a CI step.

## Roadmap
- Real AWS Service Quotas provider
- Slack / PagerDuty output
- Per-resource thresholds

## License
MIT
