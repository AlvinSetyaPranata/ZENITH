package auth

func CheckRoleCannotDuplicate(rolename string) bool {
	// if it's true
	// then it cannot have multiple role and vice versa

	switch {
	case rolename == "superadmin":
		return true
	default:
		return false
	}
}
