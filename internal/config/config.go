package config

import (
	"fmt"
	"os"
	"path"

	"github.com/layer5io/meshery-adapter-library/config"
	configprovider "github.com/layer5io/meshery-adapter-library/config/provider"
	"github.com/layer5io/meshkit/utils"
)

const (
	NginxOperation = "Nginx"
	Development    = "development"
	Production     = "production"
)

var (
	configRootPath = path.Join(utils.GetHome(), ".meshery")

	kubeconfigFilename = "kubeconfig"
	kubeconfigFiletype = "yaml"
	KubeconfigPath     = path.Join(configRootPath, fmt.Sprintf("%s.%s", kubeconfigFilename, kubeconfigFiletype))
)

// New creates a new config instance
func New(provider string) (config.Handler, error) {

	// Default config
	opts := configprovider.Options{}
	environment := os.Getenv("MESHERY_ENV")
	if len(environment) < 1 {
		environment = Development
	}

	// Config environment
	switch environment {
	case Production:
		opts = ProductionConfig
	case Development:
		opts = DevelopmentConfig
	}

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}

	return nil, ErrEmptyConfig
}

func NewKubeconfigBuilder(provider string) (config.Handler, error) {

	opts := configprovider.Options{}
	environment := os.Getenv("MESHERY_ENV")
	if len(environment) < 1 {
		environment = Development
	}

	// Config environment
	switch environment {
	case Production:
		opts.ProviderConfig = productionKubeConfig
	case Development:
		opts.ProviderConfig = developmentKubeConfig
	}

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, ErrEmptyConfig
}
