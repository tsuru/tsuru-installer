package dockermachine

import (
	"os/exec"

	"github.com/andrewsmedina/yati/tsuru/iaas"
)

func init() {
	iaas.Register("docker-machine", &dmIaas{})
}

type dmIaas struct{}

func (i *dmIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	cmd := exec.Command("docker-machine", "create", "tsuru", "-d", "virtualbox")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return &iaas.Machine{}, nil
}

func (i *dmIaas) DeleteMachine(m *iaas.Machine) error {
	return nil
}
