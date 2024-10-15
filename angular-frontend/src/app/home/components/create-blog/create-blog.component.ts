import { Component } from '@angular/core';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { BlogsService, CreateNewBlogDto } from '../../blogs.service';
import { AsyncPipe, JsonPipe, NgIf, TitleCasePipe } from '@angular/common';
import { BehaviorSubject, Observable } from 'rxjs';
import { MarkdownService } from 'ngx-markdown';

@Component({
  selector: 'app-create-blog',
  standalone: true,
  imports: [FormsModule, ReactiveFormsModule, AsyncPipe, JsonPipe, NgIf, TitleCasePipe],
  providers: [BlogsService, MarkdownService],
  templateUrl: './create-blog.component.html',
  styleUrl: './create-blog.component.css'
})
export class CreateBlogComponent {
  onSuccessMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');
  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');

  createBlogForm: FormGroup = new FormGroup({
    title: new FormControl("", [Validators.required]),
    subtitle: new FormControl("", [Validators.required]),
    content: new FormControl("", [Validators.required, Validators.minLength(20)]),
  })

  constructor(private readonly blogsService: BlogsService, private readonly markdownService: MarkdownService) { }

  onSubmit() {
    const newBlogPost: CreateNewBlogDto = this.createBlogForm.value;
    this.blogsService.createBlogpost$(newBlogPost).subscribe(
      (resp) => {
        this.onSuccessMessage$.next(`${resp.message}. You will be redirected!`);
        this.createBlogForm.reset();
        setTimeout(() => {
          window.location.replace('');
        }, 2000);
      }, (error) => {
        this.onErrorMessage$.next(JSON.stringify(error));
      }
    )
  }

  subscribeToFormUpdate$() {
    return this.createBlogForm.valueChanges as Observable<CreateNewBlogDto>;
  }

  parseMarkdownContent(content: string) {
    return this.markdownService.parse(content);
  }
}


