package config

import (
	"github.com/dnevsky/restaurant-back/internal/pkg/envutil"
	"time"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

var Config struct {
	Env          string
	JwtSecret    string
	Debug        bool
	PprofEnabled bool

	PgDsn string

	AccessTokenTTL time.Duration

	HTTPConfig struct {
		WriteTimeout       time.Duration
		ReadTimeout        time.Duration
		MaxHeaderMegabytes int
		Port               string
	}

	Limiter struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}
}

func InitConfig() {
	Config.Env = envutil.GetString("ENV", "dev")
	Config.JwtSecret = envutil.GetString("JWT_SECRET", "secret")
	Config.Debug = envutil.GetBool("DEBUG", true)
	Config.PprofEnabled = envutil.GetBool("PPROF_ENABLED", true)

	Config.PgDsn = envutil.GetString("PG_DSN", "host=localhost user=user password=pgpwd4 dbname=restaurant port=54324 sslmode=disable TimeZone=Europe/Moscow")

	Config.AccessTokenTTL = envutil.GetDuration("ACCESS_TOKEN_TTL", time.Hour)

	Config.HTTPConfig.Port = envutil.GetString("HTTP_PORT", "8000")
	Config.HTTPConfig.WriteTimeout = envutil.GetDuration("HTTP_WRITE_TIMEOUT", 10*time.Second)
	Config.HTTPConfig.ReadTimeout = envutil.GetDuration("HTTP_READ_TIMEOUT", 10*time.Second)
	Config.HTTPConfig.MaxHeaderMegabytes = envutil.GetInt("HTTP_MAX_HEADER_MEGS", 1)

	Config.Limiter.RPS = envutil.GetInt("HTTP_LIMIT_RPS", 20)
	Config.Limiter.Burst = envutil.GetInt("HTTP_LIMIT_BURST", 40)
	Config.Limiter.TTL = envutil.GetDuration("HTTP_LIMIT_TTL", 10*time.Minute)

}
