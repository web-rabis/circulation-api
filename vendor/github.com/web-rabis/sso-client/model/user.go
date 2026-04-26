package model

import "github.com/web-rabis/sso-client/protobuf"

type User struct {
	Id         int64       `json:"id" bson:"id"`
	Name       string      `json:"name" bson:"name"`
	Username   string      `json:"username" bson:"username"`
	Password   string      `json:"password" bson:"password"`
	Email      string      `json:"email" bson:"email"`
	State      string      `json:"state" bson:"state"`
	Department *Department `json:"department" bson:"depart_id"`
}
type Department struct {
	Id   int64  `json:"id" bson:"id"`
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
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
func NewDepartmentFromProto(department *protobuf.Department) *Department {
	if department == nil {
		return nil
	}
	return &Department{
		Id:   department.Id,
		Code: department.Code,
		Name: department.Name,
		Type: department.Type,
	}
}
