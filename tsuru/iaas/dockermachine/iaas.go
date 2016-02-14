package dockermachine

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/andrewsmedina/yati/tsuru/iaas"
	"github.com/docker/machine/drivers/virtualbox"
	"github.com/docker/machine/libmachine"
)

func init() {
	iaas.Register("docker-machine", &dmIaas{})
}

type dmIaas struct{}

func (i *dmIaas) createMachine() error {
	client := libmachine.NewClient("/tmp/automatic", "/tmp/automatic/certs")
	defer client.Close()
	driver := virtualbox.NewDriver("tsuru", "/tmp/automatic")
	data, err := json.Marshal(driver)
	if err != nil {
		return err
	}
	h, err := client.NewHost("virtualbox", data)
	if err != nil {
		return err
	}
	return client.Create(h)
}

func (i *dmIaas) getIP() (string, error) {
	cmd := exec.Command("docker-machine", "ip", "tsuru")
	output, err := cmd.Output()
	return string(output), err
}

func (i *dmIaas) getConfig() (map[string]string, error) {
	cmd := exec.Command("docker-machine", "config", "tsuru")
	output, err := cmd.Output()
	return map[string]string{"config": string(output)}, err
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
	config, err := i.getConfig()
	if err != nil {
		return nil, err
	}
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
