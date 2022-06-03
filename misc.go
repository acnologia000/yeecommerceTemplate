package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"time"
	//"golang.org/x/crypto/bcrypt"
)

// their recommended production values are mentioned after comment

// how quickly application should scan through existing sessions and remove the expired ones
const SESSION_CYCLE_TIME = time.Second // time.Minute
// how long each session should last
const SESSION_DURATION = time.Second * 24 // time.Hour * 24
/*
this application caches semi static content like front page items
to avoid database calls and unltimately improve performance under
high load, however that brings the issue of stale cache and to
avoid that we run a little service on a saperate go routine that
periodically refreshes the cache from database and to tweak the
frequency of that process you can tweak value of below variable

>Default : once every 24h 					 | time.Hour * 24
>Recommended highest : once every 20 minutes | time.Minute * 20
*/
const (
	PAGE_ITEM_COUNT = 20
	PAGE_LIFESPAN   = time.Hour * 24
	LISTEN_ADDRESS  = "0.0.0.0:6600"
)
const SUCCESS_RESPONSE = `{"success":"true"}`

// sql statements
const (
	SEARCH_ITEM_QUERY = "SELECT * From items where name = $1"
	INSERT_NEW_USER   = "INSERT INTO USERS (fullname, email, password_hash) values ($1,$2,$3)"
	LOGIN_SELECT      = "select password_hash from users where email = $1"
)

// report non verbose errors in production
const (
	INVALID_PAGE_NUMBER    = "page number too big or invalid number" // INVAID_PAGE_NUMBER
	INVALID_JSON           = "invalid input format"                  // "BAD_REQUEST"
	ILLEGAL_REQUEST_METHOD = "request method not allowed"            // "BAD_REQUEST"
	INVALID_METHOD         = "invalid method"                        // "BAD_REQUEST"
	JSON_DECODE_ERROR      = "JSON_DECODE_ERROR"                     // "BAD_REQUEST"
	STMT_PREPARE_ERROR     = "error in preparing statement"          // "DB_ERR"
	QUERY_ERROR            = "error in executing querrys"            // "DB_ERR"
	ROW_SCAN_ERROR         = "error in Scanning rows"                // "DB_ERR"
	PASSWORD_ERROR         = "username or password incorrect"        // "WRONG_CREDENTIALS"
	INTERNAL_SERVER_ERROR  = "ERR"
)

func ToBaseEncodedMD5(input string) string {
	sum := md5.Sum([]byte(input))
	return base64.StdEncoding.EncodeToString(sum[:]) // taking a slice out of array
}

func errorResponse(e string) []byte {
	return []byte(fmt.Sprintf(`{"failure":"%s"}`, e))
}

func printDebug(format string, a ...interface{}) {
	if DEBUG {
		fmt.Printf(format, a...)
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// GenerateRandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func setupTestEnv() {
	os.Setenv("REDIS_ADDRESS", "localhost:6379")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "yeecom")
}

// using md5 hash withing sql so i dont need these functions yet
// but the thing to remember is that later MD5 logic has to be
// replaced with bcrypt/scrypt/argon2i for more secure handing
// of userdata as users often reuse passwords

/*
func HashAndPasswordMatch(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func generateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	return string(hash), err
}
*/
