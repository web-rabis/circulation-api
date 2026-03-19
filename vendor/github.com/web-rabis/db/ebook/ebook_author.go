package ebook

type EbookAuthor struct {
	Id             int64  `json:"id,omitempty" bson:"id"`
	EbookId        int64  `json:"ebook_id,omitempty" bson:"ebook_id"`
	Author         string `json:"author,omitempty" bson:"author"`
	DinNumber      string `json:"din_number,omitempty" bson:"din_number"`
	TitleAuthor    string `json:"title_author,omitempty" bson:"title_author"`
	LifeDates      string `json:"life_dates,omitempty" bson:"life_dates"`
	RoleAuthor     string `json:"role_author,omitempty" bson:"role_author"`
	NerazdelAuthor int    `json:"nerazdel_author,omitempty" bson:"nerazdel_author"`
	KeyValue       int    `json:"key_value,omitempty" bson:"key_value"`
}
