import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs';
import { NG_APP_BLOGS_SERVICE_ENDPOINT } from '../environments';

@Injectable({
  providedIn: 'root'
})
export class BlogsService {
  private readonly JwtToken = window.localStorage.getItem("JWT_TOKEN");
  private readonly Username = window.localStorage.getItem("Username");


  constructor(private httpClient: HttpClient) { }

  getRecommendedBlogs$() {
    return this.httpClient.get<BlogsRecommendation>(`${NG_APP_BLOGS_SERVICE_ENDPOINT}/api/blogs/recommendations`, {
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
    return this.httpClient.get<GetBlog>(`${NG_APP_BLOGS_SERVICE_ENDPOINT}/api/blogs/${blogID}`, {
      headers: {
        'Content-Type': 'application/json, charset=UTF-8',
        'Accept': 'application/json, charset=UTF-8',
        'Authorization': `Bearer ${this.JwtToken}`
      }
    }).pipe(map(resp => resp.data));
  }

  getBlogsByUserID$(userID: number) {
    return this.httpClient.get<BlogsRecommendation>(`${NG_APP_BLOGS_SERVICE_ENDPOINT}/api/blogs/users/${userID}`, {
      headers: {
        'Content-Type': 'application/json, charset=UTF-8',
        'Accept': 'application/json, charset=UTF-8',
        'Authorization': `Bearer ${this.JwtToken}`
      }
    }
    ).pipe(
      map((resp: any) => resp.data.blogs)
    )
  }

  createBlogpost$(blog: CreateNewBlogDto) {
    return this.httpClient.post<{ data: { message: string }, statusCode: number }>(`${NG_APP_BLOGS_SERVICE_ENDPOINT}/api/blogs/create`, {
      title: blog.title,
      subtitle: blog.subtitle,
      content: blog.content
    }, {
      headers: {
        'Content-Type': 'application/json, charset=UTF-8',
        'Accept': 'application/json, charset=UTF-8',
        'Authorization': `Bearer ${this.JwtToken}`
      },
    }).pipe(
      map(resp => resp.data)
    )
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
export type CreateNewBlogDto = {
  title: string;
  subtitle: string;
  content: string;
}