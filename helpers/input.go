package helpers

import (
  "os"
  "bufio"
  "strings"
)

/**
  * Get input from the user
  *
  * @return string
  */
func GetInlineInput() string {
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  input = strings.TrimSuffix(input, "\n")
  input = strings.TrimSuffix(input, "\r")
  return input
}
