import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Blog, BlogsService } from '../../blogs.service';
import { BehaviorSubject, map, switchMap } from 'rxjs';
import { AsyncPipe, DatePipe, JsonPipe, NgIf } from '@angular/common';
import { AuthService } from '../../../auth/auth.service';

@Component({
  selector: 'app-blogs-view',
  standalone: true,
  imports: [AsyncPipe, NgIf, JsonPipe, DatePipe],
  providers: [BlogsService, AuthService],
  templateUrl: './blogs-view.component.html',
  styleUrl: './blogs-view.component.css'
})
export class BlogsViewComponent {
  blogDetails$: BehaviorSubject<Blog | null> = new BehaviorSubject<Blog | null>(null);
  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');

  constructor(private readonly route: ActivatedRoute, private readonly blogsService: BlogsService, private readonly authService: AuthService) {
    this.getBlogDetails();
  }

  getBlogDetails() {
    this.route.params.pipe(
      map((res: any) => res.id),
      switchMap((blogID: string) => this.blogsService.getBlogDetailsByID$(blogID))
    ).subscribe(
      (res) => {
        this.blogDetails$.next(res);
      }, (err) => {
        console.log(err.error);
        this.onErrorMessage$.next(err.error.error);
        if (err.error?.error) {
        } else {
          this.onErrorMessage$.next(err.error.message);
        }
      }
    )
  }
}
