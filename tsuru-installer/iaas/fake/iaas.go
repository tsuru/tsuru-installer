package fake

import (
	"github.com/tsuru/tsuru-installer/tsuru-installer/iaas"
)

func init() {
	iaas.Register("fake", &fakeIaas{})
}

type fakeIaas struct{}

func (i *fakeIaas) CreateMachine(params map[string]string) (*iaas.Machine, error) {
	return &iaas.Machine{}, nil
}

func (i *fakeIaas) DeleteMachine(m *iaas.Machine) error {
	return nil
}
