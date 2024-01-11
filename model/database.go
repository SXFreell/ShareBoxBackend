package model

type User struct {
	ID         int    `json:"id"`
	UID        string `json:"uid"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type Text struct {
	ID          int    `json:"id"`
	UID         string `json:"uid"`
	Code        string `json:"code"`
	Content     string `json:"content"`
	Expires     int64  `json:"expires"`
	PickupCount int    `json:"pickup_count"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type File struct {
	ID          int    `json:"id"`
	UID         string `json:"uid"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Expires     int64  `json:"expires"`
	PickupCount int    `json:"pickup_count"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
