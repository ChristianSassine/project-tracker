import { Component, EventEmitter, Output } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-task-info',
    templateUrl: './task-info.component.html',
    styleUrls: ['./task-info.component.scss'],
})
export class TaskInfoComponent {
    form : FormGroup;
    @Output() public closeInfo = new EventEmitter();
    constructor(private fb: FormBuilder, private tasksService: TasksService) {
        this.form = this.fb.group({
            username: [''],
            password: [''],
        });
    }

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
