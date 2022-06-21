export enum ErrorShown {
    ServerError = 'An error occured while contacting the server',
    LogsUnfetchable = 'Failed to fetch the logs from the server',
    ProjectsUnfetchable = 'Failed to fetch the projects from the server',
    ProjectUncreatable = 'Failed to create the project in the server',
    ProjectUnjoinable = 'Failed to join the project in the server',
    ProjectNotDeleted = 'Failed to delete the project from the server',
    TasksUnfetchable = 'Failed to fetch the tasks from the server',
    TaskUploadFailed = 'Failed to upload the task to the server',
    TaskUpdateFailed = 'Failed to update the task in the server',
    TaskUpdateStateFailed = 'Failed to update the state of the task in the server',
    TaskPositionUpdateFailed = 'Failed to update the position of the task in the server',
    TaskDeleteFailed = 'Failed to delete the task in the server',
    CommentsUnfetchable = "Failed to fetch the task's comments from the server",
    CommentSendFailed = 'Failed to send the comment to the server',
    LogoutFailed = 'Failed to logout'
}