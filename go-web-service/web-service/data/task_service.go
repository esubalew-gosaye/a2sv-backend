package data

import (
	"context"
	"task-management/web-service/db"
	"task-management/web-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTaskCollection() *mongo.Collection {
	return db.GetCollection("tasks")
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := getTaskCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err = cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func AddTask(task models.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := getTaskCollection().InsertOne(ctx, task)
	return result, err
}

func FindTaskById(id string) (*models.Task, error) {
	var task models.Task

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = getTaskCollection().FindOne(ctx, bson.M{"_id": objId}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(id string, task models.Task) (*mongo.UpdateResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := getTaskCollection().UpdateOne(
		ctx,
		bson.M{"_id": objId},
		bson.M{"$set": task},
	)
	return result, err
}

func DeleteTask(id string) (*mongo.DeleteResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := getTaskCollection().DeleteOne(ctx, bson.M{"_id": objId})
	return result, err
}
