// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/google/uuid"
)

type Board struct {
	ID    uuid.UUID
	Title string
}
