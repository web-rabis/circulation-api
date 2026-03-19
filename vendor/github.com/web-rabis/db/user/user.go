package user

type User struct {
	Id         int64      `json:"id" bson:"id"`
	Name       string     `json:"name" bson:"name"`
	Username   string     `json:"username" bson:"username"`
	Password   string     `json:"password" bson:"password"`
	Email      string     `json:"email" bson:"email"`
	State      string     `json:"state" bson:"state"`
	Department Department `json:"department" bson:"depart_id"`
}
type Department struct {
	Id   int64  `json:"id" bson:"id"`
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}
