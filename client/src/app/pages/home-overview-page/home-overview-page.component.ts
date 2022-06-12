import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { DeleteTaskComponent } from 'src/app/components/delete-task/delete-task.component';
import { HistoryLog } from 'src/app/interfaces/history-log';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { ProjectLogsService } from 'src/app/services/project-logs.service';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-home-overview-page',
    templateUrl: './home-overview-page.component.html',
    styleUrls: ['./home-overview-page.component.scss'],
})
export class HomeOverviewPageComponent implements OnInit, OnDestroy {
    displayedColumns: string[] = ['date', 'logger', 'log'];
    dataSource: MatTableDataSource<HistoryLog>;
    logsUpdateSubscription: Subscription;

    constructor(private taskService: TasksService, private logService: ProjectLogsService, private dialog: MatDialog) {
        this.dataSource = new MatTableDataSource([] as HistoryLog[]);
    }

    ngOnInit(): void {
        this.logsUpdateSubscription = this.logService.recentLogsUpdatedObservable.subscribe(
            () => (this.dataSource.data = this.logService.recentProjectLogs),
        );

        this.taskService.fetchRecentTasks();
        this.logService.fetchRecentProjectLogs();
    }

    ngOnDestroy(): void {
        this.logsUpdateSubscription.unsubscribe();
    }

    get recentTasks() {
        return this.taskService.recentTasks;
    }

    onDelete(task: ProjectTask) {
        const dialogRef = this.dialog.open(DeleteTaskComponent, {
            data: task,
        });
        const dialogCloseSubscription = dialogRef.afterClosed().subscribe(() => {
            this.taskService.fetchRecentTasks();
            dialogCloseSubscription.unsubscribe();
        });
    }
}
