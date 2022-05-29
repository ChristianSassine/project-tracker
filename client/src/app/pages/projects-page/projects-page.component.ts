import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { CreateProjectComponent } from 'src/app/components/create-project/create-project.component';
import { Project } from 'src/app/interfaces/project';
import { ProjectService } from 'src/app/services/project.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-projects-page',
    templateUrl: './projects-page.component.html',
    styleUrls: ['./projects-page.component.scss'],
})
export class ProjectsPageComponent implements OnInit {
    private readonly minimumDialogWidth = '17vw';

    constructor(private projectService: ProjectService, private router: Router, private dialog: MatDialog) {}

    get projects(): Project[] {
        return this.projectService.projects;
    }

    get isLoading(): boolean {
        return this.projectService.isLoading;
    }

    ngOnInit(): void {
        this.projectService.fetchProjects();
        this.projectService.changeToHomePageObservable.subscribe((allowed) => {
            if (allowed) this.router.navigate([Paths.Home]);
        });
    }

    onProjectSelection(project: Project) {
        this.projectService.setCurrentProject(project);
        this.router.navigate([Paths.Home]);
    }

    openDialog() {
        this.dialog.open(CreateProjectComponent, { minWidth: this.minimumDialogWidth });
    }
}
