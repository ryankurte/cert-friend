/**
 * cert-friend
 * Helper utilities
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */

package certfriend

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

// LoadFile loads a yaml file into an object
// This returns true if the file existed and was loaded,
//false if the file didn't exist, and error if an error occured
func LoadFile(filename string, o interface{}) (bool, error) {

	// Check if the file exists
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	// Read file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	// Parse yaml
	err = yaml.Unmarshal(bytes, &o)
	if err != nil {
		return false, err
	}

	return true, nil
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

// SaveFile saves an object into a yaml file
// This returns nil on success and an error if an error occured
func SaveFile(filename string, o interface{}) error {
	// Marshal to yaml
	bytes, err := yaml.Marshal(o)
	if err != nil {
		return err
	}

	// Write file
	err = ioutil.WriteFile(filename, bytes, 0600)
	if err != nil {
		return err
	}

	return nil
}
