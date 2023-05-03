package storeroles

const (
	QueryCreateRoles          = `INSERT INTO nanne.roles (user_parent_id, topic, rules, goals, child_description, role_name, role_description) VALUES (?, ?, ?, ?, ?, ?, ?);`
	QueryGetOneRole           = `SELECT topic, rules, goals, child_description, role_name, role_description FROM nanne.roles WHERE user_parent_id=? AND role_name=?;`
	QueryGetListRole          = `SELECT role_name, topic, rules, goals, child_description, role_description FROM nanne.roles WHERE user_parent_id=? ORDER BY updated DESC LIMIT ?,?;`
	QueryUpdateSingleRoleById = `UPDATE nanne.roles SET topic=?, rules=?, goals=?, child_description=?, role_description=? WHERE user_parent_id=? AND role_name=?;`
)
