package cache

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"strings"
)

type Config struct {
	IsClusterMode bool
	TLSEnabled    bool
	SkipVerify    bool
	Addr          string
	Password      string
	MaxRetries    int
}

func New(cfg Config) (redis.UniversalClient, error) {
	var tlsConfig *tls.Config
	var client redis.UniversalClient

	if cfg.TLSEnabled {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: cfg.TLSEnabled,
		}
	}
	if cfg.IsClusterMode {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      strings.Split(cfg.Addr, ","),
			Password:   cfg.Password,
			MaxRetries: cfg.MaxRetries,
			TLSConfig:  tlsConfig,
		})
		_, err := client.Ping(client.Context()).Result()
		if err != nil {
			log.Error().Err(err).Msg("could not connect to redis cluster client")
			return nil, err
		}
		return client, nil
	}
	client = redis.NewClient(&redis.Options{
		Addr:       cfg.Addr,
		Password:   cfg.Password,
		DB:         0,
		MaxRetries: cfg.MaxRetries,
		TLSConfig:  tlsConfig,
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		log.Error().Err(err).Msg("could not connect to redis client")
		return nil, err
	}
	return client, nil
}
