package env

type Spec struct {
	DB DB
}

type DB struct {
	Username string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PWD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     uint16 `envconfig:"DB_PORT" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}
