package dockermachine

import (
	"encoding/json"
	"fmt"

	"github.com/andrewsmedina/yati/tsuru/iaas"
	"github.com/docker/machine/drivers/virtualbox"
	"github.com/docker/machine/libmachine"
)

func init() {
	iaas.Register("docker-machine", &dmIaas{})
}

type dmIaas struct{}

func (i *dmIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	client := libmachine.NewClient("/tmp/automatic", "/tmp/automatic/certs")
	defer client.Close()
	driver := virtualbox.NewDriver("tsuru", "/tmp/automatic")
	data, err := json.Marshal(driver)
	if err != nil {
		return nil, err
	}
	host, err := client.NewHost("virtualbox", data)
	if err != nil {
		return nil, err
	}
	err = client.Create(host)
	ip, err := host.Driver.GetIP()
	if err != nil {
		return nil, err
	}
	options := host.AuthOptions()
	config := map[string]string{
		"ca":   options.CaCertPath,
		"cert": options.ClientCertPath,
		"key":  options.ClientKeyPath,
	}
	m := iaas.Machine{
		Address: fmt.Sprintf("https://%s:2376", ip),
		Iaas:    "docker-machine",
		Config:  config,
	}
	return &m, nil
}

func (i *dmIaas) DeleteMachine(m *iaas.Machine) error {
	client := libmachine.NewClient("/tmp/automatic", "/tmp/automatic/certs")
	defer client.Close()
	h, err := client.Load("tsuru")
	if err != nil {
		return err
	}
	return h.Driver.Remove()
}
