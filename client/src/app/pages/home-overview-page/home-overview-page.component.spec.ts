import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeOverviewPageComponent } from './home-overview-page.component';

describe('HomeOverviewPageComponent', () => {
  let component: HomeOverviewPageComponent;
  let fixture: ComponentFixture<HomeOverviewPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeOverviewPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeOverviewPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
