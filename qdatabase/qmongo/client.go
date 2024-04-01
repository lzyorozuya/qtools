package qmongo

import (
	"context"
	"fmt"
	"github.com/lz01wcy/qtools/qlog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
)

type Client struct {
	MongoClient *mongo.Client
}

func NewClient(info *Info, models ...Collection) (*Client, error) {
	var url string
	if info.Url == "" {
		if err := info.DecodePassword(); err != nil {
			return nil, err
		}
		if err := info.Check(); err != nil {
			return nil, err
		}

		url = fmt.Sprintf("mongodb://%s:%s@%s:%s", info.Account, info.Password, info.Host, info.Port)
	} else {
		url = info.Url
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	c := &Client{MongoClient: client}

	for _, model := range models {
		if err := c.CreateIndex(model); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// GetDatabase 获取database 结构体需要实现Database
func (c *Client) GetDatabase(d Database) *mongo.Database {
	return c.MongoClient.Database(d.DatabaseName())
}

// GetCollection 获取collection 结构体需要实现Collection
func (c *Client) GetCollection(collection Collection) *mongo.Collection {
	return c.MongoClient.Database(collection.DatabaseName()).Collection(collection.CollectionName())
}

func (c *Client) CreateIndex(collection Collection) error {
	if collection == nil {
		return fmt.Errorf("collection是nil")
	}
	indexer, ok := collection.(Indexer)
	if !ok {
		qlog.Infof("%s没有实现Indexer", reflect.TypeOf(collection).String())
		return nil
	} else {
		_, err := c.GetCollection(collection).Indexes().CreateMany(context.Background(), indexer.Index())
		if err != nil {
			return fmt.Errorf("%s创建索引失败: %s", reflect.TypeOf(indexer).String(), err)
		}
		return nil
	}
}
