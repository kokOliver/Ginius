import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common'; // Import CommonModule
import * as gin from '../../assets/gin.json';

@Component({
  selector: 'container-display',
  standalone: true,
  imports: [CommonModule], // Add CommonModule to the imports array
  templateUrl: './container-display.component.html',
  styleUrls: ['./container-display.component.css'],
})
export class ContainerDisplayComponent implements OnInit {
  title = 'json-read-example';
  data: any = gin;

  ngOnInit() {
    console.log('Data', this.data);
  }
}
