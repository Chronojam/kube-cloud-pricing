package main

import (
	"github.com/chronojam/kube-cloud-pricing/pkg/operator"
	"github.com/chronojam/kube-cloud-pricing/pkg/provider/gcp"
	"log"
)

func main() {
	cfg := &operator.Config{
		Provider: "gcp",
		GcpOpts: &gcp.GcpOpts{
			Project: "my-project",
		},
	}
	o, err := operator.New(cfg)
	if err != nil {
		log.Fatalf(err)
	}
}
