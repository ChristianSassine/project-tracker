import { Component, OnInit } from '@angular/core';
import { Project } from 'src/app/interfaces/project';
import { ProjectService } from 'src/app/services/project.service';

@Component({
    selector: 'app-projects-page',
    templateUrl: './projects-page.component.html',
    styleUrls: ['./projects-page.component.scss'],
})
export class ProjectsPageComponent {
    constructor(private projectService: ProjectService) {}

    get projects(): Project[] {
        return this.projectService.projects;
    }
    submit() {}
}
