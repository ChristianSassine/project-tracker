import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from '../pages/home-page/home-page.component';
import { LoginPageComponent } from '../pages/login-page/login-page.component';
import { AuthGuard } from '../services/auth.guard';

const routes: Routes = [
    {path: 'login', component: LoginPageComponent},
    {path:'', canActivate: [AuthGuard], children :[
        {path: 'home', component: HomePageComponent},
    ]},
    {path:'**', redirectTo:'/login'}
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
