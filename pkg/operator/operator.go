package operator

import (
	"github.com/chronojam/kube-cloud-pricing/pkg/provider"
	"github.com/chronojam/kube-cloud-pricing/pkg/provider/gcp"

	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	//"github.com/chronojam/kube-cloud-pricing/pkg/provider/aws"
)

const (
	AnnotationPrefix = "alpha.billing.chronojam.co.uk/"
)

// Represents the running operator
type Operator struct {
	CloudProvider provider.CloudProvider

	informerFactory informers.SharedInformerFactory
	podInformer     coreinformers.PodInformer
	kclient         *kubernetes.Clientset

	Cost map[string]float64
}

// Represents the configuration for the operator
type Config struct {
	Provider string       `yaml:"provider"`
	GcpOpts  *gcp.GcpOpts `yaml:"gcp_options"`

	KubernetesMaster string `yaml:"master_url"`
	KubeconfigPath   string `yaml:"kube_config_path"`
}

func New(cfg *Config) (*Operator, error) {
	o := &Operator{}

	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	o.informerFactory = informers.NewSharedInformerFactory(clientset, 0)
	o.kclient = clientset
	o.podInformer = o.informerFactory.Core().V1().Pods()
	o.podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    o.podAdd,
			UpdateFunc: o.podUpdate,
			DeleteFunc: o.podDelete,
		},
	)

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

func (o *Operator) Run() {
	stopCh := make(chan struct{})
	defer close(stop)
	o.informerFactory.Start(stopCh)

	select {}
}

func (o *Operator) podAdd(obj interface{}) {
	po := obj.(*v1.Pod)
	for key, value := range po.Metadata.Annotations {
		if strings.HasPrefix(key, AnnotationPrefix) {
			groupKey := strings.Split(key, AnnotationPrefix)[1]

			nodeName := p.Spec.NodeName
			node, err := o.kclient.Core().Nodes().Get(nodeName)
			if err != nil {
				return
			}
			for _, c := range p.Spec.Containers {
				c.Resources.Requests.Cpu
				c.Resources.Requests.Mem
			}
			o.Provider.GetResourcePrice(node, cpu, memory)
		}
	}
}
func (o *Operator) podUpdate(old, new interface{}) {}
func (o *Operator) podDelete(obj interface{})      {}
