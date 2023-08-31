package services

import (
	"context"

	"github.com/dannyoka/go-server/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IItemsService interface {

}

type ItemsService struct {
	ItemsCollection *mongo.Collection
}

func NewItemsService (client *mongo.Client) IItemsService {
	collection := client.Database("mern-shopping").Collection("items")
	return &ItemsService{ItemsCollection: collection}
}

func (service *ItemsService)GetItems(ctx context.Context)([]types.Item, error){
	var items []types.Item
	cursor, err := service.ItemsCollection.Find(ctx, bson.D{}); if err != nil{
		return nil, err
	}
	cursor.All(ctx, &items)
	return items, nil
}
