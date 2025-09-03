package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/matveevfg/AI-HR/backend/api"
	"github.com/matveevfg/AI-HR/backend/pkg/deepseek"
	aiHr "github.com/matveevfg/AI-HR/backend/service/ai-hr"
	"github.com/matveevfg/AI-HR/backend/storage/postgres"
)

func main() {
	confPath := flag.String("config", "./config.yaml", "path to config file")
	flag.Parse()

	logger := log.Logger

	conf, err := configFromYaml(confPath)
	if err != nil {
		logger.Fatal().Err(err).Str("path", *confPath).Msg("failed to load config file")
	}

	storage, err := postgres.New(
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.DBName,
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.App.Env,
	)
	if err != nil {
		logger.Fatal().Err(err).Str("path", *confPath).Msg("failed to connect to postgres")
	}

	if err := storage.Migrate(); err != nil {
		logger.Fatal().Err(err).Msg("failed to migrate postgres")
	}

	llmService, err := deepseek.New(storage, conf.LLM.Token)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to init llm")
	}

	mainService := aiHr.New(storage, llmService)

	server := api.New(mainService)

	go func() {
		if err := server.Start(conf.App.Addr); err != nil {
			logger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	log.Info().Msg("service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("shutdown http server")
	}
	if err := storage.Close(); err != nil {
		log.Error().Err(err).Msg("close db connection")
	}

	log.Info().Msg("closed")
}

func configFromYaml(path *string) (*config, error) {
	if path == nil {
		return nil, errors.New("no config path")
	}

	data, err := os.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	var c config

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

type config struct {
	App struct {
		Addr string `yaml:"addr"`
		Env  string `yaml:"env"`
	} `yaml:"app"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbName"`
	} `yaml:"postgres"`
	LLM struct {
		Token string `yaml:"token"`
	} `yaml:"llm"`
}
