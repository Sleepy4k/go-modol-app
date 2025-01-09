package class_has_user

import (
	"modol-app/helpers"
	"modol-app/models/class"
	"modol-app/models/user"
)

/**
 * ClassHasUsers struct
 */
type ClassHasUser struct {
	ID                int
	ClassID           int
	UserID            int
}

// ClassHasUsers collection
var ClassHasUsers []ClassHasUser

/**
 * Initializes the ClassHasUsers collection
 *
 * @return void
 */
func init() {
	if helpers.FileExists("class_has_users.json") {
		content, err := helpers.ReadFile("class_has_users.json")

		if err != nil {
			panic(err)
		}

		helpers.LoadFromJSON(content, &ClassHasUsers)
	} else {
		ClassHasUsers = []ClassHasUser{
			{1, 8, 3},
		}

		content, err := helpers.SaveToJSON(ClassHasUsers)

		if err != nil {
			panic(err)
		}

		err = helpers.SaveFile("class_has_users.json", content)

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
	LENGTH := len(ClassHasUsers)

	for i := 1; i < LENGTH; i++ {
		key := ClassHasUsers[i]
		j := i - 1

		for j >= 0 && ClassHasUsers[j].ID > key.ID {
			ClassHasUsers[j+1] = ClassHasUsers[j]
			j = j - 1
		}

		ClassHasUsers[j+1] = key
	}

	return ClassHasUsers[LENGTH-1].ID + 1
}

/**
 * Get all class has users
 *
 * @return []ClassHasUser
 */
func GetAll() []ClassHasUser {
	return ClassHasUsers
}

/**
 * Get all class has users by class ID
 *
 * @param int classID
 * @return []ClassHasUser
 */
func GetAllByClassID(classID int) []ClassHasUser {
	var classHasUsers []ClassHasUser

	for _, classHasUser := range ClassHasUsers {
		if classHasUser.ClassID == classID {
			classHasUsers = append(classHasUsers, classHasUser)
		}
	}

	return classHasUsers
}

/**
 * Get all class has users by user ID
 *
 * @param int userID
 * @return []ClassHasUser
 */
func GetAllByUserID(userID int) []ClassHasUser {
	var classHasUsers []ClassHasUser

	for _, classHasUser := range ClassHasUsers {
		if classHasUser.UserID == userID {
			classHasUsers = append(classHasUsers, classHasUser)
		}
	}

	return classHasUsers
}

/**
 * Assign a user to a class
 *
 * @param int classID
 * @param int userID
 * @return void
 */
func AssignUserToClass(classID, userID int) {
	data := ClassHasUser{CreateNewID(), classID, userID}
	ClassHasUsers = append(ClassHasUsers, data)

	content, err := helpers.SaveToJSON(ClassHasUsers)

	if err != nil {
		panic(err)
	}

	err = helpers.UpdateFile("class_has_users.json", content)

	if err != nil {
		panic(err)
	}
}

/**
 * Lists all users assigned to a class
 *
 * @param int classID
 * @return []user.User
 */
func ListUsersAssignedToClass(classID int) []user.User {
	var users []user.User

	for _, classHasUser := range ClassHasUsers {
		if classHasUser.ClassID == classID {
			users = append(users, *user.FindUserByID(classHasUser.UserID))
		}
	}

	return users
}

/**
 * Lists all classes assigned to a user
 *
 * @param int userID
 * @return []class.Class
 */
func ListClassesAssignedToUser(userID int) []class.Class {
	var classes []class.Class

	for _, classHasUser := range ClassHasUsers {
		if classHasUser.UserID == userID {
			classes = append(classes, *class.FindClassByID(classHasUser.ClassID))
		}
	}

	return classes
}

/**
 * Unassign a user from a class
 *
 * @param int classID
 * @param int userID
 * @return void
 */
func UnassignUserFromClass(classID, userID int) {
	for i, classHasUser := range ClassHasUsers {
		if classHasUser.ClassID == classID && classHasUser.UserID == userID {
			ClassHasUsers = append(ClassHasUsers[:i], ClassHasUsers[i+1:]...)
		}
	}

	content, err := helpers.SaveToJSON(ClassHasUsers)

	if err != nil {
		panic(err)
	}

	err = helpers.UpdateFile("class_has_users.json", content)

	if err != nil {
		panic(err)
	}
}
