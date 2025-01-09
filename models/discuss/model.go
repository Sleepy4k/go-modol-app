package discuss

import (
  "modol-app/enums"
  "modol-app/helpers"
  "modol-app/models/discuss_reply"
)

/**
  * Discuss struct
  */
type Discuss struct {
  ID          int
  UserID      int
  Content     string
}

// Discusses collection
var Discusses []Discuss

/**
  * Initializes the Discusses collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("discusses.json") {
    content, err := helpers.ReadFile("discusses.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &Discusses)
  } else {
    Discusses = []Discuss{
      {1, 3, "Halo, apa kabar?"},
    }

    content, err := helpers.SaveToJSON(Discusses)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("discusses.json", content)

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
  LENGTH := len(Discusses)

  for i := 1; i < LENGTH; i++ {
    key := Discusses[i]
    j := i - 1

    for j >= 0 && Discusses[j].ID > key.ID {
      Discusses[j + 1] = Discusses[j]
      j = j - 1
    }

    Discusses[j + 1] = key
  }

  return Discusses[LENGTH - 1].ID + 1
}

/**
  * Lists all discusses with selection sort
  *
  * @param enums.SORT sortType
  * @return []Discuss
  */
func ListDiscusses(sortType enums.SORT) []Discuss {
  LENGTH := len(Discusses)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if Discusses[j].ID < Discusses[min].ID {
          min = j
        }
      } else {
        if Discusses[j].ID > Discusses[min].ID {
          min = j
        }
      }
    }

    Discusses[i], Discusses[min] = Discusses[min], Discusses[i]
  }

  return Discusses
}

/**
  * Inserts a discuss
  *
  * @param Discuss discuss
  * @return void
  */
func InsertDiscuss(discuss Discuss) {
  discuss.ID = CreateNewID()
  Discusses = append(Discusses, discuss)

  content, err := helpers.SaveToJSON(Discusses)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("discusses.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Lists all discusses by user ID using sequential search
  *
  * @param int userID
  * @return []Discuss
  */
func ListDiscussesByUserID(userID int) []Discuss {
  var result []Discuss

  for _, discuss := range Discusses {
    if discuss.UserID == userID {
      result = append(result, discuss)
    }
  }

  return result
}

/**
  * Lists all discusses replies by discuss ID
  *
  * @param int discussID
  * @param enums.SORT sortType
  * @return []DiscussReply
  */
func ListDiscussesRepliesByDiscussID(discussID int, sortType enums.SORT) []discuss_reply.DiscussReply {
  return discuss_reply.ListDiscussRepliesByDiscussID(discussID, sortType)
}

/**
  * Finds a discuss by their ID using binary search
  *
  * @param int ID
  * @return *Discuss
  */
func FindDiscussByID(ID int) *Discuss {
  if ID < 1 {
    return &Discuss{}
  }

  left := 0
  right := len(Discusses) - 1

  for left <= right {
    mid := left + (right - left) / 2

    if Discusses[mid].ID == ID {
      return &Discusses[mid]
    }

    if Discusses[mid].ID < ID {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return &Discuss{}
}

/**
  * Deletes a discuss by their ID using sequential search
  *
  * @param int ID
  * @return void
  */
func DeleteDiscuss(ID int) {
  var newDiscusses []Discuss

  for i := 0; i < len(Discusses); i++ {
    if Discusses[i].ID != ID {
      newDiscusses = append(newDiscusses, Discusses[i])
    }
  }

  Discusses = newDiscusses

  content, err := helpers.SaveToJSON(Discusses)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("discusses.json", content)

  if err != nil {
    panic(err)
  }
}
