import { Component, OnInit } from '@angular/core';
import {NgForm} from "@angular/forms";
import {HttpClient, HttpHeaders} from "@angular/common/http";

const httpOptions = {
  headers: new HttpHeaders({
    'Access-Control-Allow-Origin':'*'
  })
};

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less']
})
export class LoginComponent implements OnInit {

  token = '';

  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  public resolved(token: string) {
    console.log(`Resolved captcha with response: ${token}`);
    this.token = token
  }

  onAuth(signAuth : NgForm) {
    console.log(signAuth);

    this.http.post('http://127.0.0.1:8080/auth', {login:signAuth.value.login, password:signAuth.value.password, token:this.token}, httpOptions).subscribe(resp => {
      console.log(`Resp: ${JSON.stringify(resp)}`);
    })
  }
}
