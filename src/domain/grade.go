package domain

import "time"

type GradeDTO struct {
	Id          int
	Grade       int
	StudentId   int
	SubjectName string
	PublishDate time.Time
}
