package services

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Note struct {
	Title   string `bson:"title"`
	Content string `bson:"content"`
}

func CreateNote(ctx context.Context, db *mongo.Database) {
	coll := db.Collection("notes")
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your note title: ")
	t, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	t = strings.TrimSpace(t)

	fmt.Println("Enter your note content: ")
	c, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	c = strings.TrimSpace(c)

	note := Note{Title: t, Content: c}
	result, err := coll.InsertOne(ctx, note)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("New note created with id: %v\n", result.InsertedID)
}
