package crmhazar_pkg_push

import (
	"context"
	"firebase.google.com/go/messaging"
	"log"
)

func (c *Client) SendPush(ctx context.Context, title, body, image, token string) error {
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
			"title": title,
			"body":  body,
			"image": image,
		},

		Token: token,
	}

	response, err := c.FcmClient.Send(ctx, message)
	if err != nil {
		log.Println("FcmClient.Send ", err.Error())
		return err
	}
	log.Println("Response success => ", response)
	return nil
}

func (c *Client) SendMultiPush(ctx context.Context, title, body, image string, tokens []string) error {
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
			"title": title,
			"body":  body,
			"image": image,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Tokens: tokens,
	}

	response, err := c.FcmClient.SendMulticast(ctx, message)
	if err != nil {
		log.Println("FcmClient.SendMulticast ", err.Error())
	}
	if response != nil {
		log.Println("Response success count : ", response.SuccessCount)
		log.Println("Response failure count : ", response.FailureCount)
		log.Println("Response : ", response)
	}

	return nil
}
