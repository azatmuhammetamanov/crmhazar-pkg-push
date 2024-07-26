package crmhazar_pkg_push

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, options Options) (*Client, error) {
	opts := []option.ClientOption{option.WithCredentialsFile(options.JsonPath)}
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		log.Println("firebase.NewApp ", err.Error())
		return nil, err
	}

	fcmClient, err := app.Messaging(context.TODO())
	if err != nil {
		log.Println("app.Messaging ", err.Error())
		return nil, err
	}
	return &Client{FcmClient: fcmClient}, nil
}
