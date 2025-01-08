package auth

import (
  "fmt"
  "modol-app/enums"
  "modol-app/helpers"
  "modol-app/models/user"
)

var CurrentUser *user.User

/**
  * Authentic
  *
  * @param string NIM
  * @param string password
  * @return *user.User
  */
func Authentic(NIM string, password string) *user.User {
  userData := user.FindUserByNIM(NIM)

  if userData == nil || userData.ID == 0 {
    fmt.Println("User tidak ditemukan")
    return &user.User{}
  }

  if !helpers.VerifyPassword(password, userData.Password) {
    fmt.Println("Password salah")
    return &user.User{}
  }

  CurrentUser = userData

  return userData
}

/**
  * Registers a user
  *
  * @param string NIM
  * @param string nama
  * @param enums.Role role
  * @param string password
  * @return *user.User
  */
func CreateUser(NIM string, nama string, role enums.Role, password string) *user.User {
  userData := user.User{
    NIM: NIM,
    Nama: nama,
    Role: role,
    Password: helpers.HashPassword(password),
  }

  user.InsertUser(userData)

  return &userData
}

/**
  * Checks if a user is authenticated
  *
  * @param *user.User user
  * @return bool
  */
func IsAuthenticated(user *user.User) bool {
  if user == nil {
    return false
  }

  return user.ID != 0
}

/**
  * Checks if a user is an admin
  *
  * @param *user.User user
  * @return bool
  */
func IsAdmin(user *user.User) bool {
  if !IsAuthenticated(user) {
    return false
  }

  return user.Role == enums.Admin
}

/**
  * Checks if a user is a guru
  *
  * @param *user.User user
  * @return bool
  */
func IsGuru(user *user.User) bool {
  if !IsAuthenticated(user) {
    return false
  }

  return user.Role == enums.Guru
}

/**
  * Checks if a user is a siswa
  *
  * @param *user.User user
  * @return bool
  */
func IsSiswa(user *user.User) bool {
  if !IsAuthenticated(user) {
    return false
  }

  return user.Role == enums.Siswa
}
