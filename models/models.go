package models

type MessageIn struct {
	Pk int `json:"pk"`
	Title string `json:"title"`
	Body string `json:"body"`
	Action string `json:"action"`
	ExtraUID string `json:"extra_uid"`
	FcmTokens []string `json:"fcm_tokens"`
}
