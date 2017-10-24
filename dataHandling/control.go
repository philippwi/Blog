package dataHandling

func UserExists(name string) bool {
	existingUsers := GetUsers()

	for _, user := range existingUsers {
		if user.Name == name {
			return true
		}
	}
	return false
}

func PasswordCorrect(name, password string) bool {
	existingUsers := GetUsers()

	for _, user := range existingUsers {
		if user.Name == name && user.Password == password {
			return true
		}
	}
	return false
}
