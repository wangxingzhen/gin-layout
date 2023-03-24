package conf

import "github.com/pkg/errors"

const (
	EnvTest = "test" // 开发环境
	EnvDev  = "dev"  // 测试环境
	EnvLong = "long" // 灰度环境
	EnvProd = "prod" // 生产环境
)

// AppConfig internal conf
type AppConfig struct {
	Env string `yaml:"env"`

	RedisAPI *RedisConf `yaml:"redis_address"`

	DBAddress *MysqlConf `yaml:"db_address"`
}

// Verify ...
func (a *AppConfig) Verify() error {
	// verify env
	switch a.Env {
	case EnvTest, EnvDev, EnvLong, EnvProd:
	default:
		return errors.New("config env is not one of envTest, envDev, envLong, envProd")
	}

	return nil
}

type RedisConf struct {
	Address                string `yaml:"address"`
	Password               string `yaml:"password"`
	DialTimeoutMillisecond int    `yaml:"dial_timeout_millisecond"` // Dial timeout for establishing new connections.
	RWTimeoutMillisecond   int    `yaml:"rw_timeout_millisecond"`   // timeout for read or write.
	PoolSize               int    `yaml:"pool_size"`                // Maximum number of socket connections
}

type MysqlConf struct {
	DatabaseName string `yaml:"database_name"`
	Hostname     string `yaml:"hostname"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	ParseTime    bool   `yaml:"parse_time"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}
