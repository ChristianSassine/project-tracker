import { Injectable } from '@angular/core';
import { catchError, finalize, Subject } from 'rxjs';
import { ErrorShown } from 'src/common/error-shown';
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
    errorNotification: any;

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
                catchError(this.http.handleError(ErrorShown.ProjectsUnfetchable, [])),
            )
            .subscribe((data) => (this.projects = [...data]));
    }

    createProject(title: string, password: string) {
        this.http
            .createProjectRequest(title, password)
            .pipe(catchError(this.http.handleError<Project>(ErrorShown.ProjectUncreatable)))
            .subscribe((project) => {
                this.setCurrentProject(project);
                this.changeToHomePageObservable.next(true);
            });
    }

    joinProject(id: number, password: string) {
        this.http
            .joinProjectRequest(id, password)
            .pipe(catchError(this.http.handleError<Project>(ErrorShown.ProjectUnjoinable)))
            .subscribe((project) => {
                this.setCurrentProject(project);
                this.changeToHomePageObservable.next(true);
            });
    }

    deleteProject(projectId: number) {
        this.http
            .deleteProjectRequest(projectId)
            .pipe(catchError(this.http.handleError<Project>(ErrorShown.ProjectNotDeleted)))
            .subscribe(() => this.fetchProjects());
    }
}
