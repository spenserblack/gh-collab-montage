// Provides a user source interface for the user package.
package usersource

// User is a struct that represents a GitHub user.
type User struct {
	// Login is the user's GitHub username.
	Login string `json:"login"`
	// AvatarURL is the URL of the user's avatar.
	AvatarURL string `json:"avatar_url"`
	// Type is either "User" or "Bot".
	Type string `json:"type"`
}

// UserSource provides a utility to iterate over users.
type UserSource interface {
	// Next returns the next user in the source and whether or not to stop iterating.
	Next() (user User, stop bool, err error)
}

// RESTClient is a GitHub REST API client.
//
// This interface is provided for testing purposes.
type RESTClient interface {
	// Get makes a GET request to the GitHub API.
	Get(path string, resp interface{}) error
}
