package main

import (
	"fmt"
	"github.com/strideynet/bsky-furry-feed/bluesky"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

type environment struct {
	dbURL string
}

var environments = map[string]environment{
	"local": {
		dbURL: "postgres://bff:bff@localhost:5432/bff?sslmode=disable",
	},
	"production": {
		// Requires noah has run
		// ./cloud-sql-proxy --auto-iam-authn bsky-furry-feed:us-east1:main-us-east -p 15432
		// TODO: Support detecting user email ??
		dbURL: "postgres://noah@noahstride.co.uk@localhost:15432/bff?sslmode=disable",
	},
}

// TODO: Have a `login` and `logout` command that persists auth state to disk.
var username = os.Getenv("BSKY_USERNAME")
var password = os.Getenv("BSKY_PASSWORD")

func main() {
	log, _ := zap.NewDevelopment()
	var env = &environment{}
	app := &cli.App{
		Name:  "bffctl",
		Usage: "The swiss army knife of any BFF operator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "environment",
				Aliases: []string{
					"e",
				},
				Required: true,
				Action: func(c *cli.Context, s string) error {
					v, ok := environments[s]
					if !ok {
						return fmt.Errorf("unrecognized environment: %s", s)
					}
					log.Info("configured environment", zap.String("env", s))
					*env = v
					return nil
				},
			},
		},
		Commands: []*cli.Command{
			dbCmd(log, env),
			findDIDCmd(log),
			followCmd(log, env),
			scanCmd(log, env),
			queueCmd(log, env),
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func findDIDCmd(log *zap.Logger) *cli.Command {
	handle := ""
	return &cli.Command{
		Name: "find-did",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "handle",
				Usage:       "Find DID for handle",
				Destination: &handle,
				Required:    true,
			},
		},
		Action: func(cctx *cli.Context) error {
			client := bluesky.NewClient(nil)
			did, err := client.ResolveHandle(cctx.Context, handle)
			if err != nil {
				return err
			}
			log.Info("found did", zap.String("did", did.Did))
			return nil
		},
	}
}
