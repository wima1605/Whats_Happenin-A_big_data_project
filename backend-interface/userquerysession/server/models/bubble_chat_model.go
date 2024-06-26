package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Article represents the structure of an article
type OverallArticle struct {
	ID            primitive.ObjectID     `bson:"_id"`
	DataForBubble map[string]interface{} `bson:"data_for_bubble"`
	TopXAllCat    []int                  `bson:"top_24_all_cat"`
	Top5Insights  []string               `bson:"top_5_insights"`
	Articles      []Article              `bson:"articles"`
}
