package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutSheet struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StudentID      primitive.ObjectID `bson:"studentId,omitempty" json:"studentId,omitempty"`
	PersonalID     primitive.ObjectID `bson:"personalId,omitempty" json:"personalId,omitempty"`
	Type           string             `bson:"type,omitempty" json:"type,omitempty"`
	StartTimestamp time.Time          `bson:"startTimestamp,omitempty" json:"startTimestamp,omitempty"`
	EndTimestamp   time.Time          `bson:"endTimestamp,omitempty" json:"endTimestamp,omitempty"`
	Observation    string             `bson:"observation,omitempty" json:"observation,omitempty"`
	CreatedAt      time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt      time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
