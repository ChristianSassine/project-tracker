import { Component, ElementRef, EventEmitter, HostListener, Input, OnInit, Output } from '@angular/core';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';
import { TaskState } from 'src/common/task-state';

@Component({
    selector: 'app-task',
    templateUrl: './task.component.html',
    styleUrls: ['./task.component.scss'],
})
export class TaskComponent {
    @Input() task: ProjectTask;
    @Input() isDraggable: boolean = false;
    @Input() isSelected: boolean = false;
    @Output() delete: EventEmitter<unknown>;
    todoState = TaskState.TODO;
    ongoingState = TaskState.ONGOING;
    doneState = TaskState.DONE;

    constructor(private elementRef: ElementRef, private tasksService: TasksService) {
        this.delete = new EventEmitter();
    }

    // TODO: Might need to find a better way
    @HostListener('document:click')
    onClick() {
        if (this.tasksService.currentTask.id === this.task.id) {
            this.isSelected = true;
            return;
        }
        this.isSelected = false;
    }

    onDelete($event: Event) {
        this.delete.emit();
        $event.stopPropagation();
    }
}
