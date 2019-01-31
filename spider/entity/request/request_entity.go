package request

type BaseRequestBody struct {
	Offset    string `json:"offset"`
	Total     string `json:"total"`
	Limit     string `json:"limit"`
	CsrfToken string `json:"csrf_token"`
}

// comment request body
type CommentRequestBody struct {
	Rid string `json:"rid"`
	BaseRequestBody
}

type PlayRecordRequestBody struct {
	BaseRequestBody
	Type string `json:"type"`
	Uid  string `json:"uid"`
}
