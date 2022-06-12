import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TaskState } from 'src/common/task-state';

@Component({
    selector: 'app-task',
    templateUrl: './task.component.html',
    styleUrls: ['./task.component.scss'],
})
export class TaskComponent {
    @Input() task: ProjectTask;
    @Output() delete: EventEmitter<unknown>;
    todoState = TaskState.TODO;
    ongoingState = TaskState.ONGOING;
    doneState = TaskState.DONE;

    constructor() {
        this.delete = new EventEmitter();
    }

    onDelete($event: Event){
        this.delete.emit();
        $event.stopPropagation();
    }
}
