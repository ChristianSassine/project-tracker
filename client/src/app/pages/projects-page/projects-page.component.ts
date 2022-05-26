import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { CreateProjectComponent } from 'src/app/components/create-project/create-project.component';
import { Project } from 'src/app/interfaces/project';
import { HttpHandlerService } from 'src/app/services/http-handler.service';
import { ProjectService } from 'src/app/services/project.service';

@Component({
    selector: 'app-projects-page',
    templateUrl: './projects-page.component.html',
    styleUrls: ['./projects-page.component.scss'],
})
export class ProjectsPageComponent implements OnInit {
    private readonly minimumDialogWidth = '17vw'

    constructor(private projectService: ProjectService, private dialog: MatDialog) {}

    get projects(): Project[] {
        return this.projectService.projects;
    }

    get isLoading(): boolean {
        return this.projectService.isLoading;
    }

    ngOnInit(): void {
        this.projectService.fetchProjects();
    }

    openDialog() {
        this.dialog.open(CreateProjectComponent, {minWidth: this.minimumDialogWidth});
    }

}
