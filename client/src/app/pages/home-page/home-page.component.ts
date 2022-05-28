import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';
import { ProjectService } from 'src/app/services/project.service';
import { TasksService } from 'src/app/services/tasks.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-home-page',
    templateUrl: './home-page.component.html',
    styleUrls: ['./home-page.component.scss'],
})
export class HomePageComponent implements OnInit {
    username : string;

    constructor(private projectService: ProjectService, private tasksService: TasksService, private authService: AuthService, private router: Router) {}

	ngOnInit(): void {
		if (!this.projectService.isProjectSelected) {
            this.router.navigate([Paths.Projects]);
            return;
        }
        this.tasksService.getTasks();
        this.username = this.capitalizeFirstLetter(this.authService.username);
	}

    get project(){
        return this.projectService.currentProject
    }

    get tasks(){
        return this.tasksService.tasks;
    }

    private capitalizeFirstLetter(word: string) : string {
        return word.charAt(0).toUpperCase() + word.slice(1);
      }
}
