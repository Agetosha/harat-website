import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ApiService } from './services/api';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.html',
  styleUrl: './app.scss'
})

export class App implements OnInit {
  protected title = 'dk-harat-angular-project';
  message = '';

  constructor(private api: ApiService) {}

  ngOnInit() {
    this.api.getHello().subscribe((res: any) => {
      this.message = res.message;
    });
  }
}

