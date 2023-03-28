import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';
import { compare } from 'bcryptjs';


interface User{
  username: string
  password: string
}

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
  public username = ''
  public password = ''
  public success = false
  
  public users: User[] = []
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){}


  async checkUser()
  {
    const users$ = await this.httpClient.get<User[]>('/api/users', {})
    this.users = await lastValueFrom(users$)
    for(var user of this.users)
    {
      if(user.username == this.username && await compare(this.password, user.password))
      {
        this.success = true
      }
    }
    this.username = ''
    this.password = ''
    if(this.success)
    {
      this.router.navigate(['sidebar'])
    }
  }
  
}

