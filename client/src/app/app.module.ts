import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms'
import { HttpClientModule } from '@angular/common/http'


import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatToolbarModule} from '@angular/material/toolbar';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './login-page/login-page.component';
import { SignupComponent } from './signup-page/signup-page.component';
import { RouterModule } from '@angular/router'
import { AppRoutingModule } from './app-routing.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { UserProfileComponent } from './user-profile/user-profile.component';

@NgModule({
  declarations: [
    LoginPageComponent,
    SignupComponent,
    AppComponent,
    UserProfileComponent

  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatToolbarModule,
    RouterModule,
    AppRoutingModule,
    FlexLayoutModule

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
