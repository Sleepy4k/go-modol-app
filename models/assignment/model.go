package assignment

import (
  "modol-app/enums"
  "modol-app/helpers"
  "modol-app/models/class"
  "modol-app/models/course"
)

/**
  * Assignment struct
  */
type Assignment struct {
  ID                  int
  CourseID            int
  ClassID             int
  Title               string
  Description         string
}

// Assignments collection
var Assignments []Assignment

/**
  * Initializes the Assignments collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("assignments.json") {
    content, err := helpers.ReadFile("assignments.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Assignments)
  } else {
    Assignments = []Assignment{
      {1, 1, 8, "Tugas 1", "Tugas 1 PBO"},
      {2, 2, 7, "Tugas 1", "Tugas 1 Struktur Data"},
      {3, 3, 6, "Tugas 1", "Tugas 1 Algoritma dan Pemrograman"},
      {4, 4, 5, "Tugas 1", "Tugas 1 Basis Data"},
      {5, 5, 4, "Tugas 1", "Tugas 1 Pemrograman Web"},
      {6, 6, 3, "Tugas 1", "Tugas 1 Pemrograman Mobile"},
      {7, 7, 2, "Tugas 1", "Tugas 1 Pemrograman Jaringan"},
      {8, 8, 1, "Tugas 1", "Tugas 1 Pemrograman Fungsional"},
    }

    content, err := helpers.SaveToJSON(Assignments)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("assignments.json", content)

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
  LENGTH := len(Assignments)

  for i := 1; i < LENGTH; i++ {
    key := Assignments[i]
    j := i - 1

    for j >= 0 && Assignments[j].ID > key.ID {
      Assignments[j + 1] = Assignments[j]
      j = j - 1
    }

    Assignments[j + 1] = key
  }

  return Assignments[LENGTH - 1].ID + 1
}

/**
  * Get all assignments
  *
  * @return []Assignment
  */
func GetAll() []Assignment {
  return Assignments
}

/**
  * Get all assignments by class ID
  *
  * @param int classID
  * @param enums.SORT sortType
  * @return []Assignment
  */
func GetAllByClassID(classID int, sortType enums.SORT) []Assignment {
  LENGTH := len(Assignments)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Assignments[j].ID < Assignments[min].ID {
          min = j
        }
      } else {
        if Assignments[j].ID > Assignments[min].ID {
          min = j
        }
      }
    }

    Assignments[i], Assignments[min] = Assignments[min], Assignments[i]
  }

  var assignments []Assignment

  for _, assignment := range Assignments {
    if assignment.ClassID == classID {
      assignments = append(assignments, assignment)
    }
  }

  return assignments
}

/**
  * Get all assignments by course ID
  *
  * @param int courseID
  * @param enums.SORT sortType
  * @return []Assignment
  */
func GetAllByCourseID(courseID int, sortType enums.SORT) []Assignment {
  LENGTH := len(Assignments)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Assignments[j].ID < Assignments[min].ID {
          min = j
        }
      } else {
        if Assignments[j].ID > Assignments[min].ID {
          min = j
        }
      }
    }

    Assignments[i], Assignments[min] = Assignments[min], Assignments[i]
  }

  var assignments []Assignment

  for _, assignment := range Assignments {
    if assignment.CourseID == courseID {
      assignments = append(assignments, assignment)
    }
  }

  return assignments
}

/**
  * Inserts an assignment
  *
  * @param Assignment assignment
  * @return void
  */
func InsertAssignment(assignment Assignment) {
  assignment.ID = CreateNewID()
  Assignments = append(Assignments, assignment)

  content, err := helpers.SaveToJSON(Assignments)

  if err != nil {
    panic(err)
  }

  err = helpers.SaveFile("assignments.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Get course by assignment
  *
  * @param Assignment assignment
  * @return *course.Course
  */
func GetCourseByAssignment(assignment Assignment) *course.Course {
  return course.FindCourseByID(assignment.CourseID)
}

/**
  * Get class by assignment
  *
  * @param Assignment assignment
  * @return *class.Class
  */
func GetClassByAssignment(assignment Assignment) *class.Class {
  return class.FindClassByID(assignment.ClassID)
}

/**
  * Find an assignment by their ID using binary search
  *
  * @param int ID
  * @return *Assignment
  */
func FindAssignmentByID(ID int) *Assignment {
  if ID < 1 {
    return &Assignment{}
  }

  left := 0
  right := len(Assignments) - 1

  for left <= right {
    mid := left + (right - left) / 2

    if Assignments[mid].ID == ID {
      return &Assignments[mid]
    }

    if Assignments[mid].ID < ID {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return &Assignment{}
}

/**
  * Find an assignment by their title using sequential search
  *
  * @param string title
  * @return *Assignment
  */
func FindAssignmentByTitle(title string) *Assignment {
  for _, assignment := range Assignments {
    if assignment.Title == title {
      return &assignment
    }
  }

  return &Assignment{}
}

/**
  * Deletes an assignment by their ID using sequential search
  *
  * @param int ID
  * @return void
  */
func DeleteAssignment(ID int) {
  var newAssignments []Assignment

  for i := 0; i < len(Assignments); i++ {
    if Assignments[i].ID != ID {
      newAssignments = append(newAssignments, Assignments[i])
    }
  }

  Assignments = newAssignments

  content, err := helpers.SaveToJSON(Assignments)

  if err != nil {
    panic(err)
  }

  err = helpers.SaveFile("assignments.json", content)

  if err != nil {
    panic(err)
  }
}
