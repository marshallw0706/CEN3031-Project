import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { GlobalConstants } from 'src/common/global-constants';


@Component({
  selector: 'app-explore',
  templateUrl: './explore.component.html',
  styleUrls: ['./explore.component.css']
})
export class ExploreComponent {
  searchText: string = '';
  public stringids = GlobalConstants.idArray

  ngOnInit(): void {
    for(var id of this.stringids)
    {
      console.log("id: " + id)
    }
    for(var id of GlobalConstants.idArray)
    {
      console.log("id: " + id)
    }
  }
  
  constructor(
    private router: Router
  ){}


  onSearchTextEntered(searchValue: string){
    this.searchText = searchValue;
    //console.log(this.searchText);
  }

  gotoprofile(id: number)
  {
    console.log("Going to profile id: " + id)
    GlobalConstants.viewprofileid = BigInt(id)
    this.router.navigate(['profile'])
  }
}

