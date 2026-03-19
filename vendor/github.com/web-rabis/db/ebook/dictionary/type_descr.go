package dictionary

type TypeDescription struct {
	Dictionary `bson:"_"`
	NameKz     string `json:"name_kz" bson:"name_kz"`
}
