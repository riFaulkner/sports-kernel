// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type League struct {
	ID         string    `json:"id"`
	LeagueName string    `json:"leagueName"`
	LogoURL    string    `json:"logoUrl"`
	StartDate  time.Time `json:"startDate"`
	Teams      []*Team   `json:"teams"`
}

type NewUser struct {
	OwnerName string `json:"ownerName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type Team struct {
	ID          string    `json:"id"`
	FoundedDate time.Time `json:"foundedDate"`
	TeamName    string    `json:"teamName"`
}

type User struct {
	ID        string `json:"id"`
	OwnerName string `json:"ownerName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UserPreferences struct {
	ID                string    `json:"id"`
	OwnerName         string    `json:"ownerName"`
	PreferredLeagueID *string   `json:"preferredLeagueId"`
	Leagues           []*League `json:"leagues"`
}
