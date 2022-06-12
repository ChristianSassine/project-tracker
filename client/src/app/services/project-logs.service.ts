import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { HistoryLog } from '../interfaces/history-log';
import { HttpHandlerService } from './http-handler.service';
import { ProjectService } from './project.service';

@Injectable({
    providedIn: 'root',
})
export class ProjectLogsService {
    projectLogs: HistoryLog[];
    projectLogsUpdated: Subject<boolean>;
    constructor(private projectService: ProjectService, private http: HttpHandlerService) {
        this.projectLogsUpdated = new Subject();
        this.projectLogs = [];
    }

    getProjectLogs() {
        if (!this.projectService.currentProject) return;
        this.http.getAllProjectLogs(this.projectService.currentProject.id).subscribe((logs) => {
            this.projectLogs = [...logs];
            this.projectLogsUpdated.next(true);
        });
    }
}
