import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Paths } from 'src/common/paths';
import { HomeHistoryPageComponent } from '../pages/home-history-page/home-history-page.component';
import { HomeOverviewPageComponent } from '../pages/home-overview-page/home-overview-page.component';
import { HomePageComponent } from '../pages/home-page/home-page.component';
import { HomeTasksPageComponent } from '../pages/home-tasks-page/home-tasks-page.component';
import { LoginPageComponent } from '../pages/login-page/login-page.component';
import { ProjectsPageComponent } from '../pages/projects-page/projects-page.component';
import { AuthGuard } from '../services/guards/auth.guard';
import { LoginGuard } from '../services/guards/login.guard';

const routes: Routes = [
    { path: Paths.Login, canActivate: [LoginGuard], component: LoginPageComponent },
    {
        path: Paths.Home,
        canActivate: [AuthGuard],
        component: HomePageComponent,
        children: [
            { path: Paths.OverviewChild, canActivate: [AuthGuard], component: HomeOverviewPageComponent },
            { path:  Paths.TasksChild, canActivate: [AuthGuard], component: HomeTasksPageComponent },
            { path:  Paths.HistoryChild, canActivate: [AuthGuard], component: HomeHistoryPageComponent },
            { path: '**', redirectTo: '/' + Paths.OverviewFull }
        ],
    },
    { path: Paths.Projects, canActivate: [AuthGuard], component: ProjectsPageComponent },
    { path: '**', redirectTo: '/' + Paths.Login },
];

@NgModule({
    imports: [RouterModule.forRoot(routes, { "useHash" : true })],
    exports: [RouterModule],
})
export class AppRoutingModule {}
