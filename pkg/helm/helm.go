package helm

import (
	"errors"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
//	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
)

type ChartRequest struct {
	ReleaseName string
	Namespace string
	Chart string
	Values map[string]interface{}
}

func Install(cfg *action.Configuration, request ChartRequest) error{
	client := action.NewInstall(cfg)
	client.ReleaseName = request.ReleaseName
	client.Namespace = request.Namespace

	if client.Version == "" && client.Devel {
		client.Version = ">0.0.0-0"
	}

	chartRequested, err := loader.Load(request.Chart)
	if err != nil {
		return err
	}

	_, err = client.Run(chartRequested, request.Values)
	return err
}

func Uninstall(cfg *action.Configuration, releaseName string) error{
	_, err := cfg.Releases.History(releaseName)
	if errors.Is(err, driver.ErrReleaseNotFound) {
		return nil
	}

	client := action.NewUninstall(cfg)
	client.DisableHooks = true
	_, err = client.Run(releaseName)
	if err != nil {
		return err
	}

	return nil
}
