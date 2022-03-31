package database

type Server struct {
	Host string
	Port int32
	User string
	Pwd  string
}

type Config struct {
	Enable  bool
	Name    string
	Source  []Server
	Replica []Server
}
