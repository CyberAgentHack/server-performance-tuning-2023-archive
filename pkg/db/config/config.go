package config

type DBConfig struct {
	RawDBConfig            *RawDBConfig
	SecretsManagerDBConfig *SecretsManagerDBConfig
}

type RawDBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       string `json:"dbname"`
}

type SecretsManagerDBConfig struct {
	SecretID string
}
