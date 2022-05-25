import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Project } from 'src/app/interfaces/project';
import { ProjectService } from 'src/app/services/project.service';
import { environment } from 'src/environments/environment';

@Component({
    selector: 'app-projects-page',
    templateUrl: './projects-page.component.html',
    styleUrls: ['./projects-page.component.scss'],
})
export class ProjectsPageComponent {
    constructor(private projectService: ProjectService, private http: HttpClient) {}

    get projects(): Project[] {
        return this.projectService.projects;
    }
    submit() {
        this.http.get(`${environment.serverUrl}/data/projects`, {withCredentials: true}).subscribe();
    }
}
