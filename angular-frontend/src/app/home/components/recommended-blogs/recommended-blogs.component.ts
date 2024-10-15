import { Component, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Blog, BlogsService } from '../../blogs.service';
import { AsyncPipe, DatePipe, NgForOf, NgIf, TitleCasePipe } from '@angular/common';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { AuthService } from '../../../auth/auth.service';

@Component({
  selector: 'app-recommended-blogs',
  standalone: true,
  imports: [NgIf, NgForOf, AsyncPipe, TitleCasePipe, DatePipe, RouterLink],
  templateUrl: './recommended-blogs.component.html',
  styleUrl: './recommended-blogs.component.css'
})
export class RecommendedBlogsComponent {

  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');
  recommendedBlogs$: BehaviorSubject<Blog[]> = new BehaviorSubject<Blog[]>([]);

  constructor(private readonly router: Router, private readonly route: ActivatedRoute, private readonly blogsService: BlogsService, private readonly authService: AuthService) {
    this.getRecommendedBlogs();
  }

  getRecommendedBlogs() {
    this.blogsService.getRecommendedBlogs$().subscribe(
      (resp) => {
        this.recommendedBlogs$.next(resp);
      }, (err) => {
        console.error(err);
        // logout
        this.authService.logout();
        this.router.navigate(['']);
        this.onErrorMessage$.next(err.error.error);
      }
    )
  }


  onClickGoToBlogView(blog: Blog) {
    this.router.navigate(['home', 'blogs', blog.id, blog.title], {
      preserveFragment: true,
      state: blog
    })
  }

}
