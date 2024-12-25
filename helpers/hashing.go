package helpers

import "strconv"

// Set the salt for the hashing algorithm (this should be kept secret in a real application
const salt string = "?^&*%$#@!:;.,~`"

/**
 * Hashes a password using a simple hashing algorithm
 *
 * @param string password
 * @return string
 */
func HashPassword(password string) string {
	hashed := 0

	for i := 0; i < len(password); i++ {
		hashed ^= (hashed << 5) + (hashed >> 2) + int(password[i])
	}

	convertedHash := strconv.Itoa(hashed)

	saltedHash := ""

	for i := 0; i < len(convertedHash); i++ {
		saltedHash += string(salt[int(convertedHash[i])-'0'])
	}

	return saltedHash
}

/**
 * Verifies a password against a hashed password
 *
 * @param string password
 * @param string hashedPassword
 * @return bool
 */
func VerifyPassword(password, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}
