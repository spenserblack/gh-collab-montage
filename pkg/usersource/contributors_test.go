package usersource

import (
	"errors"
	"testing"
)

func TestNewContributors(t *testing.T) {
	var api RESTClient
	contributors := NewContributors(api, "owner", "repo")
	if contributors.endpoint != "repos/owner/repo/contributors" {
		t.Errorf(`contributors.endpoint = %q, want "repos/owner/repo/contributors"`, contributors.endpoint)
	}
	if contributors.page != 1 {
		t.Errorf("contributors.page = %d, want 1", contributors.page)
	}
}

func TestContributors_Next(t *testing.T) {
	mockUser := User{
		Login:     "user",
		AvatarURL: "https://example.com/avatar.png",
		Type:      "User",
	}
	mockErr := errors.New("error")
	tests := []struct {
		name string
		api  RESTClient
		user User
		stop bool
		err  error
	}{
		{
			name: "no users",
			api: &MockContributorRESTClient{
				Users: [][]User{{}, {}},
				Err:   nil,
			},
			user: User{},
			stop: true,
			err:  nil,
		},
		{
			name: "error returned by api.Get",
			api: &MockContributorRESTClient{
				Users: [][]User{{}, {}},
				Err:   mockErr,
			},
			user: User{},
			stop: true,
			err:  mockErr,
		},
		{
			name: "one user",
			api: &MockContributorRESTClient{
				Users: [][]User{
					{mockUser},
				},
				Err: nil,
			},
			user: mockUser,
			stop: false,
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contributors := NewContributors(tt.api, "owner", "repo")
			user, stop, err := contributors.Next()
			if user != tt.user {
				t.Errorf("user = %v, want %v", user, tt.user)
			}
			if stop != tt.stop {
				t.Errorf("stop = %v, want %v", stop, tt.stop)
			}
			if err != tt.err {
				t.Errorf("err = %v, want %v", err, tt.err)
			}
		})
	}
}

// Tests that the page is incremented when the users slice is empty.
func TestContributors_Next_pageIncremented(t *testing.T) {
	mockUser := User{
		Login:     "user",
		AvatarURL: "https://example.com/avatar.png",
		Type:      "User",
	}
	contributors := NewContributors(&MockContributorRESTClient{
		Users: [][]User{
			{mockUser},
		},
		Err: nil,
	}, "owner", "repo")
	_, _, _ = contributors.Next()
	if contributors.page != 2 {
		t.Errorf("contributors.page = %d, want 2", contributors.page)
	}
}

// MockContributorRESTClient is used for testing the Contributors type.
type MockContributorRESTClient struct {
	// Users is the list of pages of users to return.
	Users [][]User
	// Err is the error to return.
	Err error
}

// Get returns the next page of users.
func (c *MockContributorRESTClient) Get(path string, resp interface{}) error {
	var users []User
	users, c.Users = c.Users[0], c.Users[1:]
	*resp.(*[]User) = users
	return c.Err
}
