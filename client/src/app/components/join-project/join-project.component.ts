import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProjectService } from 'src/app/services/project.service';

@Component({
    selector: 'app-join-project',
    templateUrl: './join-project.component.html',
    styleUrls: ['./join-project.component.scss'],
})
export class JoinProjectComponent implements OnInit {
    form: FormGroup;
    constructor(private fb: FormBuilder, private projectService : ProjectService) {}

    ngOnInit(): void {
        this.form = this.fb.group({
            id: ['', [Validators.pattern("^[0-9]*$"), Validators.required] ],
            password: ['', Validators.required]
        });
    }

    submit(){
        this.projectService.joinProject(parseInt(this.form.value.id), this.form.value.password)
    }
}
