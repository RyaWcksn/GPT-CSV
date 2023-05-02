package storeroles

const (
	QueryCreateRoles = `INSERT INTO nanne.roles (role_id, topic, child_description, role_name, role_description) VALUES (?, ?, ?, ?, ?);`
)
