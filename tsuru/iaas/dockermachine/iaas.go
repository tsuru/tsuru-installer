package dockermachine

import (
	"os"
	"os/exec"

	"github.com/andrewsmedina/yati/tsuru/iaas"
)

func init() {
	iaas.Register("docker-machine", &dmIaas{})
}

type dmIaas struct{}

func (i *dmIaas) createMachine() error {
	cmd := exec.Command("docker-machine", "create", "tsuru", "-d", "virtualbox")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (i *dmIaas) getIP() (string, error) {
	cmd := exec.Command("docker-machine", "ip", "tsuru")
	output, err := cmd.Output()
	return string(output), err
}

func (i *dmIaas) deleteMachine() error {
	cmd := exec.Command("docker-machine", "rm", "tsuru", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (i *dmIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	err := i.createMachine()
	if err != nil {
		return nil, err
	}
	ip, err := i.getIP()
	if err != nil {
		return nil, err
	}
	config := map[string]string{}
	m := iaas.Machine{
		Address: ip,
		Iaas:    "docker-machine",
		Config:  config,
	}
	return &m, nil
}

func (i *dmIaas) DeleteMachine(m *iaas.Machine) error {
	return i.deleteMachine()
}
