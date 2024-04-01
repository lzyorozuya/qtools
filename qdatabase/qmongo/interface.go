package qmongo

import "go.mongodb.org/mongo-driver/mongo"

type Database interface {
	DatabaseName() string
}

type Collection interface {
	CollectionName() string
	Database
}

type Indexer interface {
	Index() []mongo.IndexModel
}
