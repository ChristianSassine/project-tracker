import { AfterViewInit, Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { HistoryLog } from 'src/app/interfaces/history-log';
import { ProjectLogsService } from 'src/app/services/project-logs.service';
import { MatTable } from '@angular/material/table';

@Component({
    selector: 'app-home-history-page',
    templateUrl: './home-history-page.component.html',
    styleUrls: ['./home-history-page.component.scss'],
})
export class HomeHistoryPageComponent implements AfterViewInit, OnDestroy {
    @ViewChild(MatPaginator) paginator: MatPaginator;
    @ViewChild(MatSort) sort: MatSort;

    displayedColumns: string[] = ['creationDate', 'creator', 'log'];
    dataSource = new MatTableDataSource<HistoryLog>();

    private logUpdateSubscription: Subscription;

    constructor(private logsService: ProjectLogsService) {}

    ngAfterViewInit(): void {
        this.dataSource.sort = this.sort;
        this.logUpdateSubscription = this.logsService.projectLogsUpdated.subscribe(() => {
            this.dataSource.data = this.logsService.projectLogs;
            this.dataSource._updateChangeSubscription();
        });

        this.dataSource.paginator = this.paginator;
        // TODO: Fix sort, it broke '-'

        this.sort.sortChange.subscribe(() => (this.paginator.pageIndex = 0));
        this.dataSource.data = this.logsService.projectLogs;
        this.logsService.getProjectLogs();
    }

    ngOnDestroy(): void {
        this.logUpdateSubscription.unsubscribe();
    }

    applyFilter(event: Event) {
        const filterValue = (event.target as HTMLInputElement).value;
        this.dataSource.filter = filterValue.trim().toLowerCase();
    }

    onRefresh() {
        // TODO : Add refreshing from server
        this.logsService.getProjectLogs();
    }
}
