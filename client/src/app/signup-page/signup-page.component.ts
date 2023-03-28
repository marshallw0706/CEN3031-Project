import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';

interface User{
  username: string
  password: string
}

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.css']
})
export class SignupComponent {
  title = 'soundSpace';
  public submitemail = ''
  public submitname = ''
  public submitpass = ''
  public submitpass2 = ''
  public passmatcherror = false
  public emptyerror = false
  public existinguser = false
  public success = false
  public users: User[] = []
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){}

  async addUser()
  {
    this.passmatcherror = false
    this.success = false
    this.emptyerror = false
    this.existinguser = false
    if(this.submitemail == '' || this.submitname == '' || this.submitpass == '')
    {
      this.emptyerror = true
      return
    }
    if(this.submitpass != this.submitpass2)
    {
      this.passmatcherror = true
      return
    }
    this.success = true
    const users$ = await this.httpClient.get<User[]>('/api/users', {})
    this.users = await lastValueFrom(users$)
    for(var user of this.users)
    {
      if(user.username == this.submitname)
      {
        this.existinguser = true
        return
      }
    }
    console.log("submitted")
    this.httpClient.post("/api/users", {
      username: this.submitname,
      password: this.submitpass
    }).subscribe()
    if(this.success)
    {
      this.router.navigate(['sidebar'])
    }
   this.submitname = ''
  }
}
