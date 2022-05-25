import { Injectable } from '@angular/core';
import { Project } from '../interfaces/project';

@Injectable({
    providedIn: 'root',
})
export class ProjectService {
	public projects : Project[];
	public currentProject : string;
    constructor() {
		this.projects = [{title: 'Hello' , id : 1}, {title: 'MisterJohn', id : 2}]
		this.currentProject = "";
	}
}
