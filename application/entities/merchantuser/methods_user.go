package merchantuser

func (that *Model) GetUserUsername() string {
	if that.Username != "" {
		return that.Username
	}
	return that.FirstName + " " + that.LastName
}
