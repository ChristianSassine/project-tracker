import { Component, OnInit } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-tasks',
    templateUrl: './tasks.component.html',
    styleUrls: ['./tasks.component.scss'],
})
export class TasksComponent implements OnInit {
    isTODODisplayed: boolean;
    isONGOINGDisplayed: boolean;
    isDONEDisplayed: boolean;
    isINFOisplayed: boolean;
    constructor(private tasksService: TasksService) {
        this.isTODODisplayed = true;
        this.isONGOINGDisplayed = true;
        this.isDONEDisplayed = true;
        this.isINFOisplayed = false;
    }
    tasks = [{ title: 'Setup smart bulbs', creationDate: new Date() }] as ProjectTask[];

    ngOnInit(): void {
        this.tasksService.fetchTasks();
    }

    onTodoTask() {
        this.isTODODisplayed = true;
        this.isONGOINGDisplayed = false;
        this.isDONEDisplayed = false;
        this.isINFOisplayed = true;
    }

    // get tasks(): ProjectTask[]{
    // 	return this.tasksService.tasks
    // }
}
