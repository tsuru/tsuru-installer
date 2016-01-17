package iaas

type Machine struct {
	Id             string
	Iaas           string
	Status         string
	Address        string
	Port           int
	CreationParams map[string]string
}

type IaaS interface {
	CreateMachine(params map[string]string) (*Machine, error)
	DeleteMachine(m *Machine) error
}
