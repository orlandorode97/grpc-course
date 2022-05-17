package main

import (
	// blogpb "github.com/orlandorode97/grpc-golang-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}
