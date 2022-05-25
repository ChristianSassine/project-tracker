import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Paths } from 'src/common/paths';
import { HomePageComponent } from '../pages/home-page/home-page.component';
import { LoginPageComponent } from '../pages/login-page/login-page.component';
import { ProjectsPageComponent } from '../pages/projects-page/projects-page.component';
import { AuthGuard } from '../services/guards/auth.guard';

const routes: Routes = [
    {path: Paths.Login, component: LoginPageComponent},
    {path:'', canActivate: [AuthGuard], children :[
        {path: '', redirectTo: '/' + Paths.Login, pathMatch: 'full'},
        {path: Paths.Home, component: HomePageComponent},
        {path: Paths.Projects, component: ProjectsPageComponent},
    ]},
    {path:'**', redirectTo:'/' + Paths.Login}
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
