package iaas

var iaasProviders = make(map[string]Iaas)

func Register(name string, provider Iaas) {
	iaasProviders[name] = provider
}

func Get(name string) Iaas {
	return iaasProviders[name]
}

type Machine struct {
	Iaas    string
	Address string
	IP      string
	Config  map[string]string
}

type Iaas interface {
	CreateMachine(params map[string]string) (*Machine, error)
	DeleteMachine(m *Machine) error
}
