package eorder

type Department struct {
	Id   int64  `json:"id" bson:"id"`
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}
type DState struct {
	Id   int64  `json:"id" bson:"id"`
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}
