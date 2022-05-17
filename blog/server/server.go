package main

import (
	"context"
	"fmt"

	blogpb "github.com/orlandorode97/grpc-golang-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	blogpb.BlogServiceServer
	collection *mongo.Collection
}

func (s *GRPCServer) CreateBlog(ctx context.Context, r *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultGRPCTimeout)
	defer cancel()
	blog := Blog{
		AuthorID: r.GetAuthorId(),
		Title:    r.GetTitle(),
		Content:  r.GetContent(),
	}
	result, err := s.collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error: %v\n", err))
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot cast into primitive objectID: %v\n", err))
	}

	return &blogpb.CreateBlogResponse{
		Id: id.Hex(),
	}, nil
}

func (s *GRPCServer) ReadBlog(ctx context.Context, r *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultGRPCTimeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot parse object id: %v\n", err))
	}

	result := s.collection.FindOne(ctx, bson.M{"_id": id})

	var blog Blog
	if err := result.Decode(&blog); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("cannot decode bson document: %v\n", err))
	}

	return &blogpb.ReadBlogResponse{
		Id:       blog.ID.Hex(),
		AuthorId: blog.AuthorID,
		Title:    blog.Title,
		Content:  blog.Content,
	}, nil
}

func (s *GRPCServer) UpdateBlog(ctx context.Context, r *blogpb.UpdateBlogRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultGRPCTimeout)
	defer cancel()

	if _, err := s.ReadBlog(ctx, &blogpb.ReadBlogRequest{Id: r.Id}); err != nil {
		return nil, status.Errorf(codes.NotFound, "could not update blog because does not exist %v\n", err)
	}

	id, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot parse object id: %v\n", err))
	}

	filter := bson.M{"_id": id}

	updated := Blog{
		AuthorID: r.AuthorId,
		Title:    r.Title,
		Content:  r.Content,
	}

	result := s.collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": updated})
	if result.Err() != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("internal error: %v\n", result.Err()))
	}

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) DeleteBlog(ctx context.Context, r *blogpb.DeleteBlogRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultGRPCTimeout)
	defer cancel()

	if _, err := s.ReadBlog(ctx, &blogpb.ReadBlogRequest{Id: r.Id}); err != nil {
		return nil, status.Errorf(codes.NotFound, "could not delete blog because does not exist %v\n", err)
	}

	id, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot parse object id: %v\n", err))
	}

	if _, err := s.collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("internal error: %v\n", err))
	}

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) ListBlogs(empty *emptypb.Empty, stream blogpb.BlogService_ListBlogsServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultGRPCTimeout)
	defer cancel()

	blogs, err := s.collection.Find(ctx, primitive.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("internal error: %v\n", err))
	}

	defer blogs.Close(ctx)

	for blogs.Next(ctx) {
		var blog Blog
		if err := blogs.Decode(&blog); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("cannot decode bson document: %v\n", err))
		}
		if err := stream.Send(&blogpb.ListBlogResponse{
			Id:       blog.ID.Hex(),
			AuthorId: blog.AuthorID,
			Title:    blog.Title,
			Content:  blog.Content,
		}); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("cannot send back the response: %v\n", err))
		}
	}

	return nil
}
