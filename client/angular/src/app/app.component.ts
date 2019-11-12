import { Component, OnInit } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http"

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less']
})
export class AppComponent implements OnInit {

  public menuItems: Array<Object> = [
    {
      title: 'Главная',
      route: '/'
    },
    {
      title: 'Авторизация',
      route: '/auth/login'
    },
    {
      title: 'Регистрация',
      route: '/auth/sign-up'
    },
  ];



  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

}
