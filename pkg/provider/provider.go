package provider

import (
	"k8s.io/client-go/pkg/api/v1"
)

// Represents a cloud provider
type CloudProvider interface {
	// GetResourcePrice returns the price for a given amount
	// of cpu and memory in usd/mo
	GetResourcePrice(node *v1.Node, cpu, mem float64) (float64, error)
}
