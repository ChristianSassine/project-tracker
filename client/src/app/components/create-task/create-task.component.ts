import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TasksService } from 'src/app/services/tasks.service';
import { ProjectTask } from 'src/app/interfaces/project-task';

@Component({
    selector: 'app-create-task',
    templateUrl: './create-task.component.html',
    styleUrls: ['./create-task.component.scss'],
})
export class CreateTaskComponent implements OnInit {
    form: FormGroup;

    constructor(private fb: FormBuilder, private tasksService: TasksService) {}

    ngOnInit(): void {
        this.form = this.fb.group({
            title: ['', Validators.required],
            description: [''],
            state: ['TODO', Validators.required],
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
