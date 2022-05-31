import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Paths } from 'src/common/paths';
import { HistoryComponent } from '../components/history/history.component';
import { OverviewComponent } from '../components/overview/overview.component';
import { TasksComponent } from '../components/tasks/tasks.component';
import { HomePageComponent } from '../pages/home-page/home-page.component';
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
            { path: Paths.OverviewChild, canActivate: [AuthGuard], component: OverviewComponent },
            { path:  Paths.TasksChild, canActivate: [AuthGuard], component: TasksComponent },
            { path:  Paths.HistoryChild, canActivate: [AuthGuard], component: HistoryComponent },
            { path: '**', redirectTo: '/' + Paths.OverviewFull }
        ],
    },
    { path: Paths.Projects, canActivate: [AuthGuard], component: ProjectsPageComponent },
    { path: '**', redirectTo: '/' + Paths.Login },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
