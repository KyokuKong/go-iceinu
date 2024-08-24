package core

func CheckUserRole(qid int64, role int16) bool {
	// 调用User管理器里的查询封装
	user, err := GetUserByQID(qid)
	if err != nil {
		return false
	}
	if user.Role >= role {
		return true
	}
	return false
}

func GetUserRole(qid int64) (int16, error) {
	// 调用User管理器里的查询封装
	user, err := GetUserByQID(qid)
	if err != nil {
		return 0, err
	}
	return user.Role, nil
}

func SetUserRole(qid int64, role int16) error {
	// 调用User管理器里的查询封装
	user, err := GetUserByQID(qid)
	if err != nil {
		return err
	}
	user.Role = role
	return UpdateUser(user)
}
