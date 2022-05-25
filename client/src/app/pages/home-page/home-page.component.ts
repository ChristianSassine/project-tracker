import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpHandlerService } from 'src/app/services/http-handler.service';
import { ProjectService } from 'src/app/services/project.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-home-page',
    templateUrl: './home-page.component.html',
    styleUrls: ['./home-page.component.scss'],
})
export class HomePageComponent implements OnInit {
    constructor(private projectService: ProjectService, private router: Router) {}

	ngOnInit(): void {
		if (!this.projectService.currentProject) this.router.navigate([Paths.Projects]);
	}

    submit() {}
}
