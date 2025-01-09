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
  * Creates a new ID using insertion sort
  *
  * @return int
  */
func CreateNewID() int {
  LENGTH := len(Classes)

  for i := 1; i < LENGTH; i++ {
    key := Classes[i]
    j := i - 1

    for j >= 0 && Classes[j].ID > key.ID {
      Classes[j + 1] = Classes[j]
      j = j - 1
    }

    Classes[j + 1] = key
  }

  return Classes[LENGTH - 1].ID + 1
}

/**
  * Lists all classes with selection sort
  *
  * @param enums.SORT sortType
  * @return []Class
  */
func ListClasses(sortType enums.SORT) []Class {
  LENGTH := len(Classes)

  for i := 0; i < LENGTH; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Classes[j].Nama < Classes[min].Nama {
          min = j
        }
      } else {
        if Classes[j].Nama > Classes[min].Nama {
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
  class.ID = CreateNewID()
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
  * Finds a class by their ID using binary search
  *
  * @param string ID
  * @return *Class
  */
func FindClassByID(ID int) *Class {
  if ID < 1 {
    return &Class{}
  }

  left := 0
  right := len(Classes) - 1

  for left <= right {
    mid := left + (right - left) / 2

    if Classes[mid].ID == ID {
      return &Classes[mid]
    }

    if Classes[mid].ID < ID {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return &Class{}
}

/**
  * Updates a class by their ID using sequential search
  *
  * @param int ID
  * @param Class data
  * @return void
  */
func UpdateClass(ID int, data Class) {
  for i := 0; i < len(Classes); i++ {
    if Classes[i].ID == ID {
      Classes[i].Nama = data.Nama
    }
  }

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
  * Deletes a class by their ID using sequential search
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