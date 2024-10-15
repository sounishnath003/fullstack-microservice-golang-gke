import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BlogsService {
  private readonly JwtToken = window.localStorage.getItem("JWT_TOKEN");
  private readonly Username = window.localStorage.getItem("Username");

  private BLOGS_SERVICE_ENDPOINT = `http://127.0.0.1:3001`;

  constructor(private httpClient: HttpClient) { }

  getRecommendedBlogs$() {
    return this.httpClient.get<BlogsRecommendation>(`${this.BLOGS_SERVICE_ENDPOINT}/api/blogs/recommendations`, {
      headers: {
        'Content-Type': 'application/json, charset=UTF-8',
        'Accept': 'application/json, charset=UTF-8',
        'Authorization': `Bearer ${this.JwtToken}`
      },
    }).pipe(
      map(resp => resp.data)
    )
  }

  getBlogDetailsByID$(blogID: string) {
    return this.httpClient.get<GetBlog>(`${this.BLOGS_SERVICE_ENDPOINT}/api/blogs/${blogID}`, {
      headers: {
        'Content-Type': 'application/json, charset=UTF-8',
        'Accept': 'application/json, charset=UTF-8',
        'Authorization': `Bearer ${this.JwtToken}`
      }
    }).pipe(map(resp => resp.data));
  }
}


export interface BlogsRecommendation {
  data: Blog[];
  statusCode: number;
}

export interface GetBlog {
  data: Blog;
  statusCode: number;
}

export interface Blog {
  id: number;
  userID: number;
  title: string;
  subtitle: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
}