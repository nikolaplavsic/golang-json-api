package main

import "time"

type Attribute struct {
	Id        string    `json:"__id"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"createdAt"`
}

type Attributes []Attribute
