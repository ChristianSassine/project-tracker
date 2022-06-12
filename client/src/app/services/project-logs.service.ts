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
    recentProjectLogs: HistoryLog[];
    logsUpdatedObservable: Subject<boolean>;
    recentLogsUpdatedObservable: Subject<boolean>;
    constructor(private projectService: ProjectService, private http: HttpHandlerService) {
        this.logsUpdatedObservable = new Subject();
        this.recentLogsUpdatedObservable = new Subject();
        this.projectLogs = [];
    }

    fetchProjectLogs() {
        if (!this.projectService.currentProject) return;
        this.http.getAllProjectLogs(this.projectService.currentProject.id).subscribe((logs) => {
            this.projectLogs = [...logs];
            this.logsUpdatedObservable.next(true);
        });
    }

    fetchRecentProjectLogs() {
        if (!this.projectService.currentProject) return;
        this.http.getRecentProjectLogs(this.projectService.currentProject.id).subscribe((logs) => {
            this.recentProjectLogs = [...logs];
            this.recentLogsUpdatedObservable.next(true);
        });
    }
}
