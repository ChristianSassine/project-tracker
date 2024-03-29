import { Component, EventEmitter, OnDestroy, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Subscription } from 'rxjs';
import { ProjectTask } from 'src/app/interfaces/project-task';
import { AuthService } from 'src/app/services/auth.service';
import { TasksService } from 'src/app/services/tasks.service';

@Component({
    selector: 'app-task-info',
    templateUrl: './task-info.component.html',
    styleUrls: ['./task-info.component.scss'],
})
export class TaskInfoComponent implements OnInit, OnDestroy {
    form: FormGroup;
    commentForm: FormGroup;
    addingComment: boolean;

    taskSubscription: Subscription;
    @Output() public closeInfo = new EventEmitter();
    constructor(private fb: FormBuilder, private tasksService: TasksService, private authService: AuthService) {}
    ngOnInit(): void {
        this.form = this.fb.group({
            title: [this.task.title],
            description: [this.task.description],
            state: [this.task.state],
        });
        this.commentForm = this.fb.group({ content: ['', Validators.required] });

        this.addingComment = false;
        this.tasksService.fetchComments();

        this.taskSubscription = this.tasksService.newTaskSetObservable.subscribe(() => {
            this.ngOnInit();
            this.taskSubscription.unsubscribe();
        });
    }

    ngOnDestroy(): void {
        this.taskSubscription.unsubscribe();
    }

    get comments(){
        return this.tasksService.currentComments;
    }

    get task(): ProjectTask {
        return this.tasksService.currentTask;
    }

    isCurrentUser(commenter: string) {
        return commenter === this.authService.username;
    }

    onClose() {
        this.closeInfo.emit(true);
    }

    onAddComment() {
        this.addingComment = true;
    }

    onSendComment() {
        this.tasksService.addComment(this.commentForm.value.content);
        this.commentForm.setValue({ content: '' });
        this.commentForm.markAsPristine();
        this.addingComment = false;
    }

    onCancelComment() {
        this.addingComment = false;
    }

    onUpdate() {
        const sentObject = { ...this.task, ...this.form.value };
        this.tasksService.updateTask(sentObject);
        this.form.markAsPristine();
    }
}
