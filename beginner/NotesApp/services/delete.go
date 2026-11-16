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

func DeleteNote(ctx context.Context, db *mongo.Database) {
	coll := db.Collection("notes")
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the note title to delete it: ")
	t, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	t = strings.TrimSpace(t)

	filter := bson.D{{Key: "title", Value: bson.D{{Key: "$in", Value: []string{t}}}}}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nNumber of documents deleted: %d", result.DeletedCount)
}
