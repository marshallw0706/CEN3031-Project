import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';

interface User{
  name: string
  email: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public loginname = ''
  public submitname = ''
  public submittedname = ''
  public success = false
  public addsuccess = false
  title = 'SoundSpace';
  public users: User[] = []
  
  constructor(
    private httpClient: HttpClient
  ){}

  async checkUser()
  {
    this.success = false;
   const users$ = await this.httpClient.get<User[]>('/api/users', {
   })
   this.users = await lastValueFrom(users$)

   for (var user of this.users) {
    if(user.name == this.loginname)
    {
      this.success = true;
    }
  }
   this.loginname = ''
  }

  async addUser()
  {
    this.submittedname = this.submitname
    this.addsuccess = true
    this.httpClient.post("/api/users", {
      name: this.submitname
    }).subscribe()

   this.submitname = ''
  }

}

