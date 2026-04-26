package model

import "github.com/web-rabis/order-client/protobuf"

type User struct {
	Id         int64       `json:"id" bson:"id"`
	Name       string      `json:"name" bson:"name"`
	Username   string      `json:"username" bson:"username"`
	Password   string      `json:"password" bson:"password"`
	Email      string      `json:"email" bson:"email"`
	State      string      `json:"state" bson:"state"`
	Department *Department `json:"department" bson:"depart_id"`
}

func NewUserFromProto(user *protobuf.User) *User {
	if user == nil {
		return nil
	}
	return &User{
		Id:         user.Id,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		State:      user.State,
		Department: NewDepartmentFromProto(user.Department),
	}
}
func (u *User) ToProto() *protobuf.User {
	if u == nil {
		return nil
	}
	return &protobuf.User{
		Id:         u.Id,
		Name:       u.Name,
		Username:   u.Username,
		Email:      u.Email,
		State:      u.State,
		Department: u.Department.ToProto(),
	}
}
