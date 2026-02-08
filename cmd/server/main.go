package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	cli "github.com/urfave/cli/v3"

	"github.com/dxh9845/explore-ast-grep/gen/greet/v1/greetv1connect"
	"github.com/dxh9845/explore-ast-grep/internal/server"
)

func main() {
	app := &cli.Command{
		Name:  "explore-ast-grep",
		Usage: "A basic ConnectRPC server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Value:   "0.0.0.0",
				Usage:   "host to listen on",
				Sources: cli.EnvVars("HOST"),
			},
			&cli.IntFlag{
				Name:    "port",
				Value:   8080,
				Usage:   "port to listen on",
				Sources: cli.EnvVars("PORT"),
			},
		},
		Action: runServer,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runServer(ctx context.Context, cmd *cli.Command) error {
	host := cmd.String("host")
	port := cmd.Int("port")
	addr := fmt.Sprintf("%s:%d", host, port)

	mux := http.NewServeMux()

	greeter := &server.GreetServer{}
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, mux)
}
