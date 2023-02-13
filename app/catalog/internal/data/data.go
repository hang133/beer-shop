package data

import (
	"beer-shop/app/catalog/internal/conf"
	"beer-shop/app/catalog/internal/data/ent"
	"beer-shop/app/catalog/internal/data/ent/migrate"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewBeerRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data"))

	d := &Data{
		db:  entClient,
		log: log,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
