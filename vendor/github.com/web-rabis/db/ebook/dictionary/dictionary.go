package dictionary

type Dictionary struct {
	Id   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

type DState struct {
	Dictionary `bson:"_"`
}
