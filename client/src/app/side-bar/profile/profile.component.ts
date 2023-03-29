import { Component } from '@angular/core';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  profile = {
    name: 'Sound Space  User',
    jobTitle: 'Master Musician',
    userImage: 'https://via.placeholder.com/350x150',
    description: 'Insert profile page description here.'
  };
  editMode = false;

  onFileSelected(event: any) {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = () => {
      this.profile.userImage = reader.result as string;
    };
    reader.readAsDataURL(file);
  }

  toggleEditMode() {
    this.editMode = !this.editMode;
  }
}
