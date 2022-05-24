import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class ProjectService {
	public projects : string[];
    constructor() {
		this.projects = ['Hello', 'MisterJohn']
	}
}
