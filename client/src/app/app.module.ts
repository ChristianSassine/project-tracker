import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './modules/app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './modules/material.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { ProjectsPageComponent } from './pages/projects-page/projects-page.component';
import { CreateProjectComponent } from './components/create-project/create-project.component';
import { TaskComponent } from './components/task/task.component';
import { CreateTaskComponent } from './components/create-task/create-task.component';
import { TaskInfoComponent } from './components/task-info/task-info.component';
import { DeleteTaskComponent } from './components/delete-task/delete-task.component';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { HomeTasksPageComponent } from './pages/home-tasks-page/home-tasks-page.component';
import { HomeOverviewPageComponent } from './pages/home-overview-page/home-overview-page.component';
import { HomeHistoryPageComponent } from './pages/home-history-page/home-history-page.component';
import {TextFieldModule} from '@angular/cdk/text-field';
import { LogParserPipe } from './pipes/log-parser.pipe';
import { DeleteProjectComponent } from './components/delete-project/delete-project.component';
import { JoinProjectComponent } from './components/join-project/join-project.component';

@NgModule({
    declarations: [
        AppComponent,
        HomePageComponent,
        LoginComponent,
        RegisterComponent,
        LoginPageComponent,
        ProjectsPageComponent,
        CreateProjectComponent,
        TaskComponent,
        CreateTaskComponent,
        TaskInfoComponent,
        DeleteTaskComponent,
        HomeTasksPageComponent,
        HomeOverviewPageComponent,
        HomeHistoryPageComponent,
        LogParserPipe,
        DeleteProjectComponent,
        JoinProjectComponent,
    ],
    imports: [
        HttpClientModule,
        ReactiveFormsModule,
        MaterialModule,
        BrowserModule,
        AppRoutingModule,
        TextFieldModule,
        BrowserAnimationsModule,
        DragDropModule
    ],
    providers: [],
    bootstrap: [AppComponent],
})
export class AppModule {}
