package model

// Set
// Set Request
type SetReq struct {
	Type           string     `json:"type"`
	SetTextContent SetTextReq `json:"set_text_content"`
	SetFileContent SetFileReq `json:"set_file_content"`
}

type SetTextReq struct {
	Content string `json:"content"`
	Expires int    `json:"expires"`
}

type SetFileReq struct {
	Content string `json:"content"`
	Expires int    `json:"expires"`
}

// Set Response

// Get
// Get Request
type GetReq struct {
	Code string `json:"code"`
}

// Get Response

// CREATE TABLE IF NOT EXISTS text (
// id INTEGER PRIMARY KEY AUTOINCREMENT,
// uid TEXT NOT NULL,
// code TEXT NOT NULL,
// content TEXT NOT NULL,
// expires TEXT NOT NULL,
// create_time TEXT NOT NULL,
// update_time TEXT NOT NULL
