import { Component } from '@angular/core';


@Component({
  selector: 'app-explore',
  templateUrl: './explore.component.html',
  styleUrls: ['./explore.component.css']
})
export class ExploreComponent {
  searchText: string = '';
 


  onSearchTextEntered(searchValue: string){
    this.searchText = searchValue;
    //console.log(this.searchText);
  }
}
