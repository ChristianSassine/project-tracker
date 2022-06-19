import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Subscription } from 'rxjs';
import { CreateTaskComponent } from 'src/app/components/create-task/create-task.component';
import { DeleteTaskComponent } from 'src/app/components/delete-task/delete-task.component';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';
import { TaskState } from 'src/common/task-state';

@Component({
    selector: 'app-home-tasks-page',
    templateUrl: './home-tasks-page.component.html',
    styleUrls: ['./home-tasks-page.component.scss'],
})
export class HomeTasksPageComponent implements OnInit, OnDestroy {
    // TODO : Might need to refactor the names
    stateTODO = TaskState.TODO;
    stateONGOING = TaskState.ONGOING;
    stateDONE = TaskState.DONE;

    isTODODisplayed: boolean;
    isONGOINGDisplayed: boolean;
    isDONEDisplayed: boolean;
    isINFOisplayed: boolean;

    taskChangeSubscription: Subscription;
    fetchingTasksInterval: unknown;
    constructor(private tasksService: TasksService, private dialog: MatDialog) {
        this.isTODODisplayed = true;
        this.isONGOINGDisplayed = true;
        this.isDONEDisplayed = true;
        this.isINFOisplayed = false;

        this.taskChangeSubscription = new Subscription();
    }

    ngOnInit(): void {
        this.tasksService.fetchStateTasks();
        this.taskChangeSubscription = this.tasksService.newTaskSetObservable.subscribe((task) => this.showView(task.state));
        const minuteInMilliseconds = 1000 * 60;
        this.fetchingTasksInterval = setInterval(() => this.tasksService.fetchStateTasks(), minuteInMilliseconds);
    }

    ngOnDestroy(): void {
        this.taskChangeSubscription.unsubscribe();
        clearInterval(this.fetchingTasksInterval as number);
    }

    get tasksTODO(): ProjectTask[] {
        return this.tasksService.stateTasks.get(TaskState.TODO) as ProjectTask[];
    }

    get tasksONGOING(): ProjectTask[] {
        return this.tasksService.stateTasks.get(TaskState.ONGOING) as ProjectTask[];
    }

    get tasksDONE(): ProjectTask[] {
        return this.tasksService.stateTasks.get(TaskState.DONE) as ProjectTask[];
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

    onDelete(task: ProjectTask) {
        this.dialog.open(DeleteTaskComponent, {
            data: task,
        });
    }

    onDrop(event: CdkDragDrop<ProjectTask[]>) {
        if (event.previousContainer === event.container) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            this.tasksService.updateTaskPosition(event.previousIndex, event.currentIndex, event.item.data.id);
            return;
        }
        transferArrayItem(event.previousContainer.data, event.container.data, event.previousIndex, event.currentIndex);
        this.tasksService.updateTaskState(event.container.id as TaskState, event.currentIndex, event.item.data.id);
    }

    onClose() {
        this.showView(null);
    }
}
