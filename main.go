package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

// Quota represents the usage of a single cloud resource against its limit.
type Quota struct {
	Resource string
	Used     int
	Limit    int
}

// Utilisation returns used/limit as a fraction (0..1).
func (q Quota) Utilisation() float64 {
	if q.Limit == 0 {
		return 0
	}
	return float64(q.Used) / float64(q.Limit)
}

// Provider returns the quotas to evaluate. Implement this against a real
// cloud API (AWS Service Quotas, Azure, GCP) to use quota-watch for real.
type Provider interface {
	Quotas(ctx context.Context) ([]Quota, error)
}

func main() {
	threshold := flag.Float64("threshold", 0.8, "utilisation threshold (0..1) to flag on")
	flag.Parse()

	var provider Provider = mockProvider{} // swap for awsProvider{} once wired
	quotas, err := provider.Quotas(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading quotas:", err)
		os.Exit(2)
	}

	over := report(os.Stdout, quotas, *threshold)
	if over > 0 {
		fmt.Printf("\n%d quota(s) over %.0f%% threshold.\n", over, *threshold*100)
		os.Exit(1) // non-zero so cron/CI can alert
	}
	fmt.Printf("\nAll quotas under %.0f%% threshold.\n", *threshold*100)
}

func report(w *os.File, quotas []Quota, threshold float64) int {
	tw := tabwriter.NewWriter(w, 0, 0, 3, ' ', 0)
	fmt.Fprintln(tw, "RESOURCE\tUSED\tLIMIT\tUTIL\t")
	over := 0
	for _, q := range quotas {
		flag := ""
		if q.Utilisation() >= threshold {
			flag = "<-- OVER THRESHOLD"
			over++
		}
		fmt.Fprintf(tw, "%s\t%d\t%d\t%.0f%%\t%s\n", q.Resource, q.Used, q.Limit, q.Utilisation()*100, flag)
	}
	tw.Flush()
	return over
}
