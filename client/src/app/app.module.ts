import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RecaptchaModule } from 'ng-recaptcha';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RecaptchaModule
  ],
  providers: [

  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
