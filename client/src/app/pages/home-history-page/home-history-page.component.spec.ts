import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeHistoryPageComponent } from './home-history-page.component';

describe('HomeHistoryPageComponent', () => {
  let component: HomeHistoryPageComponent;
  let fixture: ComponentFixture<HomeHistoryPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomeHistoryPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeHistoryPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
