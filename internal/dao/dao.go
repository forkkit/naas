package dao

var (
	// OAuth2Client ...
	OAuth2Client OAuth2Clienter = &oauth2Client{}
	// OAuth2ClientInfo ...
	OAuth2ClientInfo OAuth2ClientInfoer = &oauth2ClientInfo{}
	// Admin ...
	Admin Adminer = &admin{}
	// User ...
	User Userer = &user{}
	// Organization ...
	Organization Organizationer = &organization{}
	// OrganizationRole ...
	OrganizationRole OrganizationRoleer = &organizationRole{}
	// Role ...
	Role Roleer = &role{}
	// UserRole ...
	UserRole UserRoleer = &userRole{}
)
