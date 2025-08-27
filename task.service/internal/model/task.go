package model

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus string

const (
    StatusTodo       TaskStatus = "todo"
    StatusInProgress TaskStatus = "in_progress"
    StatusDone       TaskStatus = "done"
)

type Task struct {
    ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
    Title       string              `bson:"title" json:"title"`
    Description string              `bson:"description,omitempty" json:"description,omitempty"`
    Status      TaskStatus          `bson:"status" json:"status"`
    DueDate     *time.Time          `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
    UserID      *primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
    CreatedAt   time.Time           `bson:"createdAt" json:"createdAt"`
    UpdatedAt   time.Time           `bson:"updatedAt" json:"updatedAt"`
}
