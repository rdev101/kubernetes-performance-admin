package kube

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

//Client used by service
var Client Kube

func init() {
	client, err := NewKube()
	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Failed to initialize Kube client"))
		os.Exit(1)
	}
	Client = client
}

//Kube interface
type Kube interface {
}

//KubeClientConfig
type kubeClientConfig struct {
	restConfig *rest.Config
}

//NewKube Constructor
func NewKube() (Kube, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get kube config")
	}
	return &kubeClientConfig{
		restConfig: cfg,
	}, nil
}
