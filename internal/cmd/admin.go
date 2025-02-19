package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
	"github.com/vinceanalytics/vince/internal/ops"
	"github.com/vinceanalytics/vince/internal/util/data"
)

var admin = &cli.Command{
	Name:  "admin",
	Usage: "Creates admin account",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "data",
			Usage:   "directory to store data",
			Sources: cli.EnvVars("VINCE_DATA"),
			Value:   "vince-data",
		},
		&cli.StringFlag{
			Name:     "name",
			Usage:    "administrator name",
			Sources:  cli.EnvVars("VINCE_ADMIN_NAME"),
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Usage:    "administrator password",
			Sources:  cli.EnvVars("VINCE_ADMIN_PASSWORD"),
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		db, err := data.Open(c.String("data"), nil)
		if err != nil {
			return err
		}
		defer db.Close()
		return ops.CreateAdmin(db, c.String("name"), c.String("password"))
	},
}
