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
  * Creates a new ID using insertion sort
  *
  * @return int
  */
func CreateNewID() int {
  LENGTH := len(Users)

  for i := 1; i < LENGTH; i++ {
    key := Users[i]
    j := i - 1

    for j >= 0 && Users[j].ID > key.ID {
      Users[j + 1] = Users[j]
      j = j - 1
    }

    Users[j + 1] = key
  }

  return Users[LENGTH - 1].ID + 1
}
  
/**
  * Inserts a user
  *
  * @param User user
  * @return void
  */
func InsertUser(user User) {
  user.ID = CreateNewID()
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
  * Finds a user by their ID using binary search
  *
  * @param int ID
  * @return *User
  */
func FindUserByID(ID int) *User {
  if ID < 1 {
    return &User{}
  }

  left := 0
  right := len(Users) - 1

  for left <= right {
    mid := left + (right - left) / 2

    if Users[mid].ID == ID {
      return &Users[mid]
    }

    if Users[mid].ID < ID {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return &User{}
}

/**
  * Finds a user by their NIM using sequential search
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
