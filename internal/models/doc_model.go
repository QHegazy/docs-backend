package models



type Document struct {
	DocumentName  string     `json:"document_name" validate:"required,max=255"`
	MongoID       string     `json:"mongo_id" validate:"required,len=24"`       

}
