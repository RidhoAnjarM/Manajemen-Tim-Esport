package models

import (
	"gorm.io/gorm"
)

// Model untuk Login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Model untuk User
type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Model untuk Team eSports
type Team struct {
	ID           uint     `json:"id" gorm:"primarykey"`
	TeamName     string   `json:"team_name" binding:"required"`
	Game         string   `json:"game" binding:"required,oneof=Dota2 CS:GO Valorant PUBGM MLBB"`
	Achievements string   `json:"achievements" binding:"omitempty,max=255"`
	Logo         string   `json:"logo" binding:"required"`
	Players      []Player `json:"players"`
}

// Model untuk Player
type Player struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	TeamID   uint   `json:"team_id"`
	Name     string `json:"name" binding:"required"`
	Position string `json:"position" binding:"required"`
	Game     string `json:"game" binding:"required,oneof=Dota2 CS:GO Valorant PUBGM MLBB"`
	Profil   string `json:"profil" binding:"required"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Team{}, &Player{})
}
