import { AfterViewInit, Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { DeleteTaskComponent } from 'src/app/components/delete-task/delete-task.component';
import { HistoryLog } from 'src/app/interfaces/history-log';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { ProjectLogsService } from 'src/app/services/project-logs.service';
import { TasksService } from 'src/app/services/tasks.service';
import Chart, { ChartConfiguration } from 'chart.js/auto';

@Component({
    selector: 'app-home-overview-page',
    templateUrl: './home-overview-page.component.html',
    styleUrls: ['./home-overview-page.component.scss'],
})
export class HomeOverviewPageComponent implements OnInit, OnDestroy, AfterViewInit {
    @ViewChild('statsCanvas', { static: false }) private statsCanvas!: ElementRef<HTMLCanvasElement>;
    displayedColumns: string[] = ['date', 'logger', 'log'];
    dataSource: MatTableDataSource<HistoryLog>;

    private logsUpdateSubscription: Subscription;
    private statsChart: Chart;

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

    ngAfterViewInit(): void {
		this.generateStatsChart();
    }

    ngOnDestroy(): void {
        this.logsUpdateSubscription.unsubscribe();
        this.statsChart.destroy();
    }

    get recentTasks() {
        return this.taskService.recentTasks;
    }

    onDelete(task: ProjectTask) {
        const dialogRef = this.dialog.open(DeleteTaskComponent, {
            data: task,
        });
        const dialogCloseSubscription = dialogRef.afterClosed().subscribe((isDeleted) => {
            if(!isDeleted) return;
            this.taskService.fetchRecentTasks();
            this.generateStatsChart();
            this.logService.fetchRecentProjectLogs();
            dialogCloseSubscription.unsubscribe();
        });
    }

    private generateStatsChart(){
        this.taskService.getTasksStats().subscribe((stats)=> {
            if (this.statsChart) this.statsChart.destroy();
			const data = {
				labels: ['Todo', 'Ongoing', 'Done'],
				datasets: [
					{
						label: 'Project statistics',
						data: [stats.todoTasks, stats.ongoingTasks, stats.doneTasks],
						backgroundColor: ['#c4c58c', '#8cc2c5', '#8cc58c'],
						hoverOffset: 4,
					},
				],
			};

			const config = {
				type: 'doughnut',
				data: data,
				options: { maintainAspectRatio: false },
			} as ChartConfiguration;
			this.statsChart = new Chart(this.statsCanvas.nativeElement.getContext('2d') as CanvasRenderingContext2D, config);
		})
    }
}
