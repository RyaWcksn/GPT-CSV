package storechild

const (
	QueryCreateUserChild = `INSERT INTO nanne.users_child (user_parent_id, child_name, age, role_name) VALUES (?, ?, ?, ?);`
)
