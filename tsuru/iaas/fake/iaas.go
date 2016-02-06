package fake

import (
	"github.com/andrewsmedina/yati/tsuru/iaas"
)

type fakeIaas struct{}

func (i *fakeIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	return &iaas.Machine{}, nil
}

func (i *fakeIaas) DeleteMachine(m *iaas.Machine) error {
	return nil
}
