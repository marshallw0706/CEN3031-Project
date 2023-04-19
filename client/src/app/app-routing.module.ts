import { UserProfileComponent } from './user-profile/user-profile.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPageComponent } from './login-page/login-page.component';
import { SignupComponent } from './signup-page/signup-page.component';
import { HomeComponent } from './side-bar/home/home.component';
import { SideBarComponent } from './side-bar/side-bar.component';
import { PostComponent } from './side-bar/post/post.component';
import { ExploreComponent } from './side-bar/explore/explore.component';
import { ProfileComponent } from './side-bar/profile/profile.component';
import { SearchingComponent } from './side-bar/searching/searching.component';
import { ContactComponent } from './side-bar/contact/contact.component';

const routes: Routes = [
  {
    path : '',
    component: LoginPageComponent
  },
  {
    path: 'signup',
    component: SignupComponent
  },
  {
    path: 'profile',
    component: ProfileComponent
  },
  {
    path: 'sidebar',
    component: SideBarComponent
  },
  {
    path:'post', 
    component:PostComponent
  },
  {
    path:'explore', 
  component:ExploreComponent
  },
  {
    path:'home', 
    component:HomeComponent
  },
  {
    path:'searching', 
    component:SearchingComponent
  },
  {
    path:'contact',
    component:ContactComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }