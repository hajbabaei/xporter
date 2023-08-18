package domain

type XPost struct {
	Id      string `json:"id" bson:"id"`
	Url     string `json:"url" bson:"url"`
	Profile string `json:"profile" bson:"profile"`
	Content string `json:"Content" bson:"Content"`
}

type XPorterService {
	GetPost(postId string) (*XPost, error)
}