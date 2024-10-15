import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Blog, BlogsService } from '../../blogs.service';
import { BehaviorSubject, map, switchMap } from 'rxjs';
import { AsyncPipe, DatePipe, JsonPipe, NgForOf, NgIf, TitleCasePipe } from '@angular/common';
import { AuthService } from '../../../auth/auth.service';
import { MarkdownService } from 'ngx-markdown';

@Component({
  selector: 'app-blogs-view',
  standalone: true,
  imports: [AsyncPipe, NgIf, JsonPipe, DatePipe, NgForOf, TitleCasePipe,],
  providers: [BlogsService, AuthService, MarkdownService],
  templateUrl: './blogs-view.component.html',
  styleUrl: './blogs-view.component.css'
})
export class BlogsViewComponent {
  blogDetails$: BehaviorSubject<Blog | null> = new BehaviorSubject<Blog | null>(null);
  moreBlogsByUser$: BehaviorSubject<Blog[]> = new BehaviorSubject<Blog[]>([]);

  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');

  constructor(private readonly route: ActivatedRoute, private readonly router: Router, private readonly blogsService: BlogsService, private readonly authService: AuthService, private readonly markdownService: MarkdownService) {
    this.getBlogDetails();
  }

  getBlogDetails() {
    this.route.params.pipe(
      map((res: any) => res.id),
      switchMap((blogID: string) => this.blogsService.getBlogDetailsByID$(blogID))
    ).subscribe(
      (res) => {
        this.blogDetails$.next(res);
        this.getBlogsByUserID(res.userID);
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

  getBlogsByUserID(userID: number) {
    this.blogsService.getBlogsByUserID$(userID).subscribe(
      (res) => {
        this.moreBlogsByUser$.next(res);
      }, (err) => {
        this.onErrorMessage$.next(JSON.stringify(err.error));
      }
    )
  }

  onClickGoToBlogView(blog: Blog) {
    this.router.navigate(['home', 'blogs', blog.id, blog.title], {
      preserveFragment: true,
      state: blog
    })
  }

  parseMarkdown(content: string) {
    return this.markdownService.parse(content);;
  }
}
