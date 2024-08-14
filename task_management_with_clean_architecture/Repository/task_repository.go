package repository

import (
	"context"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   infrastructure.Database
	collection string
}

func NewTaskRepository(db infrastructure.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (taskRepo *taskRepository) CreateTask(c context.Context, task *domain.Task) (domain.Task, error) {
	collection := taskRepo.database.Collection(taskRepo.collection)

	_, err := collection.InsertOne(c, task)

	return *task, err
}

func (taskRepo *taskRepository) GetTasks(c context.Context) ([]domain.Task, error) {
	collection := taskRepo.database.Collection(taskRepo.collection)

	var tasks []domain.Task

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return tasks, err
	}

	err = cursor.All(context.Background(), &tasks)

	if err != nil {
		return tasks, err
	}

	return tasks, nil

}

func (taskRepo *taskRepository) GetTask(c context.Context, id string) (domain.Task, error) {
	collection := taskRepo.database.Collection(taskRepo.collection)

	var task domain.Task
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return task, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (taskRepo *taskRepository) UpdateTask(c context.Context, id string, updatedTask *domain.Task) (domain.Task, error) {
	collection := taskRepo.database.Collection(taskRepo.collection)
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return *updatedTask, err
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

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return *updatedTask, err
	}

	return *updatedTask, nil

}

func (taskRepo *taskRepository) DeleteTask(c context.Context, id string) error {
	collection := taskRepo.database.Collection(taskRepo.collection)

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		return err
	}

	return nil
}
