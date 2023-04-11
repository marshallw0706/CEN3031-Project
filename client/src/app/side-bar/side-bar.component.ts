import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GlobalConstants } from 'src/common/global-constants';

@Component({
  selector: 'app-root',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css']
})
export class SideBarComponent implements OnInit {
  public hellouser = GlobalConstants.loggedinuser
  constructor(
    private router: Router
  ){}
  ngOnInit(): void {
    this.check()
  }
  title = 'homepage';

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

  visit()
  {
    console.log("In visit with user id " + GlobalConstants.loggedinid)
    GlobalConstants.viewprofileid = 1n
    console.log("In visit with user id " + GlobalConstants.loggedinid)
    this.router.navigate(['/profile'])
  }
}