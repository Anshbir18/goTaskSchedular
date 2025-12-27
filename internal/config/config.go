package config

import(
    "os"
)

type Config struct {
    MySQLUser     string
    MySQLPassword string
    MySQLHost     string
    MySQLPort     string
    MySQLDB       string
}

func getEnv(key string , value string) string {
    if v:= os.Getenv(key); v != "" {
        return v
    }
    return value
}

func LoadConfig() *Config {
    return &Config{
        MySQLUser: getEnv("MYSQL_USER", "root"),
        MySQLPassword: getEnv("MYSQL_PASSWORD", "root"),
        MySQLHost: getEnv("MYSQL_HOST", "127.0.0.1"),
        MySQLPort: getEnv("MYSQL_PORT", "3306"),
        MySQLDB: getEnv("MYSQL_DB", "schedulerdb"),
    }
}

func (c *Config) DSN() string {
    return c.MySQLUser + ":" + c.MySQLPassword + "@tcp(" + c.MySQLHost + ":" + c.MySQLPort + ")/" + c.MySQLDB + "?parseTime=true"
}