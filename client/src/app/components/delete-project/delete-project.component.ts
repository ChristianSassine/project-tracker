import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Project } from 'src/app/interfaces/project';
import { ProjectService } from 'src/app/services/project.service';

@Component({
    selector: 'app-delete-project',
    templateUrl: './delete-project.component.html',
    styleUrls: ['./delete-project.component.scss'],
})
export class DeleteProjectComponent {
    deleteConfirmed = true;
    deleteCanceled = false;
    constructor(@Inject(MAT_DIALOG_DATA) public project : Project, private projectService : ProjectService) {}

    onDelete(){
        this.projectService.deleteProject(this.project.id);
    }

}
