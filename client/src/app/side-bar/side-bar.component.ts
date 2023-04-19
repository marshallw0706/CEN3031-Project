import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { GlobalConstants } from 'src/common/global-constants';
import { HomeComponent } from './home/home.component';
import { lastValueFrom } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http'

interface User{
  ID: BigInt
  username: string
  password: string
}

@Component({
  selector: 'app-root',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css']
})
export class SideBarComponent implements OnInit {
  public hellouser = GlobalConstants.loggedinuser
  @ViewChild(HomeComponent) homeComponent: HomeComponent;
  constructor(
    private router: Router,
    private httpClient: HttpClient
  ){}
  public users: User[] = []
  
  ngOnInit(): void {
    this.check()
    this.startingUsers()
  }
  title = 'homepage';

  isPostVisible = false;

  async fetchImageAsBlob(imagePath: string): Promise<Blob> {
    const response = await fetch(imagePath);
    const blob = await response.blob();
    return blob;
  }

  async createUserProfile(username: string, name: string, jobTitle: string, description: string, profilePicPath: string): Promise<void> {
    const response: any = await this.httpClient.post('/api/users', { "username": username }).toPromise();
    await this.httpClient.put("/api/users/" + response.ID + "/profile", {
      "name": name,
      "jobtitle": jobTitle,
      "description": description
    }).toPromise();
  
    let idArray = GlobalConstants.idArray;
    idArray.push(response.ID);
    GlobalConstants.idArray = idArray;
    console.log(`${name} has id: ${response.ID}`);
  
    const formData = new FormData();
    const imageBlob = await this.fetchImageAsBlob(profilePicPath);
    formData.append('file', imageBlob, "profilepic.png");
    await this.httpClient.post("/api/users/" + response.ID + "/files/upload", formData).toPromise();
  }

  async startingUsers()
  {
    const users$ = await this.httpClient.get<User[]>('/api/users', {})
    const users = await lastValueFrom(users$)
    if(users)
    {
      this.users = users
    if(this.users.length === 1)
    {
      await this.createUserProfile("ryanjones", "Ryan Jones", "Musician", "Check out my new album!", "../../assets/image/1.png");
      await this.createUserProfile("jazzyj", "Jazzy J", "Musician", "Check out my new album!", "../../assets/image/2.png");
      await this.createUserProfile("shawnmakesmusic", "Shawn", "Musician", "Check out my new album!", "../../assets/image/boy.jpeg");
      await this.createUserProfile("hailz_official", "Hailz", "Musician", "Check out my new album!", "../../assets/image/4.png");
      await this.createUserProfile("morgan_combs", "Morgan Combs", "Musician", "Check out my new album!", "../../assets/image/5.png");
      await this.createUserProfile("SABBY", "SABBY", "Musician", "Check out my new album!", "../../assets/image/eye.jpeg");
      await this.createUserProfile("james_stone", "James Stone", "Musician", "Check out my new album!", "../../assets/image/higher.jpeg");
      await this.createUserProfile("kenya_eastnotwest", "Kenya East", "Musician", "Check out my new album!", "../../assets/image/hop.jpeg");
      await this.createUserProfile("michaeljames", "Michael James", "Musician", "Check out my new album!", "../../assets/image/sky.jpeg");
      await this.createUserProfile("johngho", "Johngho", "Musician", "Check out my new album!", "../../assets/image/sphere.jpeg");
      await this.createUserProfile("jeanne_lebras", "Jeanne lebras", "Musician", "Check out my new album!", "../../assets/image/stranger.jpeg");
      await this.createUserProfile("iamkarl", "Karl", "Musician", "Check out my new album!", "../../assets/image/sunset.jpeg");
      await this.createUserProfile("ciara", "Ciara", "Musician", "Check out my new album!", "../../assets/image/girl.jpeg");
      await this.createUserProfile("kieth___lori", "Kieth Lori", "Musician", "Check out my new album!", "../../assets/image/timeaway.png");
      await this.createUserProfile("RPL", "RPL", "Musician", "Check out my new album!", "../../assets/image/15.png");
      await this.createUserProfile("thetwo_official", "The Two", "Musician", "Check out my new album!", "../../assets/image/first.jpeg");
      await this.createUserProfile("RoyEdri", "Roy Edri", "Musician", "Check out my new album!", "../../assets/image/ghost.jpeg");
      await this.createUserProfile("GeceyeMiKussem__", "Geceye Mi Kussem", "Musician", "Check out my new album!", "../../assets/image/patron.png");
      await this.createUserProfile("The_DrMic", "Dr Mic", "Musician", "Check out my new album!", "../../assets/image/mic.jpeg");
      await this.createUserProfile("Ilove21savage", "Drizzy Drake :P", "Musician", "Check out my new album!", "../../assets/image/20.png");
      
    }
  }
  }

  handleHomeButtonClick() {
    if (this.isPostVisible) {
      this.togglePostVisibility();
    }
    this.homeComponent.getFiles();
  }

  togglePostVisibility() {
    this.isPostVisible = !this.isPostVisible;
  }

  check()
  {
  if(!GlobalConstants.loggedin)
    {
      this.router.navigate([''])
    }
  }

  logout()
  {
    GlobalConstants.loggedinuser = ''
    GlobalConstants.loggedinid = 1n
    GlobalConstants.loggedin = false
    console.log("logout successful")
    this.check()
  }
}