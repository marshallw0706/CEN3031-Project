import { Component, EventEmitter, Output } from '@angular/core';
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
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent {
  @Output() onPostButtonClick = new EventEmitter<void>()
  public data = ''
  public description = ''
  public files: APIFile[] = []
  public fileName = ''
  public uploadfile: File
  public filedisplay: APIFile = {
    filename: '',
    type: '',
    data: ''
  }
  public filedisplay2: APIFile = {
    filename: '',
    type: '',
    data: ''
  }
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ){}

  onFileSelected(event) {
    this.uploadfile = event.target.files[0]
}

  uploadFile()
  {
    const formData = new FormData()
    formData.append('file', this.uploadfile)
    this.httpClient.post("/api/users/"+GlobalConstants.loggedinid+"/files/upload", formData).subscribe(
      (response: any) => this.httpClient.put("/api/users/"+GlobalConstants.loggedinid+"/files/"+response.ID, {"description": this.description}).subscribe(),
      (error) => console.log("error uploading file")
    )
  }
}