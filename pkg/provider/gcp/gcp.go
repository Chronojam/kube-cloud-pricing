package main

import (
	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"k8s.io/client-go/pkg/api/v1"
)

const (
	GkeLabelArch         = "beta.kubernetes.io/arch"
	GkeLabelInstanceType = "beta.kubernetes.io/instance-type"
	GkeLabelRegion       = "failure-domain.beta.kubernetes.io/region"
	GkeLabelZone         = "failure-domain.beta.kubernetes.io/zone"

	GcpPricingUrl = "https://cloudpricingcalculator.appspot.com/static/data/pricelist.json"
)

type gcpPricing struct {
	priceList map[string]interface{} `json:"gcp_price_list"`
}

type GcpProvider struct {
	//	compute *compute.Service
	//	project string
	pricing map[string]interface{}
}

func New(opts *GcpOpts) (*GcpProvider, error) {
	g := &GcpProvider{}
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		return nil, err
	}

	computeService, err := compute.New(client)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(GcpPricingUrl)
	if err != nil {
		return nil, err
	}
	var p gcpPricing

	err = json.Unmarshal(resp, &p)
	if err != nil {
		return nil, err
	}

	g.pricing = p.priceList
	//	g.project = opts.Project
	//	g.compute = computeService

	return g, nil
}

type GcpOpts struct {
	Project string
}

// GCP implementation will look for labels on the node to determine the zone
func (gcp *GcpProvider) GetResourcePrice(node *v1.Node, cpu, mem float64) (float64, error) {
	/*iType, err := gcp.ReadLabel(GkeLabelInstanceType, node)
	if err != nil {
		return 0, err
	}*/

	region, err := gcp.ReadLabel(GkeLabelRegion, node)
	if err != nil {
		return 0, err
	}

	// Yeah this is pretty inaccurate but its tough to guess the exact pricing for
	// a given instance, so we'll assume everything is of the custom type - as we've got
	// exact data for that.
	cpuPrice := gcp.pricing["CP-COMPUTEENGINE-CUSTOM-VM-CORE"].(map[string]interface{})[region].(float64)
	memPrice := gcp.pricing["CP-COMPUTEENGINE-CUSTOM-VM-RAM"].(map[string]interface{})[region].(float64)

	return cpu*cpuPrice + mem*memPrice
}

func (gcp *GcpProvider) ReadLabel(label string, node *v1.Node) (string, error) {
	val, ok := node.Metadata.Labels[label]
	if !ok {
		return "", errors.New("could not read label %s of node", GkeLabelInstanceType)
	}

	return val, nil
}
