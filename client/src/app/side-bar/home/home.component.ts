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

interface Comment{
  ID: BigInt
  content: string
  PostedBy: User
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
  handle: string
  created_at: BigInt
  description: string
  comments: Comment[]
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  public files: APIFile[] = []
  public usersfiles: APIFile[] = []
  public currlikes: number
  public user: User
  public currlikedby: User[]
  public users: User[]
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
    const users$ = await this.httpClient.get<User[]>('/api/users/'+GlobalConstants.loggedinid+'/following', {})
    this.users = await lastValueFrom(users$)
    if(this.users != null)
    {
    for(var user of this.users)
    {
      const usersfiles$ = await this.httpClient.get<APIFile[]>('/api/users/'+user.ID+'/files', {})
      this.usersfiles = await lastValueFrom(usersfiles$)
      if(this.usersfiles != null)
      {
        console.log("Attempting to add files")
        this.files = this.files.concat(this.usersfiles)
      }
    }
  }
    for(var file of this.files)
    {
      const user$ = this.httpClient.get<User>('/api/users/'+file.owner_id, {})
      this.user = await lastValueFrom(user$)
      file.handle = this.user.username
      const currlikedby$ = await this.httpClient.get<User[]>('/api/users/'+file.owner_id+'/files/'+file.ID+'/likedby', {})
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
    this.reversedFiles = this.files.sort((a, b) => {
      if (a.created_at < b.created_at) {
        return 1;
      }
      if (a.created_at > b.created_at) {
        return -1;
      }
      return 0;
    });

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

  gotoprofile(id: string)
  {
    GlobalConstants.viewprofileid = BigInt(id)
    this.router.navigate(['profile'])
  }

}