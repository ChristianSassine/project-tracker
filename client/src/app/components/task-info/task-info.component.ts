import { Component, EventEmitter, Output } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';
import { TaskState } from 'src/common/task-state';

@Component({
    selector: 'app-task-info',
    templateUrl: './task-info.component.html',
    styleUrls: ['./task-info.component.scss'],
})
export class TaskInfoComponent {
    @Output() public closeInfo = new EventEmitter();
    constructor(private tasksService: TasksService) {}

    get task(): ProjectTask {
        return this.tasksService.currentTask;
    }

    onClose() {
        this.closeInfo.emit(true);
    }

    onUpdate() {
        this.tasksService.updateTask(this.task);
    }
}
