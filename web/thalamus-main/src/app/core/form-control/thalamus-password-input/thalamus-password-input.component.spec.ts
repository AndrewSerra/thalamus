import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThalamusPasswordInputComponent } from './thalamus-password-input.component';

describe('ThalamusPasswordInputComponent', () => {
  let component: ThalamusPasswordInputComponent;
  let fixture: ComponentFixture<ThalamusPasswordInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThalamusPasswordInputComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ThalamusPasswordInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
