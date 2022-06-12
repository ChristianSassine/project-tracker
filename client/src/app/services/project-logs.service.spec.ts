import { TestBed } from '@angular/core/testing';

import { ProjectLogsService } from './project-logs.service';

describe('ProjectLogsService', () => {
  let service: ProjectLogsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ProjectLogsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
