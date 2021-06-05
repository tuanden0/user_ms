package main

import (
	"log"
	"os"
	"user_ms/tools/build/protoc"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "protoc",
				Usage:  "Gen proto file",
				Action: protoc.Action,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
