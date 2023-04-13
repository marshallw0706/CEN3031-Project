import { Component, EventEmitter, Output } from '@angular/core';
import { GlobalConstants } from 'src/common/global-constants';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent {
  @Output() onPostButtonClick = new EventEmitter<void>()
  @Output() onHomeButtonClick = new EventEmitter<void>();
  setRoute()
  {
    console.log("setroute")
    GlobalConstants.viewprofileid = GlobalConstants.loggedinid
  }
}
