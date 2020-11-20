package github

// {
// 	"name": "Hello-World",
// 	"description": "This is your first repository",
// 	"homepage": "https://github.com",
// 	"private": false,
// 	"has_issues": true,
// 	"has_projects": true,
// 	"has_wiki": true
//  }

type CreateRepo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomePage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projcets"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"push"`
	HasPush bool `json:"pull"`
}
