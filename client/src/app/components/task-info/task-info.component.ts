import { Component, EventEmitter, OnDestroy, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Subscription } from 'rxjs';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-task-info',
    templateUrl: './task-info.component.html',
    styleUrls: ['./task-info.component.scss'],
})
export class TaskInfoComponent implements OnInit, OnDestroy {
    form: FormGroup;

    taskSubscription: Subscription;
    @Output() public closeInfo = new EventEmitter();
    constructor(private fb: FormBuilder, private tasksService: TasksService) {}
    ngOnInit(): void {
        this.form = this.fb.group({
            title: [this.task.title],
            description: [this.task.description],
            state: [this.task.state],
        });

        this.taskSubscription = this.tasksService.newTaskSetObservable.subscribe(() => {
            this.ngOnInit();
            this.taskSubscription.unsubscribe();
        });
    }

    ngOnDestroy(): void {
        this.taskSubscription.unsubscribe();
    }

    get task(): ProjectTask {
        return this.tasksService.currentTask;
    }

    onClose() {
        this.closeInfo.emit(true);
    }

    onUpdate() {
        const sentObject = { ...this.task, ...this.form.value };

        if (this.task.state != this.form.value.state) {
            const updatedStateIndex = 0;
            this.tasksService.updateTaskState(this.form.value.state, updatedStateIndex, this.task.id);
        }
        this.tasksService.updateTask(sentObject);
        this.form.markAsPristine();
    }
}
