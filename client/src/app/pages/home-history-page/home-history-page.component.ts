import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { HistoryLog } from 'src/app/interfaces/history-log';

@Component({
    selector: 'app-home-history-page',
    templateUrl: './home-history-page.component.html',
    styleUrls: ['./home-history-page.component.scss'],
})
export class HomeHistoryPageComponent implements AfterViewInit {
    @ViewChild(MatPaginator) paginator: MatPaginator;
    @ViewChild(MatSort) sort: MatSort;

    displayedColumns: string[] = ['creationDate', 'creator', 'log'];
    dataSource = new MatTableDataSource<HistoryLog>();
    constructor() {}
    ngAfterViewInit(): void {
        this.dataSource.paginator = this.paginator;
        this.dataSource.sort = this.sort;

        this.sort.sortChange.subscribe(() => (this.paginator.pageIndex = 0));
    }

    applyFilter(event: Event) {
        const filterValue = (event.target as HTMLInputElement).value;
        this.dataSource.filter = filterValue.trim().toLowerCase();
    }

	onRefresh(){
		// TODO: implement refresh
	}
}