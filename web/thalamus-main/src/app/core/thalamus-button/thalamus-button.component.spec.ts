import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThalamusButtonComponent } from './thalamus-button.component';

describe('ThalamusButtonComponent', () => {
  let component: ThalamusButtonComponent;
  let fixture: ComponentFixture<ThalamusButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThalamusButtonComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ThalamusButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
