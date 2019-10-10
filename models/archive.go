package models

import (
	"time"
)

type Archive struct {
	Time     time.Time
	Articles SortArticles `bson:"-"`
}

type SortArchives []*Archive

func (s SortArchives) Len() int           { return len(s) }
func (s SortArchives) Less(i, j int) bool { return s[i].Time.After(s[j].Time) }
func (s SortArchives) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }



