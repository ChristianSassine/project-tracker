import { Injectable } from '@angular/core';
import { finalize, Subject } from 'rxjs';
import { Project } from '../interfaces/project';
import { AuthService } from './auth.service';
import { HttpHandlerService, tapOnSubscribe } from './http-handler.service';

const LOCAL_STORAGE_PROJECT_PREFIX = 'current_project_';

@Injectable({
    providedIn: 'root',
})
export class ProjectService {
    public projects: Project[];

    public isLoading: boolean;
    public changeToHomePageObservable: Subject<boolean>;

    constructor(private http: HttpHandlerService, private authService: AuthService) {
        this.projects = [
            { title: 'Hello', id: 1 },
            { title: 'MisterJohn', id: 2 },
        ];
        this.isLoading = true;
        this.changeToHomePageObservable = new Subject();
    }

    setCurrentProject(project: Project) {
        const stringifiedProject = JSON.stringify(project);
        localStorage.setItem(LOCAL_STORAGE_PROJECT_PREFIX + this.authService.username.toLowerCase(), stringifiedProject);
    }

    get currentProject(): Project | null {
        if (!this.authService.username) return null;
        const stringifiedProject = localStorage.getItem(LOCAL_STORAGE_PROJECT_PREFIX + this.authService.username.toLowerCase());
        const project = stringifiedProject ? JSON.parse(stringifiedProject) : null;
        return project;
    }

    fetchProjects() {
        this.http
            .getAllProjects()
            .pipe(
                tapOnSubscribe(() => (this.isLoading = true)),
                finalize(() => (this.isLoading = false)),
            )
            .subscribe((data) => (this.projects = [...data]));
    }

    createProject(title: string) {
        this.http.createProjectRequest(title).subscribe((project) => {
            this.setCurrentProject(project);
            this.changeToHomePageObservable.next(true);
        });
    }
}
