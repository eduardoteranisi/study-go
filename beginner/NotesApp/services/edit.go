package services

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func EditNote(ctx context.Context, db *mongo.Database) {
	coll := db.Collection("notes")
	opts := options.Replace().SetUpsert(true)
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the note title to update: ")
	fmt.Println("(If not found, we will create a new note)")
	t, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	t = strings.TrimSpace(t)

	fmt.Println("Enter the new note title: ")
	new_t, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	new_t = strings.TrimSpace(new_t)

	fmt.Println("Enter the new note content: ")
	new_c, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	new_c = strings.TrimSpace(new_c)

	filter := bson.D{
		{Key: "title", Value: bson.D{
			{Key: "$in", Value: []string{t}},
		}},
	}

	replacement := Note{Title: new_t, Content: new_c}

	result, err := coll.ReplaceOne(ctx, filter, replacement, opts)

	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents replaced: %v\n", result.ModifiedCount)
	fmt.Printf("Documents upserted: %v | New document id: %v\n", result.UpsertedCount, result.UpsertedID)
}
