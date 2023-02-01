package models

type Book struct {
  ID     string  `json:"id" bson:"_id"`
  Title  string  `json:"title"`
  Author string  `json:"author"`
  Year   string  `json:"year"`
}
