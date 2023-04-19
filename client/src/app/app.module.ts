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


import { HeaderComponent } from './side-bar/header/header.component';
import { SidenavComponent } from './side-bar/sidenav/sidenav.component';
import { HomeComponent } from './side-bar/home/home.component';
import { ProfileComponent } from './side-bar/profile/profile.component';
import { MatSidenavModule} from '@angular/material/sidenav';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import { PostComponent } from './side-bar/post/post.component';
import { ExploreComponent } from './side-bar/explore/explore.component';
import { SideBarComponent } from './side-bar/side-bar.component';
import { SearchComponent } from './side-bar/search/search.component';
import { CommunityComponent } from './community/community.component';
import { SearchingComponent } from './side-bar/searching/searching.component';
import { ContactComponent } from './side-bar/contact/contact.component';

@NgModule({
  declarations: [
    LoginPageComponent,
    SignupComponent,
    AppComponent,
    UserProfileComponent,
    HeaderComponent,
    SidenavComponent,
    HomeComponent,
    ProfileComponent,
    PostComponent,
    ExploreComponent,
    SideBarComponent,
    SearchComponent,
    CommunityComponent,
    ContactComponent

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
    FlexLayoutModule,
    MatSidenavModule,
    MatToolbarModule,
    MatMenuModule,
    MatIconModule,
    MatDividerModule,
    MatListModule,

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
