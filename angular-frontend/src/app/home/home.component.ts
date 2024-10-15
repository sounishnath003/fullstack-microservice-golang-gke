import { Component } from '@angular/core';
import { ActivatedRoute, Router, RouterOutlet } from '@angular/router';
import { BehaviorSubject } from 'rxjs';
import { BlogsService } from './blogs.service';
import { AsyncPipe, NgIf } from '@angular/common';
import { AuthService } from '../auth/auth.service';
import { RecommendedBlogsComponent } from './components/recommended-blogs/recommended-blogs.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [RouterOutlet, NgIf, AsyncPipe, RecommendedBlogsComponent],
  providers: [BlogsService, AuthService],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');
  constructor(private readonly router: Router, private readonly route: ActivatedRoute) {
    this.router.navigate(['recommended-blogs'], {
      relativeTo: this.route
    })
  }
}
