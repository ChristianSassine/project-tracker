import { Injectable } from '@angular/core';
import { HttpHandlerService } from './http-handler.service';
import { ProjectService } from './project.service';
import { ProjectTask } from '../interfaces/project-task';

@Injectable({
    providedIn: 'root',
})
export class TasksService {
    tasks : ProjectTask[]
	constructor(private http: HttpHandlerService, private projectService: ProjectService) {}

	fetchTasks(){
		if (!this.projectService.currentProject) return;
		this.http.getAllTasks(this.projectService.currentProject.id).subscribe((data)=> this.tasks = [...data])
	}
}
