import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';

@Component({
    selector: 'app-task',
    templateUrl: './task.component.html',
    styleUrls: ['./task.component.scss'],
})
export class TaskComponent {
    @Input() task: ProjectTask;
    @Output() delete: EventEmitter<unknown>;
    constructor() {
        this.delete = new EventEmitter();
    }

    onDelete($event: Event){
        this.delete.emit();
        $event.stopPropagation();
    }
}
