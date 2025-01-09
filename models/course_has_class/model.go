package course_has_class

import (
	"modol-app/helpers"
	"modol-app/models/class"
	"modol-app/models/course"
)

/**
 * CourseHasClass struct
 */
type CourseHasClass struct {
	ID                  int
	CourseID            int
	ClassID             int
}

// CourseHasClasses collection
var CourseHasClasses []CourseHasClass

/**
 * Initializes the CourseHasClasses collection
 *
 * @return void
 */
func init() {
	if helpers.FileExists("course_has_classes.json") {
		content, err := helpers.ReadFile("course_has_classes.json")

		if err != nil {
			panic(err)
		}

		helpers.LoadFromJSON(content, &CourseHasClasses)
	} else {
		CourseHasClasses = []CourseHasClass{
			{1, 1, 8},
			{2, 2, 7},
			{3, 3, 6},
			{4, 4, 5},
			{5, 5, 4},
			{6, 6, 3},
			{7, 7, 2},
			{8, 8, 1},
		}

		content, err := helpers.SaveToJSON(CourseHasClasses)

		if err != nil {
			panic(err)
		}

		err = helpers.SaveFile("course_has_classes.json", content)

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
	LENGTH := len(CourseHasClasses)

	for i := 1; i < LENGTH; i++ {
		key := CourseHasClasses[i]
		j := i - 1

		for j >= 0 && CourseHasClasses[j].ID > key.ID {
			CourseHasClasses[j+1] = CourseHasClasses[j]
			j = j - 1
		}

		CourseHasClasses[j+1] = key
	}

	return CourseHasClasses[LENGTH-1].ID + 1
}

/**
 * Assigns a course to a class
 *
 * @param int courseID
 * @param int classID
 * @return void
 */
func AssignCourseToClass(courseID, classID int) {
	data := CourseHasClass{CreateNewID(), courseID, classID}
	CourseHasClasses = append(CourseHasClasses, data)

	content, err := helpers.SaveToJSON(CourseHasClasses)

	if err != nil {
		panic(err)
	}

	err = helpers.UpdateFile("course_has_classes.json", content)

	if err != nil {
		panic(err)
	}
}

/**
 * Lists all courses assigned to a class
 *
 * @param int classID
 * @return []course.Course
 */
func ListCoursesAssignedToClass(classID int) []course.Course {
	var courses []course.Course

	for _, courseHasClass := range CourseHasClasses {
		if courseHasClass.ClassID == classID {
			courses = append(courses, *course.FindCourseByID(courseHasClass.CourseID))
		}
	}

	if len(courses) == 0 {
		return []course.Course{}
	}

	return courses
}

/**
 * Lists all classes assigned to a course
 *
 * @param int courseID
 * @return []class.Class
 */
func ListClassesAssignedToCourse(courseID int) []class.Class {
	var classes []class.Class

	for _, courseHasClass := range CourseHasClasses {
		if courseHasClass.CourseID == courseID {
			classes = append(classes, *class.FindClassByID(courseHasClass.ClassID))
		}
	}

	if len(classes) == 0 {
		return []class.Class{}
	}

	return classes
}
