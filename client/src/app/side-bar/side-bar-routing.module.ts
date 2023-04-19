import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ExploreComponent } from './explore/explore.component';
import { HomeComponent } from './home/home.component';
import { PostComponent } from './post/post.component';
import { ProfileComponent } from './profile/profile.component';
import { CommunityComponent } from '../community/community.component';
import { SearchComponent } from './search/search.component';
import { SearchingComponent } from './searching/searching.component';

const routes: Routes = [
  {path: '', redirectTo:'home', pathMatch: 'full'},
  {path:'home', component:HomeComponent},
  {path:'profile', component:ProfileComponent},
  {path:'post', component:PostComponent},
  {path:'explore', component:ExploreComponent},
  {path: 'searching', component:SearchingComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
