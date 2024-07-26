package crmhazar_pkg_push

import "firebase.google.com/go/messaging"

type Client struct {
	FcmClient *messaging.Client
	Token     string
	Tokens    []string
	Title     string
	Body      string
	Image     string
}
