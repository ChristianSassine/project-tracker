import { Injectable } from '@angular/core';
import { HttpHandlerService } from './http-handler.service';
import { ProjectService } from './project.service';
import { ProjectTask } from '../interfaces/project-task';
import { TaskState } from 'src/common/task-state';
import { Project } from '../interfaces/project';

@Injectable({
    providedIn: 'root',
})
export class TasksService {
    tasksTODO: ProjectTask[];
    tasksONGOING: ProjectTask[];
    tasksDONE: ProjectTask[];

    constructor(private http: HttpHandlerService, private projectService: ProjectService) {}

    fetchStateTasks() {
        if (!this.projectService.currentProject) return;
        this.http.getTasksByState(this.projectService.currentProject.id, TaskState.TODO).subscribe((data) => (this.tasksTODO = [...data]));
        this.http.getTasksByState(this.projectService.currentProject.id, TaskState.ONGOING).subscribe((data) => (this.tasksONGOING = [...data]));
        this.http.getTasksByState(this.projectService.currentProject.id, TaskState.DONE).subscribe((data) => (this.tasksDONE = [...data]));
    }

    fetchTasksByState(){
        if (!this.projectService.currentProject) return;
        this.http.getTasksByState(this.projectService.currentProject.id, 'TODO').subscribe();
    }

    uploadTask(task: ProjectTask) {
        if (!this.projectService.currentProject?.id) return;
        this.http.createTask(task, this.projectService.currentProject?.id as number).subscribe();
    }

    updateTask(task: ProjectTask){
        this.http.updateTask(task, (this.projectService.currentProject as Project).id)
    }
}
