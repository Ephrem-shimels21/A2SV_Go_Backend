package data

import (
	"context"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var tasks = make(map[int]models.Task)
var currentID = 1

func GetTasks() ([]models.Task, error) {
	var taskList []models.Task
	cursor, err := TaskCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return taskList, err
	}
	err = cursor.All(context.Background(), &taskList)
	if err != nil {
		return taskList, err
	}

	return taskList, nil

}

func GetTask(id string) (models.Task, error) {
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return task, err
	}

	err = TaskCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)

	if err != nil {
		return task, err
	}

	return task, nil

}

func CreateTask(task models.Task) (models.Task, error) {
	task.ID = primitive.NewObjectID()
	_, err := TaskCollection.InsertOne(context.Background(), task)

	if err != nil {
		return task, err
	}

	return task, nil

}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return updatedTask, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	_, err = TaskCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return updatedTask, err
	}
	return updatedTask, nil

}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = TaskCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return nil
}
