package reader

type Dictionary struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
type DSocialStatus struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

type DTypeCard struct {
	Id         int64  `json:"id" bson:"id"`
	Type       string `json:"type" bson:"type"`
	Code       string `json:"code" bson:"code"`
	Web        string `json:"web" bson:"saitbibl"`
	Email      string `json:"email" bson:"emailbibl"`
	Instagram  string `json:"instagram" bson:"instagram"`
	IsTerm     bool   `json:"isTerm" bson:"srokcard"`
	Photo      []byte `json:"photo" bson:"photo"`
	PhotoTitle []byte `json:"photoTitle" bson:"phototitul"`
}
