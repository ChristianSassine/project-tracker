import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeTasksPageComponent } from './home-tasks-page.component';

describe('HomeTasksPageComponent', () => {
  let component: HomeTasksPageComponent;
  let fixture: ComponentFixture<HomeTasksPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeTasksPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeTasksPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
