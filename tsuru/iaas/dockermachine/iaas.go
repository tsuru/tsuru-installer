package dockermachine

import (
	"github.com/andrewsmedina/yati/tsuru/iaas"
)

type dmIaas struct{}

func (i *dmIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	return &iaas.Machine{}, nil
}

func (i *dmIaas) DeleteMachine(m *iaas.Machine) error {
	return nil
}

