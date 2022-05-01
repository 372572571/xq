package redisx

type Config struct {
	Enable   bool
	Addrs    []string
	Leader   map[string]string // master: master
	Password string
	Index    int
	Mode     string // alone |sentinel |cluster
}
