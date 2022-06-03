import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './modules/app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './modules/material.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ServiceWorkerModule } from '@angular/service-worker';
import { environment } from '../environments/environment';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { ProjectsPageComponent } from './pages/projects-page/projects-page.component';
import { CreateProjectComponent } from './components/create-project/create-project.component';
import { OverviewComponent } from './components/overview/overview.component';
import { TasksComponent } from './components/tasks/tasks.component';
import { HistoryComponent } from './components/history/history.component';
import { TaskComponent } from './components/task/task.component';
import { CreateTaskComponent } from './components/create-task/create-task.component';
import { TaskInfoComponent } from './components/task-info/task-info.component';

@NgModule({
    declarations: [AppComponent, HomePageComponent, LoginComponent, RegisterComponent, LoginPageComponent, ProjectsPageComponent, CreateProjectComponent, OverviewComponent, TasksComponent, HistoryComponent, TaskComponent, CreateTaskComponent, TaskInfoComponent],
    imports: [
        HttpClientModule,
        ReactiveFormsModule,
        MaterialModule,
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        ServiceWorkerModule.register('ngsw-worker.js', {
            enabled: environment.production,
            // Register the ServiceWorker as soon as the application is stable
            // or after 30 seconds (whichever comes first).
            registrationStrategy: 'registerWhenStable:30000',
        }),
    ],
    providers: [],
    bootstrap: [AppComponent],
})
export class AppModule {}
