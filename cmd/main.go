package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go-starter/codebase"
	"go-starter/codeshared"
	"log"
	"os"
	"os/exec"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	app := &cli.App{
		Name:        "go-starter",
		Description: "Getting Started Go Monorepo Project with Template",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Init monorepo project",
				Action: func(ctx *cli.Context) error {
					if err := codeshared.NewInitProject(ctx).Init(); err != nil {
						log.Fatal(err)
					}
					if err := codeshared.NewInitSharedPkg(ctx).Init(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:  "add-service",
				Usage: "Add new monorepo service",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Service name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "port",
						Usage:    "Service port",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "module",
						Usage:    "Module name",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					if err := codebase.NewAddService(ctx).AddService(); err != nil {
						log.Fatal(err)
					}
					if err := codebase.NewAddModule(ctx).AddModule(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:  "add-module",
				Usage: "Add new module to existing service",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "service",
						Usage:    "Service name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "module",
						Usage:    "Module name",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					if err := codebase.NewAddModule(ctx).AddModule(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:  "add-resource",
				Usage: "Add new resource / collection to existing service",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "service",
						Usage:    "Service name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "resource",
						Usage:    "Resource name",
						Required: true,
					},
					&cli.BoolFlag{
						Name:  "collection",
						Usage: "Create collection",
					},
				},
				Action: func(ctx *cli.Context) error {
					if err := codebase.NewAddResource(ctx).Add(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:  "run-services",
				Usage: "Run all monorepo services",
				Action: func(ctx *cli.Context) error {
					cmd := exec.Command("/bin/sh", "-c", "go run ./cmd/run_services.go")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					if err := cmd.Start(); err != nil {
						log.Fatal(err)
					}

					defer func() {
						cmd.Wait()
						cmd.Process.Kill()
					}()

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
