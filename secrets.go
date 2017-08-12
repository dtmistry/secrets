package secrets

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	DEFAULT_SECRETS_LOCATION string = "/run/secrets/"
)

var ErrMissingLocation = errors.New("location to read secrets from is required")
var ErrMissingSecretName = errors.New("secretName is required")

type Secrets struct {
	Location string
}

func (s *Secrets) readLines(fileName string) ([]string, error) {
	var lines []string
	f, err := os.Open(filepath.Join(s.Location, fileName))
	if err != nil {
		return lines, errors.Wrapf(err, "Unable to read secret file %s, fileName")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, errors.Wrapf(err, "Unable to read secret file %s, fileName")
	}
	return lines, nil
}

func (s *Secrets) readFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(s.Location, fileName))
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read secret file %s", fileName)
	}
	//Docker allows to create empty secret
	if len(data) == 0 {
		return "", fmt.Errorf("Secret %s appears to be empty", fileName)
	}
	return string(data), nil
}

func (s *Secrets) ReadAsMap(secretName string) (map[string]string, error) {
	secret := make(map[string]string)
	if len(secretName) == 0 {
		return secret, ErrMissingSecretName
	}
	data, err := s.readLines(secretName)
	if err != nil {
		return secret, err
	}
	for i := range data {
		parts := strings.Split(data[i], "=")
		if len(parts) != 2 {
			return secret, fmt.Errorf("Invalid content in secret file %s", secretName)
		} else {
			secret[parts[0]] = parts[1]
		}
	}
	return secret, nil
}

func (s *Secrets) Read(secretName string) (string, error) {
	var secret string
	if len(secretName) == 0 {
		return secret, ErrMissingSecretName
	}
	return s.readFile(secretName)
}

func NewSecrets(location string) (*Secrets, error) {
	if len(location) == 0 {
		return nil, ErrMissingLocation
	}
	return &Secrets{
		Location: location,
	}, nil
}

func NewDefaultSecrets() *Secrets {
	return &Secrets{
		Location: DEFAULT_SECRETS_LOCATION,
	}
}
