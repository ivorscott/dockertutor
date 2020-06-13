package lessons

type ActiveLesson int
type ActiveCategory string

// Categories represents category of lessons
type Categories []string{"docker","docker-compose","swarm"}

// Lesson represents the state of an exercise
type Lesson struct {
	Id int
	Title string
	Exercise string
	Answer string
	Explanation string
	Complete bool
	AutoClean bool
	DependsOn []Lesson
	Resources
}

type Lessons []Lesson

// Resources aim to perserve the state of a lesson in the event of 
// exiting the program and starting back up again. Direct resources
// and the resources of dependent lessons 
type Resources struct {
	Images []string
	Containers []string
	Volumes	[]string
	Networks []string
}

// Next fetches the next lesson
func (l *Lesson) Next() {}

// Teach displays the lesson to the user
func (l *Lesson) Teach() {}

// Success represents a lesson succeeded
func (l *Lesson) Success() {}

// Failure represents a lesson failed
func (l *Lesson) Failure() {}

// Exit quits the lesson
func (l *Lesson) Exit() {}

// Reset resources and lesson progress
func (l *Lesson) Reset() {}