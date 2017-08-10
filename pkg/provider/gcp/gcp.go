package main

import (
	compute "google.golang.org/api/compute/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type GcpProvider struct {
	compute *compute.Service
}

func New(opts *GcpOpts) (*GcpProvider, error) {

}

// GCP implementation will look for labels on the node to determine the zone and project.
func (gcp *GcpProvider) GetResourcePrice(node *v1.Node, cpu, mem float64) (float64, error) {
	node.Metadata.Labels[]
	instance, err := gcp.compute.Instances.Get(project, zone, instance).Do()
	if err != nil {
		return err
	}
}
