package authentication

import (
	"testing"

	wdbErrors "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/stretchr/testify/assert"
)

var (
	hashableString string = "hashThis^string$01"

	// hash values of hashableString using various hashing algorithms
	// generated using: https://www.browserling.com/tools/all-hashes
	hashedStringMD5      string = "59e7bd0713e94abf223e8a34f298794b"
	hashedStringSHA1     string = "26f32439b6d2dcdddcd11d2fd4d73ca6b24b4a77"
	hashableStringSHA256 string = "65d9de2c2b995bf8a123e2af1ddf566b8e16b3b9873c5c13ad2c60b7fc3ab249"

	// sample auth headers
	sampleAuthHeader1                                    string = "Basic dXNlcm5hbWU6cGFzc3dvcmQ="
	sampleAuthHeader1Username, sampleAuthHeader1Password string = "username", "password"

	// fix this example
	// sampleAuthHeader2                                    string = "Basic YWJjOjEyMw==="
	// sampleAuthHeader2Username, sampleAuthHeader2Password string = "abc", "123"
)

func Test_Hash(t *testing.T) {
	testCases := []struct {
		hashingAlgorithm     string
		expectedHashedString string
	}{
		{
			hashingAlgorithm:     MD5,
			expectedHashedString: hashedStringMD5,
		},
		{
			hashingAlgorithm:     SHA1,
			expectedHashedString: hashedStringSHA1,
		},
		{
			hashingAlgorithm:     SHA256,
			expectedHashedString: hashableStringSHA256,
		},
		{
			hashingAlgorithm:     "SHA512",             // not supported, should use sha256 algorithm instead
			expectedHashedString: hashableStringSHA256, // return sha256 hash as sha512 not supported
		},
	}

	for _, tc := range testCases {
		got := Hash(hashableString, tc.hashingAlgorithm)
		assert.Equal(t, tc.expectedHashedString, got)
	}
}

func Test_HandleUserCredentials(t *testing.T) {
	testCases := []struct {
		authorizationHeaderString string
		expectedUsername          *string
		expectedPassword          *string
		expectedError             *wdbErrors.WdbError
	}{
		{
			authorizationHeaderString: sampleAuthHeader1,
			expectedUsername:          &sampleAuthHeader1Username,
			expectedPassword:          &sampleAuthHeader1Password,
			expectedError:             nil,
		},
		// TODO: fix this case, debug to get issue
		// {
		// 	authorizationHeaderString: sampleAuthHeader2,
		// 	expectedUsername:          &sampleAuthHeader2Username,
		// 	expectedPassword:          &sampleAuthHeader2Password,
		// 	expectedError:             nil,
		// },
		{
			authorizationHeaderString: "",
			expectedUsername:          nil,
			expectedPassword:          nil,
			expectedError:             &wdbErrors.InvalidCredentialsError,
		},
	}

	for _, tc := range testCases {
		gotUsername, gotPassword, err := HandleUserCredentials(tc.authorizationHeaderString)
		assert.Equal(t, tc.expectedUsername, gotUsername)
		assert.Equal(t, tc.expectedPassword, gotPassword)
		assert.Equal(t, tc.expectedError, err)
	}

}

func Test_GetActor(t *testing.T) {
	testCases := []struct {
		authorizationHeaderString string
		expectedActor             string
	}{
		{
			authorizationHeaderString: sampleAuthHeader1,
			expectedActor:             sampleAuthHeader1Username,
		},
		{
			authorizationHeaderString: "",
			expectedActor:             "",
		},
	}

	for _, tc := range testCases {
		gotActor := GetActor(tc.authorizationHeaderString)
		assert.Equal(t, tc.expectedActor, gotActor)
	}
}
