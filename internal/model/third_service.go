package model

type IntrospectRes struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    *TokenData `json:"data"`
}

type TokenData struct {
	UserID   string             `json:"user_id"`
	UserName string             `json:"user_name"`
	OrgID    string             `json:"org_id"`
	Roles    []map[int64]string `json:"roles"`
}
