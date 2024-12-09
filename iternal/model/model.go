package model

import "time"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Для данных, если необходимо
}

type UserSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type User struct {
	ID           int64        `json:"id" bson:"_id,omitempty"`
	Name         string       `json:"name" bson:"name"`
	Email        string       `json:"email" bson:"email"`
	Password     string       `json:"password" bson:"password"`
	RegisteredAt time.Time    `json:"registeredAt" bson:"registeredAt"`
	Verification Verification `json:"verification" bson:"verification"`
}

type Verification struct {
	Code     string `json:"code" bson:"code"`
	Verified bool   `json:"verified" bson:"verified"`
}

type Tokens struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

type AuthenticationRequest struct {
	Uuid string `json:"uuid"`
	Ip   string
}

type RefreshSession struct {
	UserUuid     string
	RefreshToken []byte
	Ip           string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	Email        string
}
