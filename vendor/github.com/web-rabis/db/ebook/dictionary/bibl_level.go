package dictionary

type BibliographicLevel struct {
	Dictionary `bson:"_"`
	TypeEbooks string `json:"type_ebooks" bson:"type_ebooks"`
}
