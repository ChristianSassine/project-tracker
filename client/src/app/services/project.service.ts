import { Injectable } from '@angular/core';
import { finalize, Subject } from 'rxjs';
import { Project } from '../interfaces/project';
import { HttpHandlerService, tapOnSubscribe } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class ProjectService {
	public projects : Project[];
	public currentProject : Project;

	public isProjectSelected : boolean;
	public isLoading : boolean;

	public changeToHomePageObservable : Subject<boolean>;

    constructor(private http: HttpHandlerService) {
		this.projects = [{title: 'Hello' , id : 1}, {title: 'MisterJohn', id : 2}]
		this.currentProject = {} as Project;
		this.isProjectSelected = false;
		this.isLoading = true;
		this.changeToHomePageObservable = new Subject();
	}

	setCurrentProject(project: Project){
		this.isProjectSelected = true;
		this.currentProject = project;
	}

	fetchProjects(){
        this.http.getAllProjects().pipe(tapOnSubscribe(()=> this.isLoading = true), finalize(()=> this.isLoading = false)).subscribe((data)=> this.projects = [...data]);
	}

	createProject(title : string){
		this.http.createProjectRequest(title).subscribe((project)=>{
			this.setCurrentProject(project);
			this.changeToHomePageObservable.next(true);
		});
	}
}
