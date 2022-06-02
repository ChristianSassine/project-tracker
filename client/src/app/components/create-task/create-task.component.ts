import { Component, Inject, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TasksService } from 'src/app/services/tasks.service';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TaskState } from 'src/common/task-state';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
    selector: 'app-create-task',
    templateUrl: './create-task.component.html',
    styleUrls: ['./create-task.component.scss'],
})
export class CreateTaskComponent implements OnInit {
    form: FormGroup;

    constructor(@Inject(MAT_DIALOG_DATA) public data: TaskState, private fb: FormBuilder, private tasksService: TasksService) {}

    ngOnInit(): void {
        this.form = this.fb.group({
            title: ['', Validators.required],
            description: [''],
            state: [this.data, Validators.required],
        });
    }

    submit() {
        const newTask = {
            title: this.form.value.title,
            description: this.form.value.description,
            state: this.form.value.state,
            creationDate: new Date(),
        } as ProjectTask;
        this.tasksService.uploadTask(newTask);
    }
}
