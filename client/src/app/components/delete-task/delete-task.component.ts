import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-delete-task',
    templateUrl: './delete-task.component.html',
    styleUrls: ['./delete-task.component.scss'],
})
export class DeleteTaskComponent {
    constructor(@Inject(MAT_DIALOG_DATA) public task : ProjectTask, private taskService: TasksService) {}

	onDelete(){
		this.taskService.deleteTask(this.task.id);
	}
}
