package rbac

// AuthorizationRequest represents a request to grant or remove access permissions
type AuthorizationRequest struct {
	// The set of access which is granted or removed
	Access  []string    `json:"access,omitempty"`
	Subject *SubjectRef `json:"subject,omitempty"`
	Target  *TargetRef  `json:"target,omitempty"`
}