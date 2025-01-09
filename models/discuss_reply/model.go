package discuss_reply

import (
  "modol-app/enums"
  "modol-app/helpers"
)

/**
  * DiscussReply struct
  */
type DiscussReply struct {
  ID          int
  DiscussID   int
  UserID      int
  Content     string
}

// DiscussReplies collection
var DiscussReplies []DiscussReply

/**
  * Initializes the DiscussReplies collection
  *
  * @return void
  */
func init() {
  if helpers.FileExists("discuss_replies.json") {
    content, err := helpers.ReadFile("discuss_replies.json")

    if err != nil {
      panic(err)
    }

    helpers.LoadFromJSON(content, &DiscussReplies)
  } else {
    DiscussReplies = []DiscussReply{
      {1, 1, 2, "Halo juga, saya baik."},
    }

    content, err := helpers.SaveToJSON(DiscussReplies)

    if err != nil {
      panic(err)
    }

    err = helpers.SaveFile("discuss_replies.json", content)

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
  LENGTH := len(DiscussReplies)

  for i := 1; i < LENGTH; i++ {
    key := DiscussReplies[i]
    j := i - 1

    for j >= 0 && DiscussReplies[j].ID > key.ID {
      DiscussReplies[j + 1] = DiscussReplies[j]
      j = j - 1
    }

    DiscussReplies[j + 1] = key
  }

  return DiscussReplies[LENGTH - 1].ID + 1
}

/**
  * Lists all discuss replies by discuss ID with sequential search
  *
  * @param int discussID
  * @param enums.SORT sortType
  * @return []DiscussReply
  */
func ListDiscussRepliesByDiscussID(discussID int, sortType enums.SORT) []DiscussReply {
  var discussReplies []DiscussReply

  for _, discussReply := range DiscussReplies {
    if discussReply.DiscussID == discussID {
      discussReplies = append(discussReplies, discussReply)
    }
  }

  LENGTH := len(discussReplies)

  for i := 0; i < LENGTH - 1; i++ {
    min := i

    for j := i + 1; j < LENGTH; j++ {
      if sortType == enums.ASC {
        if discussReplies[j].ID < discussReplies[min].ID {
          min = j
        }
      } else {
        if discussReplies[j].ID > discussReplies[min].ID {
          min = j
        }
      }
    }

    discussReplies[i], discussReplies[min] = discussReplies[min], discussReplies[i]
  }

  return discussReplies
}

/**
  * Inserts a discuss reply
  *
  * @param DiscussReply discussReply
  * @return void
  */
func InsertDiscussReply(discussReply DiscussReply) {
  discussReply.ID = CreateNewID()
  DiscussReplies = append(DiscussReplies, discussReply)

  content, err := helpers.SaveToJSON(DiscussReplies)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("discuss_replies.json", content)

  if err != nil {
    panic(err)
  }
}

/**
  * Deletes a discuss reply by its ID
  *
  * @param int discussReplyID
  * @return void
  */
func DeleteDiscussReplyByID(discussReplyID int) {
  var newDiscussReplies []DiscussReply

  for i := 0; i < len(DiscussReplies); i++ {
    if DiscussReplies[i].ID != discussReplyID {
      newDiscussReplies = append(newDiscussReplies, DiscussReplies[i])
    }
  }

  DiscussReplies = newDiscussReplies

  content, err := helpers.SaveToJSON(DiscussReplies)

  if err != nil {
    panic(err)
  }

  err = helpers.UpdateFile("discuss_replies.json", content)

  if err != nil {
    panic(err)
  }
}