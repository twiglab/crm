package config

import (
	"context"
	"log"

	"github.com/twiglab/crm/erp/orm"
	"github.com/twiglab/crm/erp/orm/ent"
)

type Web struct {
	Addr string `yaml:"addr" mapstructure:"addr"`
}
type DB struct {
	DSN string `yaml:"dsn" mapstructure:"dsn"`
}

type App struct {
	ID  string `yaml:"id" mapstructure:"id"`
	Web Web    `yaml:"web" mapstructure:"web"`
	DB  DB     `yaml:"db" mapstructure:"db"`
}

func EntClient(dbc DB) *ent.Client {
	db, err := orm.FromURL(context.Background(), dbc.DSN)
	if err != nil {
		log.Fatal(err)
	}

	return orm.OpenClient(db)
}
