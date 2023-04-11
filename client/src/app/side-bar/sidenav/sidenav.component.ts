import { Component } from '@angular/core';
import { GlobalConstants } from 'src/common/global-constants';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent {
  setRoute()
  {
    console.log("setroute")
    GlobalConstants.viewprofileid = GlobalConstants.loggedinid
  }
}
