package entities

import (
	"github.com/kamva/mgm/v3"
)

type AmigurumiPattern struct {
   mgm.DefaultModel `bson:",inline"`
   Name             string `json:"name" bson:"name"`
}

func NewAmigurumiPattern(name string) *AmigurumiPattern {
   return &AmigurumiPattern{
      Name:  name,
   }
}