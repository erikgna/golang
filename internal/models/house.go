package models

import (
	"golang/internal/utils"
)

type house struct {
    UUID        string `json:"uuid"`
    Name        string `json:"name"`
    Price       int    `json:"price"`
    Description string `json:"description"`
    CreatedAt   utils.Date   `json:"createdAt"`
}