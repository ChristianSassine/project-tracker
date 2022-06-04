import { Component, EventEmitter, Output } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-task-info',
    templateUrl: './task-info.component.html',
    styleUrls: ['./task-info.component.scss'],
})
export class TaskInfoComponent {
    @Output() public closeInfo = new EventEmitter();
    constructor(private tasksService: TasksService) {}
    task = {
        title: 'Make a HelloWorld',
        description: 'Start by plugging in the keyboard and recreating the universe',
        creationDate: new Date(),
        state: 'TODO',
    } as ProjectTask;

    onClose() {
        this.closeInfo.next(true);
    }

    onUpdate() {
        this.tasksService.updateTask(this.task);
    }
}
