package iaas

var iaasProviders = make(map[string]Iaas)

func Register(name string, provider Iaas) {
	iaasProviders[name] = provider
}

func Get(name string) Iaas {
	return iaasProviders[name]
}

type Machine struct {
	Id             string
	Iaas           string
	Status         string
	Address        string
	Port           int
	CreationParams map[string]string
}

type Iaas interface {
	CreateMachine(params map[string]string) (*Machine, error)
	DeleteMachine(m *Machine) error
}
