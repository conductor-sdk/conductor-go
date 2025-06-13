package security

// Access represents the access level for permissions
type Access string

const (
	CREATE  Access = "CREATE"
	READ    Access = "READ"
	EXECUTE Access = "EXECUTE"
	UPDATE  Access = "UPDATE"
	DELETE  Access = "DELETE"
)