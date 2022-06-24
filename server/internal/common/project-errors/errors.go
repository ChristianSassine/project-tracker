package projectErrors

import "errors"

var (
	FailedAuthToken         error = errors.New("Authentication token failed, declining authentication")
	FailedRefreshToken            = errors.New("Token refresh failed, no new authentication token sent to the client")
	FailedToLogin                 = errors.New("Failed to login the client, no tokens were sent")
	FailedToLogout                = errors.New("Failed to logout the client, the cookies containing the tokens weren't cleared")
	FailedToCreateUser            = errors.New("Failed to create the user, no changes were made to the database")
	FailedToFetchUsername         = errors.New("Failed to fetch the user's username, no username was sent")
	FailedToAddComment            = errors.New("Failed to add a comment, no comment was added to the database")
	FailedToFetchComments         = errors.New("Failed to get the comments from the database, no comment was sent to the user")
	FailedToFetchLogs             = errors.New("Failed to get the logs of the project from the database, no logs sent to the user")
	FailedToFetchProjects         = errors.New("Failed to get the projects of the user from the database, no projects were sent to the user")
	FailedToAddProject            = errors.New("Failed to add a project to the database, no projects were added to the database")
	FailedToAddUserProject        = errors.New("Failed to add user to a project, the user failed to join the project and the database was not altered")
	FailedToDeleteProject         = errors.New("Failed to delete a project, the project remains in the database")
	FailedToFetchTasks            = errors.New("Failed to get the tasks of the project from the database, no tasks were sent to the user")
	FailedToFetchTasksStats       = errors.New("Failed to get the stats of the project from the database, no stats were sent to the user")
	FailedToAddTask               = errors.New("Failed to add a task, no task was added to the database")
	FailedToUpdateTask            = errors.New("Failed to update a task, no task was updated in the database")
	FailedToDeleteTask            = errors.New("Failed to delete a task, no task was removed from the database")
	FailedToValidateUser          = errors.New("Failed to validate user in the project, declining request")
)
