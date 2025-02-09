import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThalamusEmailInputComponent } from './thalamus-email-input.component';

describe('ThalamusEmailInputComponent', () => {
  let component: ThalamusEmailInputComponent;
  let fixture: ComponentFixture<ThalamusEmailInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThalamusEmailInputComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ThalamusEmailInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
