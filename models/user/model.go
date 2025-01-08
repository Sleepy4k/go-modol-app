package user

import (
  "modol-app/enums"
  "modol-app/helpers"
)

// User model
type User struct {
  ID          int
  Nama        string
  NIM         string
  Role        enums.Role
  Password    string
}

// Users collection
var Users []User

/**
  * Initializes the Users collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("users.json") {
    content, err := helpers.ReadFile("users.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Users)
  } else {
    Users = []User{
      {1, "Admin", "admin", enums.Admin, helpers.HashPassword("password")},
      {2, "Arif Samsudin", "A12345678", enums.Guru, helpers.HashPassword("password")},
      {3, "Regita", "2411102001", enums.Siswa, helpers.HashPassword("password")},
    }

    content, err := helpers.SaveToJSON(Users)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("users.json", content)

    if err != nil {
      panic(err)
    }
  }
}

/**
  * Inserts a user
  *
  * @param User user
  * @return void
  */
func InsertUser(user User) {
  user.ID = len(Users) + 1
  Users = append(Users, user)

  content, err := helpers.SaveToJSON(Users)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("users.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Finds a user by their NIM
  *
  * @param string NIM
  * @return *User
  */
func FindUserByNIM(NIM string) *User {
  for _, user := range Users {
    if user.NIM == NIM {
      return &user
    }
  }

  return &User{}
}
