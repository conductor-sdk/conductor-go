package integration

// Category represents the type of integration
type Category string

const (
	API            Category = "API"
	AI_MODEL       Category = "AI_MODEL"
	VECTOR_DB      Category = "VECTOR_DB"
	RELATIONAL_DB  Category = "RELATIONAL_DB"
	MESSAGE_BROKER Category = "MESSAGE_BROKER"
	GIT            Category = "GIT"
	EMAIL          Category = "EMAIL"
)