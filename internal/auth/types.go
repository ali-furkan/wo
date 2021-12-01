package auth

type Repository struct {
	Name      string
	URL       string
	CreatedAt string
}

type User struct {
	Name  string
	URL   string
	Email string
	Repos []*Repository
}
