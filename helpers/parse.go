package helpers

import "encoding/json"

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
