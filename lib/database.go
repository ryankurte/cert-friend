package cafriend

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Database contains a list of certificates for a given CA
type Database struct {
	root         []Certificate
	intermediate []Certificate
	server       []Certificate
	client       []Certificate
}

// LoadDatabase loads a certificate database file
func LoadDatabase(name string) (*Database, error) {

	// Read file
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// Parse yaml
	var db Database
	err = yaml.Unmarshal(bytes, &db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}

// Save the database to a file
func (db *Database) Save(name string) error {
	// Marshal to yaml
	bytes, err := yaml.Marshal(db)
	if err != nil {
		return err
	}

	// Write file
	err = ioutil.WriteFile(name, bytes, 0600)
	if err != nil {
		return err
	}

	return nil
}
