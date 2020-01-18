package models

import "time"

type document struct {
	id       string
	title    string
	subTitle string
	tags     []string
	attr     int
	time     time.Time
	abstract []byte
	nextDoc  string
	prefDoc  string
}

