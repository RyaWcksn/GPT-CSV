package storechild

const (
	QueryCreateUserChild   = `INSERT INTO nanne.users_child (user_parent_id, child_name, age, role_name) VALUES (?, ?, ?, ?);`
	QueryGetOneUserChild   = `SELECT child_name, role_name, age FROM nanne.users_child WHERE user_parent_id=? AND child_name=?;`
	QueryGetListChild      = `SELECT child_name, role_name, age FROM nanne.users_child WHERE user_parent_id=? ORDER BY updated DESC LIMIT ?,?;`
	QueryUpdateSingleChild = `UPDATE nanne.users_child SET role_name=?, age=? WHERE user_parent_id=? AND child_name=?;`
)
