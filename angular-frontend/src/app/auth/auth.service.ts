import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private AUTH_SERVICE_ENDPOINT = `http://127.0.0.1:3000`

  constructor(private httpClient: HttpClient) { }

  login$(loginDetails: LoginFormDto): Observable<LoginSuccessful> {
    return this.httpClient.post<LoginSuccessful>(`${this.AUTH_SERVICE_ENDPOINT}/api/auth/login`, { username: loginDetails.username, password: loginDetails.password }).pipe(
      tap(res => {
        window.localStorage.setItem('JWT_TOKEN', res.data.token);
        window.localStorage.setItem('USERNAME', res.data.username);
      })
    )
  }

  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => error.error);
  }
}



export type LoginFormDto = {
  username: string;
  password: string;
}

export type LoginSuccessful = {
  data: {
    token: string,
    username: string,
  },
  statusCode: number
}