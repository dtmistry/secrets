package secrets

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	EMPTY_SECRET_FILE                 = "empty-test-secret-file"
	SECRET_FILE                       = "test-secret-file"
	SECRET_FILE_CONTENTS              = "test-value"
	SECRET_PROP_FILE                  = "test-secret-prop-file"
	SECRET_PROP_FILE_CONTENTS         = "db-user=db-pass"
	INVALID_SECRET_PROP_FILE          = "invalid-test-secret-prop-file"
	INVALID_SECRET_PROP_FILE_CONTENTS = "db-userdb-pass"
)

type SecretsTester struct {
	suite.Suite
	Dir string
}

func TestSecretsTester(t *testing.T) {
	s := new(SecretsTester)
	suite.Run(t, s)
}

func (s *SecretsTester) SetupSuite() {
	s.Dir = os.TempDir()
	ioutil.WriteFile(filepath.Join(s.Dir, SECRET_FILE), []byte(SECRET_FILE_CONTENTS), 0755)
	ioutil.WriteFile(filepath.Join(s.Dir, SECRET_PROP_FILE), []byte(SECRET_PROP_FILE_CONTENTS), 0755)
}

func (s *SecretsTester) TeardownSuite() {
	os.Remove(filepath.Join(s.Dir, SECRET_FILE))
	os.Remove(filepath.Join(s.Dir, SECRET_PROP_FILE))
}

func (s *SecretsTester) TestNewSecrets() {
	secrets, err := NewSecrets("test-dir")
	assert.NotNil(s.T(), secrets)
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), "test-dir", secrets.Location)
	}
}

func (s *SecretsTester) TestNewSecretsErr() {
	secrets, err := NewSecrets("")
	assert.Nil(s.T(), secrets)
	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), ErrMissingLocation, err)
	}
}

func (s *SecretsTester) TestNewDefaultSecrets() {
	secrets := NewDefaultSecrets()
	assert.NotNil(s.T(), secrets)
	assert.Equal(s.T(), DEFAULT_SECRETS_LOCATION, secrets.Location)
}

func (s *SecretsTester) TestRead() {
	secrets, err := NewSecrets(s.Dir)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.Read(SECRET_FILE)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secret)
	assert.Equal(s.T(), SECRET_FILE_CONTENTS, secret)
}

func (s *SecretsTester) TestReadBadLocation() {
	secrets, err := NewSecrets("/tmp/location/")
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.Read(SECRET_FILE)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "", secret)
}

func (s *SecretsTester) TestReadEmptyFile() {
	secrets, err := NewSecrets(s.Dir)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.Read(EMPTY_SECRET_FILE)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), "", secret)
}

func (s *SecretsTester) TestReadAsMap() {
	secrets, err := NewSecrets(s.Dir)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.ReadAsMap(SECRET_PROP_FILE)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secret)
	assert.Equal(s.T(), 1, len(secret))
}

func (s *SecretsTester) TestReadAsMapInvalidFile() {
	secrets, err := NewSecrets(s.Dir)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.ReadAsMap(INVALID_SECRET_PROP_FILE)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), 0, len(secret))
}

func (s *SecretsTester) TestReadNoName() {
	secrets := NewDefaultSecrets()
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.Read("")
	assert.Error(s.T(), err)
	assert.Equal(s.T(), ErrMissingSecretName, err)
	assert.Equal(s.T(), "", secret)
}

func (s *SecretsTester) TestReadAsMapNoName() {
	secrets := NewDefaultSecrets()
	assert.NotNil(s.T(), secrets)
	secret, err := secrets.ReadAsMap("")
	assert.Error(s.T(), err)
	assert.Equal(s.T(), ErrMissingSecretName, err)
	assert.Equal(s.T(), 0, len(secret))
}
