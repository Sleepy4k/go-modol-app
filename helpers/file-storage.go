package helpers

import (
  "fmt"
  "os"
)

// Set the path to the file storage directory
const fileStoragePath string = "./storage"

/**
  * Checks if a file exists in the file storage directory
  *
  * @param string filename
  * @return void
  */
func init() {
  // Create the file storage directory if it doesn't exist
  if _, err := os.Stat(fileStoragePath); os.IsNotExist(err) {
    os.Mkdir(fileStoragePath, 0755)
  }
}

/**
  * Checks if a file exists in the file storage directory
  *
  * @param string filename
  * @return bool
  */
func FileExists(filename string) bool {
  _, err := os.Stat(fmt.Sprintf("%s/%s", fileStoragePath, filename))
  return !os.IsNotExist(err)
}

/**
  * Saves a file to the file storage directory
  *
  * @param string filename
  * @param []byte data
  * @return error
  */
func SaveFile(filename string, data []byte) error {
  if FileExists(filename) {
    return fmt.Errorf("file already exists")
  }

  return os.WriteFile(fmt.Sprintf("%s/%s", fileStoragePath, filename), data, 0666)
}

/**
  * Updates a file in the file storage directory
  *
  * @param string filename
  * @param []byte data
  * @return error
  */
func UpdateFile(filename string, data []byte) error {
  if !FileExists(filename) {
    err := os.WriteFile(fmt.Sprintf("%s/%s", fileStoragePath, filename), data, 0666)
    return err
  }

  // Backup the existing file
  err := os.Rename(fmt.Sprintf("%s/%s", fileStoragePath, filename), fmt.Sprintf("%s/%s.bak", fileStoragePath, filename))

  if err != nil {
    fmt.Printf("Error backing up file: %s\n", err)
    return err
  }

  err = os.WriteFile(fmt.Sprintf("%s/%s", fileStoragePath, filename), data, 0666)

  if err != nil {
    fmt.Printf("Error updating file: %s\n", err)
    os.Rename(fmt.Sprintf("%s/%s.bak", fileStoragePath, filename), fmt.Sprintf("%s/%s", fileStoragePath, filename))
  } else {
    os.Remove(fmt.Sprintf("%s/%s.bak", fileStoragePath, filename))
  }

  return err
}

/**
  * Reads a file from the file storage directory
  *
  * @param string filename
  * @return []byte
  * @return error
  */
func ReadFile(filename string) ([]byte, error) {
  return os.ReadFile(fmt.Sprintf("%s/%s", fileStoragePath, filename))
}

/**
  * Deletes a file from the file storage directory
  *
  * @param string filename
  * @return error
  */
func DeleteFile(filename string) error {
  return os.Remove(fmt.Sprintf("%s/%s", fileStoragePath, filename))
}
