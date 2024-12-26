package domain

import "time"

// Block representa os elementos para criação de um bloco da blockchain
// @Description Representa os elementos para criação de um bloco da blockchain
type Block struct {
	Id        string    `json:"id" bson:"_id"`
	Message   string    `json:"message" bson:"message"`
	Author    string    `json:"author" bson:"author"`
	Timestamp int64     `json:"timestamp" bson:"timestamp"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
}
