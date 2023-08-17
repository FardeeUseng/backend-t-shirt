package configs

type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
}

type Fiber struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	Protocal string
	Username string
	Password string
	Database string
	SSLMode  string
}
