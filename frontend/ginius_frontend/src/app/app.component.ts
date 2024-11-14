import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ButtonComponent } from './button/button.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, ButtonComponent], // Import the ButtonComponent here
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'], // Fixed typo: `styleUrl` → `styleUrls`
})
export class AppComponent {
  title = 'Oliver is still here';
}
