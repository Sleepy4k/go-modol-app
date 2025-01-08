package class

import (
  "modol-app/enums"
  "modol-app/helpers"
)

// User model
type Class struct {
  ID          int
  Nama        string
}

// Classes collection
var Classes []Class

/**
  * Initializes the classes collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("classes.json") {
    content, err := helpers.ReadFile("classes.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Classes)
  } else {
    Classes = []Class{
      {1, "IF 11 01"},
      {2, "IF 11 02"},
      {3, "IF 11 03"},
      {4, "IF 11 04"},
      {5, "IF 11 05"},
      {6, "IF 11 06"},
      {7, "IF 11 07"},
      {8, "IF 11 08"},
    }

    content, err := helpers.SaveToJSON(Classes)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("classes.json", content)

    if err != nil {
      panic(err)
    }
  }
}

/**
  * Lists all classes with selection sort
  *
  * @param enums.SORT sortType
  * @return []Class
  */
func ListClasses(sortType enums.SORT) []Class {
  for i := 0; i < len(Classes); i++ {
    min := i

    for j := i + 1; j < len(Classes); j++ {
      if sortType == enums.ASC {
        if Classes[j].ID < Classes[min].ID {
          min = j
        }
      } else {
        if Classes[j].ID > Classes[min].ID {
          min = j
        }
      }
    }

    Classes[i], Classes[min] = Classes[min], Classes[i]
  }

  return Classes
}

/**
  * Inserts a class
  *
  * @param Class class
  * @return void
  */
func InsertClass(class Class) {
  class.ID = len(Classes) + 1
  Classes = append(Classes, class)

  content, err := helpers.SaveToJSON(Classes)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("classes.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Finds a user by their ID
  *
  * @param string ID
  * @return *Class
  */
func FindUserByID(ID int) *Class {
  if ID < 1 {
    return &Class{}
  }

  low := 0
  high := len(Classes) - 1

  for low <= high {
    mid := (low + high) / 2

    if Classes[mid].ID == ID {
      return &Classes[mid]
    }

    if Classes[mid].ID < ID {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }

  return &Class{}
}

/**
  * Deletes a class by their ID
  *
  * @param int ID
  * @return void
  */
func DeleteClass(ID int) {
  var newClasses []Class

  for i := 0; i < len(Classes); i++ {
    if Classes[i].ID != ID {
      newClasses = append(newClasses, Classes[i])
    }
  }

  Classes = newClasses

  content, err := helpers.SaveToJSON(Classes)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("classes.json", content)

  if err != nil {
    panic(err)
  }
}