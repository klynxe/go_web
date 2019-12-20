import { Component, OnInit } from '@angular/core';
import {NgForm} from "@angular/forms";
import {HttpClient, HttpHeaders} from "@angular/common/http";


const httpOptions = {
  headers: new HttpHeaders({
    'Access-Control-Allow-Origin':'*'
  })
};

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.less']
})
export class SignUpComponent implements OnInit {

  token = '';

  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  public resolved(token: string) {
    console.log(`Resolved captcha with response: ${token}`);
    this.token = token


  }

  onSignUp(signUpForm : NgForm) {
    console.log(signUpForm);

    this.http.post('http://127.0.0.1:8080/sign-up', {login:signUpForm.value.login, email: signUpForm.value.email, password:signUpForm.value.password, token:this.token}, httpOptions).subscribe(resp => {
      console.log(`Resp: ${JSON.stringify(resp)}`);
    })
  }
}
