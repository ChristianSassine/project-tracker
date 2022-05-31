import { Component, Input } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';

@Component({
    selector: 'app-task',
    templateUrl: './task.component.html',
    styleUrls: ['./task.component.scss'],
})
export class TaskComponent {
    @Input() task: ProjectTask;
    constructor() {}
}
