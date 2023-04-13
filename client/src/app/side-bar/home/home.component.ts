import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';
import { GlobalConstants } from 'src/common/global-constants';

interface User{
  ID: BigInt
  username: string
  password: string
}

interface APIFile{
  ID: BigInt
	filename: string
  owner_id: string
	type: string
	data: string
  likes: number
  likedby: User[]
  liked: boolean
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  public user = GlobalConstants.loggedinuser
  public files: APIFile[] = []
  public currlikes: number
  public currlikedby: User[]
  public reversedFiles: APIFile[] = [];
  public uploadfile: File
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){}

  ngOnInit(): void {
    this.getFiles()
  }
  
  async getFiles()
  {
    const files$ = await this.httpClient.get<APIFile[]>('/api/users/'+GlobalConstants.loggedinid+'/files', {})
    this.files = await lastValueFrom(files$)
    for(var file of this.files)
    {
      const currlikedby$ = await this.httpClient.get<User[]>('/api/users/'+GlobalConstants.loggedinid+'/files/'+file.ID+'/likedby', {})
      this.currlikedby = await lastValueFrom(currlikedby$)
      if(this.currlikedby != null)
      {
        console.log("liked by not null")
        for(var user of this.currlikedby)
        {
          if(user.ID == GlobalConstants.loggedinid)
          {
            console.log("Liked by user")
            file.liked = true
          }
        }
    }
      if(file.liked == null)
      {
        file.liked = false
      }
    }
    this.reversedFiles = this.files.slice().reverse();

  }

  async addLike(likefile: APIFile)
  {
    likefile.liked = !likefile.liked;

  if (likefile.liked) {
    likefile.likes++
    this.httpClient.post('/api/users/' + GlobalConstants.loggedinid + '/like/' + likefile.owner_id + '/' + likefile.ID, {}).subscribe(
      (response: any) => console.log("successful like: " + response),
      (error) => console.log("failure to like: " + error)
    );
  } else {
    likefile.likes--
    this.httpClient.delete('/api/users/' + GlobalConstants.loggedinid + '/unlike/' + likefile.owner_id + '/' + likefile.ID, {}).subscribe(
      (response: any) => console.log("successful unlike: " + response),
      (error) => console.log("failure to unlike: " + error)
    );
  }

  }

}