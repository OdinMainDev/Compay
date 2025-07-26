package merchantuser

func (that *UserModel) GetUserUsername() string {
	if that.Username != "" {
		return that.Username
	}
	return that.FirstName + " " + that.LastName
}
