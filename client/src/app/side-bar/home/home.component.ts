import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';
import { GlobalConstants } from 'src/common/global-constants';

interface APIFile{
	filename: string
	type: string
	data: string
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  public user = GlobalConstants.loggedinuser
  public files: APIFile[] = []
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
    this.reversedFiles = this.files.slice().reverse();

  }

}