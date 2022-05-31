import { Component, OnInit } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-tasks',
    templateUrl: './tasks.component.html',
    styleUrls: ['./tasks.component.scss'],
})
export class TasksComponent implements OnInit {
    constructor(private tasksService: TasksService,) {}
    tasks = [{ title: 'Setup smart bulbs', creationDate: new Date() }] as ProjectTask[];

    ngOnInit(): void {
		this.tasksService.fetchTasks();
	}

	// get tasks(): ProjectTask[]{
	// 	return this.tasksService.tasks
	// }
}
