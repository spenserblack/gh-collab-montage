package usersource

import "fmt"

// Contributors provides a repository's contributors from the GitHub API.
type Contributors struct {
	api RESTClient
	// endpoint is the GitHub API endpoint.
	endpoint string
	// perPage is the number of users to request per page.
	perPage int
	// page is the current page of users.
	page int
	// users is an internal cache of users to iterate over.
	users []User
}

// NewContributors returns a new Contributors.
func NewContributors(api RESTClient, owner, repo string) *Contributors {
	const perPage int = 30
	endpoint := fmt.Sprintf("repos/%s/%s/contributors", owner, repo)
	return &Contributors{
		api:      api,
		endpoint: endpoint,
		perPage:  perPage,
		page:     1,
		users:    make([]User, 0, perPage),
	}
}

// Next returns the next user in the source and whether or not to stop iterating.
func (c *Contributors) Next() (user User, stop bool, err error) {
	if len(c.users) == 0 {
		err = c.fetchUsers()
		if err != nil {
			stop = true
			return
		}
	}
	if len(c.users) == 0 {
		stop = true
		return
	}
	user, c.users = c.users[0], c.users[1:]
	stop = false
	return
}

// FetchUsers fetches the next page of users from the GitHub API.
func (c *Contributors) fetchUsers() error {
	response := []User{}
	err := c.api.Get(c.url(), &response)
	if err != nil {
		return err
	}
	c.users = response
	c.page++
	return nil
}

// URL returns the URL to request.
func (c Contributors) url() string {
	return fmt.Sprintf("%s?page=%d&per_page=%d", c.endpoint, c.page, c.perPage)
}
