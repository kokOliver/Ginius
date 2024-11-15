import { Component } from '@angular/core';
import { NavigationBarComponent } from './navigation-bar/navigation-bar.component';
import { ContainerDisplayComponent } from './container-display/container-display.component';
import { FooterBarComponent } from './footer-bar/footer-bar.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    NavigationBarComponent,
    ContainerDisplayComponent,
    FooterBarComponent,
  ],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'], // Fixed typo: `styleUrl` → `styleUrls`
})
export class AppComponent {
  title = 'Oliver is still here';
}
