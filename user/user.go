type User struct {
	firstname string
	created time.Time
}

func newUser(firstname string) (*User, error) {
	if firstname == "" {
		return nil, errors.New("firstname needed")
	} else {
		return &User {
			firstname: firstname,
			created: time.Now(),
		}, nil
	}
}

func (u User) outputDetails() {
	fmt.Println(u.firstname)
	fmt.Println(u.created)
}

func (u *User) clearUserData() {
	u.firstname = ""
}