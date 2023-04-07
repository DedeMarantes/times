package models

import "gorm.io/gorm"

type Time struct {
	gorm.Model
	Nome     string `json:"nome"`
	Estadio  string `json:"estadio"`
	Titulos  int    `json:"titulos"`
	Fundacao int    `json:"fundacao"`
}

type Jogador struct {
	gorm.Model
	Nome   string `json:"nome"`
	Idade  int    `json:"idade"`
	TimeID uint   `json:"time_id"`
	Time   Time   `json:"time" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
