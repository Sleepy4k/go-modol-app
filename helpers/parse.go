package helpers

import (
  "strings"
  "encoding/json"
)

/**
  * Loads a JSON file into a struct
  *
  * @param []byte content
  * @param interface{} v
  * @return error
  */
func LoadFromJSON(content []byte, v interface{}) error {
  return json.Unmarshal(content, v)
}

/**
  * Saves a struct to a JSON file
  *
  * @param interface{} v
  * @return []byte
  * @return error
  */
func SaveToJSON(v interface{}) ([]byte, error) {
  return json.Marshal(v)
}

/**
  * Make word to uppercase level
  *
  * @param string word
  * @return string
  */
func ToUpper(word string) string {
  return strings.ToUpper(word)
}

/**
  * Make word to lowercase level
  *
  * @param string word
  * @return string
  */
func ToLower(word string) string {
  return strings.ToLower(word)
}
