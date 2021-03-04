package repository

const (
	createEmailQuery = `INSERT INTO emails (addressfrom, addressto, subject, message) 
	VALUES ($1, $2, $3, $4) 
	RETURNING email_id, addressfrom, addressto, subject, message, created_at`

	getByIDQuery = `SELECT email_id, addressfrom, addressto, subject, message, created_at FROM emails WHERE email_id = $1`

	searchTotalCountQuery = `SELECT count(email_id) FROM emails WHERE addressto ILIKE '%' || $1 || '%' 
	ORDER BY created_at OFFSET $2 LIMIT  $3`

	searchQuery = `SELECT email_id, addressfrom, addressto, subject, message, created_at 
	FROM emails WHERE addressto ILIKE '%' || $1 || '%' ORDER BY created_at OFFSET $2 LIMIT  $3`
)
