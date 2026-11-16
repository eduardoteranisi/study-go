package services

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetNote(ctx context.Context, db *mongo.Database) {
	coll := db.Collection("notes")
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text to search: ")
	txt, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	txt = strings.TrimSpace(txt)

	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "title", Value: bson.D{
			{Key: "$regex", Value: txt},
			{Key: "$options", Value: "i"},
		}}},
		bson.D{{Key: "content", Value: bson.D{
			{Key: "$regex", Value: txt},
			{Key: "$options", Value: "i"},
		}}},
	}}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}

	var results []Note
	if err = cursor.All(ctx, &results); err != nil {
		fmt.Println(err)
	}
	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}
