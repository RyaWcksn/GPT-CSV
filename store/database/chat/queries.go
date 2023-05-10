package storechat

const (
	QueryCreateNewMessage = `INSERT INTO nanne.chat_history (user_parent_id, child_name, role_name, question, answer) VALUES (?, ?, ?, ?, ?);`
)
