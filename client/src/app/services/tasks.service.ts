import { Injectable } from '@angular/core';
import { HttpHandlerService } from './http-handler.service';
import { ProjectService } from './project.service';
import { ProjectTask } from '../interfaces/project-task';
import { TaskState } from 'src/common/task-state';
import { Project } from '../interfaces/project';
import { catchError, Observable, of, Subject } from 'rxjs';
import { TasksStats } from '../interfaces/tasks-stats';
import { TaskComment } from '../interfaces/task-comment';
import { ErrorShown } from 'src/common/error-shown';

@Injectable({
    providedIn: 'root',
})
export class TasksService {
    stateTasks: Map<TaskState, ProjectTask[]>;
    recentTasks: ProjectTask[];

    currentTask: ProjectTask;
    currentComments: TaskComment[];
    newTaskSetObservable: Subject<ProjectTask>;

    constructor(private http: HttpHandlerService, private projectService: ProjectService) {
        this.stateTasks = new Map();
        this.stateTasks.set(TaskState.TODO, []);
        this.stateTasks.set(TaskState.ONGOING, []);
        this.stateTasks.set(TaskState.DONE, []);

        this.recentTasks = [];
        this.currentComments = [];
        this.currentTask = {} as ProjectTask;

        this.newTaskSetObservable = new Subject();
    }

    fetchStateTasks() {
        if (!this.projectService.currentProject) return;
        this.http
            .getTasksByState(this.projectService.currentProject.id, TaskState.TODO)
            .pipe(catchError(this.http.handleError(ErrorShown.TasksUnfetchable, [])))
            .subscribe((data) => this.stateTasks.set(TaskState.TODO, [...data]));
        this.http
            .getTasksByState(this.projectService.currentProject.id, TaskState.ONGOING)
            .pipe(catchError(this.http.handleError(ErrorShown.TasksUnfetchable, [])))
            .subscribe((data) => this.stateTasks.set(TaskState.ONGOING, [...data]));
        this.http
            .getTasksByState(this.projectService.currentProject.id, TaskState.DONE)
            .pipe(catchError(this.http.handleError(ErrorShown.TasksUnfetchable, [])))
            .subscribe((data) => this.stateTasks.set(TaskState.DONE, [...data]));
    }

    fetchRecentTasks() {
        if (!this.projectService.currentProject) return;
        this.http
            .getRecentTasks(this.projectService.currentProject.id)
            .pipe(catchError(this.http.handleError(ErrorShown.TasksUnfetchable, [])))
            .subscribe((tasks) => (this.recentTasks = [...tasks]));
    }

    getTasksStats(): Observable<TasksStats> {
        if (!this.projectService.currentProject) return of();
        return this.http.getTasksStats(this.projectService.currentProject.id);
    }

    setCurrentTask(task: ProjectTask) {
        this.currentTask = task;
        this.newTaskSetObservable.next(task);
    }

    clearCurrentTask() {
        this.currentTask = {} as ProjectTask;
    }

    uploadTask(task: ProjectTask) {
        if (!this.projectService.currentProject?.id) return;
        this.http
            .createTask(task, (this.projectService.currentProject as Project).id)
            .pipe(catchError(this.http.handleError(ErrorShown.TaskUploadFailed)))
            .subscribe(() => this.fetchStateTasks());
    }

    updateTask(task: ProjectTask) {
        this.http
            .updateTask(task, (this.projectService.currentProject as Project).id)
            .pipe(catchError(this.http.handleError(ErrorShown.TaskUpdateFailed)))
            .subscribe(() => {
                this.setCurrentTask(task);
                this.fetchStateTasks();
            });
    }

    updateTaskPosition(previousIndex: number, currentIndex: number, taskId: number) {
        if (!this.projectService.currentProject) return;
        this.http
            .updateTaskPosition(previousIndex, currentIndex, taskId, this.projectService.currentProject.id)
            .pipe(catchError(this.http.handleError(ErrorShown.TaskPositionUpdateFailed)))
            .subscribe();
    }

    updateTaskState(newState: TaskState, currentIndex: number, taskId: number) {
        if (!this.projectService.currentProject) return;
        (this.stateTasks.get(newState) as ProjectTask[])[currentIndex].state = newState;
        this.http
            .updateTaskState(newState, currentIndex, taskId, this.projectService.currentProject.id)
            .pipe(catchError(this.http.handleError(ErrorShown.TaskUpdateStateFailed)))
            .subscribe();
    }

    deleteTask(taskId: number) {
        if (!this.projectService.currentProject) return;
        this.http
            .deleteTask(taskId, this.projectService.currentProject.id)
            .pipe(catchError(this.http.handleError(ErrorShown.TaskDeleteFailed)))
            .subscribe(() => {
                if (taskId === this.currentTask.id) this.setCurrentTask({} as ProjectTask);
                this.fetchStateTasks();
            });
    }

    fetchComments() {
        if (!this.projectService.currentProject) return;
        this.http
            .getTaskComments(this.currentTask.id, this.projectService.currentProject.id)
            .pipe(catchError(this.http.handleError(ErrorShown.CommentsUnfetchable, [])))
            .subscribe((comments) => (this.currentComments = [...comments]));
    }

    addComment(content: string) {
        if (!this.projectService.currentProject) return;
        this.http.addTaskComment(this.currentTask.id, content, this.projectService.currentProject.id)
        .pipe(catchError(this.http.handleError(ErrorShown.CommentSendFailed)))
        .subscribe(() => this.fetchComments());
    }
}
