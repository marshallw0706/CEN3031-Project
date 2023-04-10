import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';
import { compare } from 'bcryptjs';
import { GlobalConstants } from 'src/common/global-constants';


interface User{
  ID: BigInt
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
  public turnred = false
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
        GlobalConstants.loggedin = true
        GlobalConstants.loggedinid = user.ID
        GlobalConstants.loggedinuser = this.username

      }
    }
    this.username = ''
    this.password = ''
    if(this.success)
    {
      this.router.navigate(['sidebar'])
    }
    else
    {
      this.turnred = true
      console.log("turn red")
    }
  }
  
}