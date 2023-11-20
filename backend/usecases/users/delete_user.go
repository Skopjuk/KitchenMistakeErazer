package users

type DeleteUserProfile struct {
	repository DeleteUser
}

func NewDeleteUserProfile(repository DeleteUser) *DeleteUserProfile {
	return &DeleteUserProfile{repository: repository}
}

func (d *DeleteUserProfile) Execute(id int) error {
	return d.repository.DeleteUser(id)
}
