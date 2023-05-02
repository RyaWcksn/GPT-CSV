package storeroles

const (
	QueryCreateRoles    = `INSERT INTO nanne.roles (role_id, topic, rules, goals, child_description, role_name, role_description) VALUES (?, ?, ?, ?, ?, ?, ?);`
	QueryGetOneRoleById = `SELECT topic, rules, goals, child_description, role_name, role_description FROM nanne.roles WHERE role_id=?;`
	QueryGetListRole    = `SELECT topic, rules, goals, child_description, role_name, role_description FROM nanne.roles ORDER BY updated DESC LIMIT ?,?;`
)
