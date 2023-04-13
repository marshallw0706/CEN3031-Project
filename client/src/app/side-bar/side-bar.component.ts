import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { GlobalConstants } from 'src/common/global-constants';
import { HomeComponent } from './home/home.component';

@Component({
  selector: 'app-root',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css']
})
export class SideBarComponent implements OnInit {
  public hellouser = GlobalConstants.loggedinuser
  @ViewChild(HomeComponent) homeComponent: HomeComponent;
  constructor(
    private router: Router
  ){}
  ngOnInit(): void {
    this.check()
  }
  title = 'homepage';

  isPostVisible = false;

  handleHomeButtonClick() {
    if (this.isPostVisible) {
      this.togglePostVisibility();
    }
    this.homeComponent.getFiles();
  }

  togglePostVisibility() {
    this.isPostVisible = !this.isPostVisible;
  }

  check()
  {
  if(!GlobalConstants.loggedin)
    {
      this.router.navigate([''])
    }
  }

  logout()
  {
    GlobalConstants.loggedinuser = ''
    GlobalConstants.loggedinid = 1n
    GlobalConstants.loggedin = false
    console.log("logout successful")
    this.check()
  }
}