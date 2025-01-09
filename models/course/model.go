package course

import (
  "modol-app/enums"
  "modol-app/helpers"
)

// Course model
type Course struct {
  ID          int
  Nama        string
}

// Courses collection
var Courses []Course

/**
  * Initializes the courses collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("courses.json") {
    content, err := helpers.ReadFile("courses.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Courses)
  } else {
    Courses = []Course{
      {1, "Pemrograman Berorientasi Objek"},
      {2, "Struktur Data"},
      {3, "Algoritma dan Pemrograman"},
      {4, "Basis Data"},
      {5, "Pemrograman Web"},
      {6, "Pemrograman Mobile"},
      {7, "Pemrograman Jaringan"},
      {8, "Pemrograman Fungsional"},
    }

    content, err := helpers.SaveToJSON(Courses)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("courses.json", content)

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
  LENGTH := len(Courses)

  for i := 1; i < LENGTH; i++ {
    key := Courses[i]
    j := i - 1

    for j >= 0 && Courses[j].ID > key.ID {
      Courses[j + 1] = Courses[j]
      j = j - 1
    }

    Courses[j + 1] = key
  }

  return Courses[LENGTH - 1].ID + 1
}

/**
  * Lists all courses with selection sort
  *
  * @param enums.SORT sortType
  * @return []Course
  */
func ListCourses(sortType enums.SORT) []Course {
  LENGTH := len(Courses)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Courses[j].ID < Courses[min].ID {
          min = j
        }
      } else {
        if Courses[j].ID > Courses[min].ID {
          min = j
        }
      }
    }

    Courses[i], Courses[min] = Courses[min], Courses[i]
  }

  return Courses
}

/**
  * Finds a course by their ID using binary search
  *
  * @param string ID
  * @return *Courses
  */
func FindCourseByID(ID int) *Course {
  if ID < 1 {
    return &Course{}
  }

  left := 0
  right := len(Courses) - 1

  for left <= right {
    mid := left + (right - left) / 2

    if Courses[mid].ID == ID {
      return &Courses[mid]
    }

    if Courses[mid].ID < ID {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return &Course{}
}
