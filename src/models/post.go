package models

import "time"

//Represents a post made by an user
type Post struct {
	ID uint64 `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	AuthorID uint64 `json:"authorId,omitempty"`
	AuthorNick string `json:"authorNick,omitempty"`
	Likes uint64 `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`

}