package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/graphql-services/graphql-configurator/gen"
	"github.com/graphql-services/graphql-configurator/src"

	"github.com/aws/aws-xray-sdk-go/xray"
)

func main() {
	xray.Configure(xray.Config{
		LogLevel:       "info", // default
		ServiceVersion: "1.2.3",
	})
	db := gen.NewDBFromEnvVars()

	eventController, err := gen.NewEventController()
	if err != nil {
		panic(err)
	}

	handler := gen.GetHTTPServeMux(src.New(db, &eventController), db, src.GetMigrations(db))
	algnhsa.ListenAndServe(handler, &algnhsa.Options{
		UseProxyPath: true,
	})
}
