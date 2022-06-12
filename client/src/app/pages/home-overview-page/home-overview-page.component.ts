import { Component, OnInit } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-home-overview-page',
    templateUrl: './home-overview-page.component.html',
    styleUrls: ['./home-overview-page.component.scss'],
})
export class HomeOverviewPageComponent {
    constructor(private taskService : TasksService) {
		this.taskService.fetchRecentTasks();
	}
	
	get recentTasks(){
		return this.taskService.recentTasks;
	}
}
