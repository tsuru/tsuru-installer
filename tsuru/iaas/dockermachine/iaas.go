package dockermachine

import (
	"github.com/andrewsmedina/yati/tsuru/iaas"
	"github.com/docker/machine/commands/mcndirs"
	"github.com/docker/machine/libmachine"
)

func init() {
	iaas.Register("docker-machine", &dmIaas{})
}

type dmIaas struct{}

func (i *dmIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	api := libmachine.NewClient(mcndirs.GetBaseDir())
	driverName := "virtualbox"
	driver, _ := api.NewPluginDriver(driverName, nil)
	driver.Create()
	return &iaas.Machine{}, nil
}

func (i *dmIaas) DeleteMachine(m *iaas.Machine) error {
	return nil
}
