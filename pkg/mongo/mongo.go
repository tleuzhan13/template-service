package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Database    string `env:"MONGO_DB"`
	URI         string `env:"MONGO_DB_URI"`
	Username    string `env:"MONGO_USERNAME"`
	Password    string `env:"MONGO_PWD"`
	TLSFilePath string `env:"TLS_FILE_PATH"`
	TLSEnable   bool   `env:"TLS_ENABLE"`
	ReplicaSet  string `env:"MONGO_REPLICA_SET"`
}

var clientOptions *options.ClientOptions

// NewConnect creates connection to mongo and returns the mongo conn struct
func NewConnect(ctx context.Context, cfg Config) (*mongo.Database, error) {
	clientOptions = options.Client().ApplyURI(cfg.genConnectURL())
	if cfg.TLSEnable {
		tlsConfig, err := cfg.getTLSConfig(ctx)
		if err != nil {
			return nil, err
		}
		clientOptions.SetTLSConfig(tlsConfig)
	}
	if cfg.ReplicaSet != "" {
		clientOptions.SetReplicaSet(cfg.ReplicaSet)
	}
	log.Printf("Connection to mongoDB ---> %s", cfg.Database)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connection to mongoDB Error: %w ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ping connection mongoDB Error: %w ", err)
	}

	go ping(ctx, client)

	return client.Database(cfg.Database), err
}

// connectionCheck implements db reconnection
func ping(ctx context.Context, client *mongo.Client) {
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ticker.C:
			err := client.Ping(ctx, nil)
			if err != nil {
				log.Printf("Lost connection to mongoDB: %v", err.Error())
				client, _ = mongo.Connect(ctx, clientOptions)

				err = client.Ping(ctx, nil)
				if err == nil {
					log.Printf("Reconnect to mongoDB successfully%v", "!")
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			err := client.Disconnect(ctx)
			if err != nil {
				log.Printf("mongo close connection error %v", err.Error())

				return
			}

			log.Printf("mongo close connection successful")
		}
	}
}
