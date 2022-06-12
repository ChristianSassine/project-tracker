import { Injectable } from '@angular/core';
import { HttpHandlerService } from './http-handler.service';
import { ProjectService } from './project.service';
import { ProjectTask } from '../interfaces/project-task';
import { TaskState } from 'src/common/task-state';
import { Project } from '../interfaces/project';
import { Subject } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class TasksService {
    stateTasks: Map<TaskState, ProjectTask[]>;
    recentTasks: ProjectTask[];

    currentTask: ProjectTask;
    newTaskSetObservable: Subject<ProjectTask>;

    constructor(private http: HttpHandlerService, private projectService: ProjectService) {
        this.stateTasks = new Map();
        this.stateTasks.set(TaskState.TODO, []);
        this.stateTasks.set(TaskState.ONGOING, []);
        this.stateTasks.set(TaskState.DONE, []);
        this.recentTasks = [];

        this.currentTask = {} as ProjectTask;

        this.newTaskSetObservable = new Subject();
    }

    fetchStateTasks() {
        if (!this.projectService.currentProject) return;
        this.http
            .getTasksByState(this.projectService.currentProject.id, TaskState.TODO)
            .subscribe((data) => this.stateTasks.set(TaskState.TODO, [...data]));
        this.http
            .getTasksByState(this.projectService.currentProject.id, TaskState.ONGOING)
            .subscribe((data) => this.stateTasks.set(TaskState.ONGOING, [...data]));
        this.http
            .getTasksByState(this.projectService.currentProject.id, TaskState.DONE)
            .subscribe((data) => this.stateTasks.set(TaskState.DONE, [...data]));
    }

    fetchRecentTasks() {
        if (!this.projectService.currentProject) return;
        this.http.getRecentTasks(this.projectService.currentProject.id).subscribe((tasks) => (this.recentTasks = [...tasks]));
    }

    setCurrentTask(task: ProjectTask) {
        this.currentTask = task;
        this.newTaskSetObservable.next(task);
    }

    uploadTask(task: ProjectTask) {
        if (!this.projectService.currentProject?.id) return;
        this.http.createTask(task, (this.projectService.currentProject as Project).id).subscribe(() => this.fetchStateTasks());
    }

    updateTask(task: ProjectTask) {
        this.http.updateTask(task, (this.projectService.currentProject as Project).id).subscribe(() => {
            this.setCurrentTask(task);
            this.fetchStateTasks();
        });
    }

    updateTaskPosition(previousIndex: number, currentIndex: number, taskId: number) {
        if (!this.projectService.currentProject) return;
        this.http.updateTaskPosition(previousIndex, currentIndex, taskId, this.projectService.currentProject.id).subscribe();
    }

    updateTaskState(newState: TaskState, currentIndex: number, taskId: number) {
        if (!this.projectService.currentProject) return;
        (this.stateTasks.get(newState) as ProjectTask[])[currentIndex].state = newState;
        this.http.updateTaskState(newState, currentIndex, taskId, this.projectService.currentProject.id).subscribe();
    }

    deleteTask(taskId: number) {
        this.http.deleteTask(taskId, (this.projectService.currentProject as Project).id).subscribe(() => {
            if (taskId === this.currentTask.id) this.setCurrentTask({} as ProjectTask);
            this.fetchStateTasks();
        });
    }
}
