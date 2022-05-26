import { Injectable } from '@angular/core';
import { finalize } from 'rxjs';
import { Project } from '../interfaces/project';
import { HttpHandlerService, tapOnSubscribe } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class ProjectService {
	public projects : Project[];
	public currentProject : string;
	public isLoading : boolean;

    constructor(private http: HttpHandlerService) {
		this.projects = [{title: 'Hello' , id : 1}, {title: 'MisterJohn', id : 2}]
		this.currentProject = "";
		this.isLoading = true;
	}

	fetchProjects(){
        this.http.getAllProjects().pipe(tapOnSubscribe(()=> this.isLoading = true), finalize(()=> this.isLoading = false)).subscribe((data)=> this.projects = [...data]);
	}
}
