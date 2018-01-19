import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import 'rxjs/add/operator/map';
import { environment } from '@env/environment';

@Injectable()
export class MovieService {
  private baseUrl: string = environment.api;

  constructor(private http: Http){}

  public getMovies(id?: number){
    return this.http
      .post(`${this.baseUrl}`, {ID: id})
      .map(res => {
        return res.json()
      })
  }
}
