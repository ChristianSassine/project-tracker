import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ProjectService } from 'src/app/services/project.service';

@Component({
    selector: 'app-create-project',
    templateUrl: './create-project.component.html',
    styleUrls: ['./create-project.component.scss'],
})
export class CreateProjectComponent implements OnInit {
    form: FormGroup;

	title :string = 'Project Creation'
	subtitle : string = 'Fill in the fields below'
	projectTitleLabel:string = 'Title'
    projectTitlePlaceholder: string = 'Ex: Setting up home automation';
    projectTitleError: string = 'A title is required';
	createButtonText: string = 'Create'
	

    constructor(private fb: FormBuilder, private projectService: ProjectService, private router: Router) {}

    ngOnInit(): void {
        this.form = this.fb.group({
            title: ['', Validators.required],
        });
    }

    submit() {
        if (this.form.get('title')?.valid) {
            //TODO : Handle catch
            this.projectService
                .createProject(this.form.value.title);
        }
    }
}
