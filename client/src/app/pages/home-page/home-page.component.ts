import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { Project } from 'src/app/interfaces/project';
import { AuthService } from 'src/app/services/auth.service';
import { ProjectService } from 'src/app/services/project.service';
import { TasksService } from 'src/app/services/tasks.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-home-page',
    templateUrl: './home-page.component.html',
    styleUrls: ['./home-page.component.scss'],
})
export class HomePageComponent implements OnInit, OnDestroy {
    username: string;
    title : string;
    logoutSubscription: Subscription;

    constructor(
        private projectService: ProjectService,
        private authService: AuthService,
        private router: Router,
    ) {
        this.title = '';
        this.logoutSubscription = this.authService.logoutObservable.subscribe(() => this.router.navigate([Paths.Login]));
    }

    ngOnInit(): void {
        if (!this.projectService.currentProject) {
            this.router.navigate([Paths.Projects]);
            return;
        }
        this.username = this.capitalizeFirstLetter(this.authService.username);
    }

    ngOnDestroy(): void {
        this.logoutSubscription.unsubscribe();
    }

    get project() {
        return this.projectService.currentProject as Project;
    }

    isOverviewSelected(): boolean {
        return this.router.url === '/' + Paths.OverviewFull;
    }

    isTasksSelected(): boolean {
        return this.router.url === '/' + Paths.TasksFull;
    }

    isHistorySelected(): boolean {
        return this.router.url === '/' + Paths.HistoryFull;
    }

    onOverview() {
        this.router.navigate([Paths.OverviewFull]);
    }

    onTasks() {
        this.router.navigate([Paths.TasksFull]);
    }

    onHistory() {
        this.router.navigate([Paths.HistoryFull]);
    }

    private capitalizeFirstLetter(word: string): string {
        return word.charAt(0).toUpperCase() + word.slice(1);
    }

    onLogout() {
        console.log('Attempting logout');
        this.authService.logout();
        this.router.navigate([Paths.Login]);
    }

    onChangeProject() {
        this.router.navigate([Paths.Projects]);
    }
}
