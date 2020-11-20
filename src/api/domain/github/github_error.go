package github

type GitHubErrorResponse struct {
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"doumentation_url"`
	Errors           []GitHubError `json:"error"`
}

type GitHubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
