import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ThalamusTextInputComponent } from './thalamus-text-input.component';

describe('ThalamusTextInputComponent', () => {
  let component: ThalamusTextInputComponent;
  let fixture: ComponentFixture<ThalamusTextInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ThalamusTextInputComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ThalamusTextInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
