package storeusersparent

const (
	QuerySaveUsersParent          = `INSERT INTO nanne.users_parent (user_parent_id, password, name, email, phone_number, status) VALUES (?, ?, ?, ?, ?, ?);`
	QuerySelectOneUsersParentById = `SELECT password, name, email, phone_number, status FROM nanne.users_parent WHERE user_parent_id=?;`
)
