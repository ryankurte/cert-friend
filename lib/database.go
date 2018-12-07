package cafriend

// Database contains a list of certificates for a given CA
type Database struct {
	root         []Certificate
	intermediate []Certificate
	server       []Certificate
	client       []Certificate
}
