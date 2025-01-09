package subject

import (
  "modol-app/enums"
  "modol-app/helpers"
)

/**
  * Subject struct
  */
type Subject struct {
  ID          int
  Name        string
  Description string
}

// Subjects collection
var Subjects []Subject

/**
  * Initializes the Subjects collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("subjects.json") {
    content, err := helpers.ReadFile("subjects.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Subjects)
  } else {
    Subjects = []Subject{
      {1, "Matematika", "Pelajaran tentang angka."},
      {2, "Bahasa Indonesia", "Pelajaran tentang bahasa Indonesia."},
      {3, "Bahasa Inggris", "Pelajaran tentang bahasa Inggris."},
    }

    content, err := helpers.SaveToJSON(Subjects)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("subjects.json", content)

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
  LENGTH := len(Subjects)

  for i := 1; i < LENGTH; i++ {
    key := Subjects[i]
    j := i - 1

    for j >= 0 && Subjects[j].ID > key.ID {
      Subjects[j + 1] = Subjects[j]
      j = j - 1
    }

    Subjects[j + 1] = key
  }

  return Subjects[LENGTH - 1].ID + 1
}

/**
  * Finds a subject by ID using sequential search
  *
  * @param int id
  * @return *Subject
  */
func FindSubjectByID(id int) *Subject {
  for _, subject := range Subjects {
    if subject.ID == id {
      return &subject
    }
  }

  return &Subject{}
}

/**
  * Lists all subjects with sequential search
  *
  * @param enums.SORT sortType
  * @return []Subject
  */
func ListSubjects(sortType enums.SORT) []Subject {
  LENGTH := len(Subjects)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Subjects[j].Name < Subjects[min].Name {
          min = j
        }
      } else {
        if Subjects[j].Name > Subjects[min].Name {
          min = j
        }
      }
    }

    Subjects[i], Subjects[min] = Subjects[min], Subjects[i]
  }

  return Subjects
}

/**
  * Inserts a subject
  *
  * @param Subject subject
  * @return void
  */
func InsertSubject(subject Subject) {
  subject.ID = CreateNewID()
  Subjects = append(Subjects, subject)

  content, err := helpers.SaveToJSON(Subjects)

  if err != nil {
    panic(err)
  }

  err = helpers.SaveFile("subjects.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Updates a subject
  *
  * @param Subject subject
  * @return void
  */
func UpdateSubject(subject Subject) {
  for index, item := range Subjects {
    if item.ID == subject.ID {
      Subjects[index] = subject
    }
  }

  content, err := helpers.SaveToJSON(Subjects)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("subjects.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Deletes a subject by ID
  *
  * @param int id
  * @return void
  */
func DeleteSubjectByID(id int) {
  var newSubjects []Subject

  for i := 0; i < len(Subjects); i++ {
    if Subjects[i].ID != id {
      newSubjects = append(newSubjects, Subjects[i])
    }
  }

  Subjects = newSubjects

  content, err := helpers.SaveToJSON(Subjects)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("subjects.json", content)

  if err != nil {
    panic(err)
  }
}
