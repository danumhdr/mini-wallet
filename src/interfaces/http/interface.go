package http

import "go.mongodb.org/mongo-driver/mongo"

type Http struct {
	DB *mongo.Database
}

type iHttp interface {
	Launch()
}

func New(http *Http) iHttp {
	return http
}
