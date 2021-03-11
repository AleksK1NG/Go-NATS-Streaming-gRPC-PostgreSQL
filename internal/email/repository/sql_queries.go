package repository

const (
	createEmailQuery = `INSERT INTO emails (address_from, address_to, subject, message) 
	VALUES ($1, $2, $3, $4) 
	RETURNING email_id, address_from, address_to, subject, message, created_at`

	getByIDQuery = `SELECT email_id, address_from, address_to, subject, message, created_at FROM emails WHERE email_id = $1`

	searchTotalCountQuery = `SELECT count(email_id)
	FROM emails
	WHERE document_with_idx @@ to_tsquery($1)`

	searchQuery = `SELECT email_id, address_to, address_from, subject, message, created_at
	FROM emails
	WHERE document_with_idx @@ to_tsquery($1) ORDER BY created_at OFFSET $2 LIMIT $3`
)
