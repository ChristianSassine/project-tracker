import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Subject } from 'rxjs';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';
import { TaskState } from 'src/common/task-state';
import { CreateTaskComponent } from '../create-task/create-task.component';

@Component({
    selector: 'app-tasks',
    templateUrl: './tasks.component.html',
    styleUrls: ['./tasks.component.scss'],
})
export class TasksComponent implements OnInit {
    stateTODO = TaskState.TODO;
    stateONGOING = TaskState.ONGOING;
    stateDONE = TaskState.DONE;

    isTODODisplayed: boolean;
    isONGOINGDisplayed: boolean;
    isDONEDisplayed: boolean;
    isINFOisplayed: boolean;

    constructor(private tasksService: TasksService, private dialog: MatDialog) {
        this.isTODODisplayed = true;
        this.isONGOINGDisplayed = true;
        this.isDONEDisplayed = true;
        this.isINFOisplayed = false;
    }

    ngOnInit(): void {
        this.tasksService.fetchStateTasks();
    }

    get tasksTODO(): ProjectTask[] {
        return this.tasksService.tasksTODO;
    }

    get tasksONGOING(): ProjectTask[] {
        return this.tasksService.tasksONGOING;
    }

    get tasksDONE(): ProjectTask[] {
        return this.tasksService.tasksDONE;
    }

    onAdd(taskState: TaskState) {
        this.dialog.open(CreateTaskComponent, { data: taskState });
    }

    showView(state: TaskState | null) {
        switch (state) {
            case TaskState.TODO:
                this.isTODODisplayed = true;
                this.isONGOINGDisplayed = false;
                this.isDONEDisplayed = false;
                this.isINFOisplayed = true;
                break;
            case TaskState.ONGOING:
                this.isTODODisplayed = false;
                this.isONGOINGDisplayed = true;
                this.isDONEDisplayed = false;
                this.isINFOisplayed = true;
                break;
            case TaskState.DONE:
                this.isTODODisplayed = false;
                this.isONGOINGDisplayed = false;
                this.isDONEDisplayed = true;
                this.isINFOisplayed = true;
                break;
            default:
                this.isTODODisplayed = true;
                this.isONGOINGDisplayed = true;
                this.isDONEDisplayed = true;
                this.isINFOisplayed = false;
                break;
        }
    }

    onTask(task: ProjectTask) {
        this.tasksService.setCurrentTask(task);
        this.showView(task.state);
    }

    onClose(){
        this.showView(null);
    }
}
