package config

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const (
	User = "Ilia"
)

func (c *Config) SetUser() {
	c.Current_user_name = User
}
