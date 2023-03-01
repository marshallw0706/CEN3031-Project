import { UserProfileComponent } from './user-profile/user-profile.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPageComponent } from './login-page/login-page.component';
import { SignupComponent } from './signup-page/signup-page.component';
import { HomeComponent } from './side-bar/home/home.component';
import { SideBarComponent } from './side-bar/side-bar.component';

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
    component: UserProfileComponent
  },
  {
    path: 'sidebar',
    component: SideBarComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }