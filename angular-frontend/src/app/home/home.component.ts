import { Component } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';
import { BehaviorSubject } from 'rxjs';
import { Blog, BlogsService } from './blogs.service';
import { AsyncPipe, DatePipe, JsonPipe, NgFor, NgForOf, NgIf, TitleCasePipe } from '@angular/common';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [RouterOutlet, JsonPipe, AsyncPipe, NgIf, NgForOf, DatePipe, TitleCasePipe],
  providers: [BlogsService],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');
  recommendedBlogs$: BehaviorSubject<Blog[]> = new BehaviorSubject<Blog[]>([]);

  constructor(private readonly blogsService: BlogsService, private readonly router: Router) {
    this.getRecommendedBlogs();
  }

  getRecommendedBlogs() {
    this.blogsService.getRecommendedBlogs$().subscribe(
      (resp) => {
        console.log(resp);

        this.recommendedBlogs$.next(resp);
      }, (err) => {
        console.error(err);
        window.localStorage.clear();
        this.router.navigate(['/']);
        this.onErrorMessage$.next(err.error.error);
      }
    )
  }
}
