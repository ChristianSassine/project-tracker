import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProjectService } from 'src/app/services/project.service';
import { TasksService } from 'src/app/services/tasks.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-home-page',
    templateUrl: './home-page.component.html',
    styleUrls: ['./home-page.component.scss'],
})
export class HomePageComponent implements OnInit {
    constructor(private projectService: ProjectService, private tasksService: TasksService, private router: Router) {}

	ngOnInit(): void {
		if (!this.projectService.isProjectSelected) {
            this.router.navigate([Paths.Projects]);
            return;
        }
        this.tasksService.getTasks();
	}

    get tasks(){
        return this.tasksService.tasks;
    }
}
