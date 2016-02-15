package dockermachine

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/andrewsmedina/yati/tsuru/iaas"
	"github.com/docker/machine/drivers/virtualbox"
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/host"
)

func init() {
	iaas.Register("docker-machine", &dmIaas{})
}

type dmIaas struct{}

func (i *dmIaas) createMachine() (*host.Host, error) {
	client := libmachine.NewClient("/tmp/automatic", "/tmp/automatic/certs")
	driver := virtualbox.NewDriver("tsuru", "/tmp/automatic")
	data, err := json.Marshal(driver)
	if err != nil {
		return nil, err
	}
	h, err := client.NewHost("virtualbox", data)
	if err != nil {
		return nil, err
	}
	err = client.Create(h)
	return h, err
}

func (i *dmIaas) deleteMachine() error {
	cmd := exec.Command("docker-machine", "rm", "tsuru", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (i *dmIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	host, err := i.createMachine()
	if err != nil {
		return nil, err
	}
	config := map[string]string{}
	ip, err := host.Driver.GetIP()
	if err != nil {
		return nil, err
	}
	m := iaas.Machine{
		Address: fmt.Sprintf("https://%s:2376", ip),
		Iaas:    "docker-machine",
		Config:  config,
	}
	return &m, nil
}

func (i *dmIaas) DeleteMachine(m *iaas.Machine) error {
	return i.deleteMachine()
}
