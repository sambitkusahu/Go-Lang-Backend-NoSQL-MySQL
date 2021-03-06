package videomicro

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VideoDataModel ...
type VideoDataModel struct {
	VideoID        string `json:"video_id" bson:"video_id"`
	VideoTitle     string `json:"video_title" bson:"video_title"`
	VideoThumbnail string `json:"video_thumbnail" bson:"video_thumbnail"`
	VideoType      string `json:"video_type" bson:"video_type"`
	Like           int    `json:"like" bson:"like"`
}

// CollectionModel ...
type CollectionModel struct {
	CollectionID   string           `json:"collection_id" bson:"collection_id"`
	CollectionName string           `json:"collection_name" bson:"collection_name"`
	NosOfVideo     int              `json:"nos_of_video" bson:"nos_of_video"`
	Video          []VideoDataModel `json:"video" bson:"video"`
}

// Data ...
type Data struct {
	Data []CollectionModel `json:"collection" bson:"collection"`
}

// VideoModel ...
type VideoModel struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Success bool               `json:"success" bson:"success"`
	Data    Data               `json:"data" bson:"data"`
}

// Repository ...
type Repository interface {
	Create(ctx context.Context, videomodel VideoModel) (interface{}, error)
	Get(ctx context.Context) (interface{}, error)
	Update(ctx context.Context, id string, videomodel VideoModel) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}

// ## business logic
//  The business logic in the services contain core business logic, which should not have any knowledge of endpoint or concrete transports like HTTP or gRPC, or encoding and decoding of request and response message types. This will encourage you follow a clean architecture for the Go kit based services
//  Each service method converts as an endpoint by using an adapter and exposed by using concrete transports.

//   ## endpoint
//   In Go kit, the primary messaging pattern is RPC. An endpoint represents a single RPC method. Each service method in a Go kit service converts to an endpoint to make RPC style communication between servers and clients. Each endpoint exposes the service method to outside world using Transport layer by using concrete transports like HTTP or gRPC. A single endpoint can be exposed by using multiple transports.

//   ## transport layer
//   The transport layer in Go kit is bound to concrete transports. Go kit supports various transports for serving services using HTTP, gRPC, NATS, AMQP and Thrift. Because Go kit services are just focused on implementing business logic and don’t have any knowledge about concrete transports, you can provide multiple transports for a same service.
