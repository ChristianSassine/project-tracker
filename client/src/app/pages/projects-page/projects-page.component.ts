import { Component, OnInit } from '@angular/core';
import { Project } from 'src/app/interfaces/project';
import { HttpHandlerService } from 'src/app/services/http-handler.service';
import { ProjectService } from 'src/app/services/project.service';

@Component({
    selector: 'app-projects-page',
    templateUrl: './projects-page.component.html',
    styleUrls: ['./projects-page.component.scss'],
})
export class ProjectsPageComponent implements OnInit {
    constructor(private projectService: ProjectService) {}

    get projects(): Project[] {
        return this.projectService.projects;
    }

    get isLoading(): boolean {
        return this.projectService.isLoading;
    }

    ngOnInit(): void {
        this.projectService.fetchProjects();
    }

}
