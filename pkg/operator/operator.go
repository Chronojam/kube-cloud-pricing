package operator

import (
	"github.com/chronojam/kube-cloud-pricing/pkg/provider"
	"github.com/chronojam/kube-cloud-pricing/pkg/provider/gcp"
	//"github.com/chronojam/kube-cloud-pricing/pkg/provider/aws"
)

// Represents the running operator
type Operator struct {
	CloudProvider provider.CloudProvider
}

// Represents the configuration for the operator
type Config struct {
	Provider string       `yaml:"provider"`
	GcpOpts  *gcp.GcpOpts `yaml:"gcp_options"`
}

func New(cfg *Config) (*Operator, error) {
	o := &Operator{}
	switch cfg.Provider {
	case "gcp":
		gcpProvider, err := gcp.New(cfg.GcpOpts)
		if err != nil {
			log.Fatalf("error while creating gcp client %v", err)
		}
		o.Provider = gcpProvider
		//case "aws":

	}
}
