import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home.component';
import { RecommendedBlogsComponent } from './components/recommended-blogs/recommended-blogs.component';
import { BlogsViewComponent } from './components/blogs-view/blogs-view.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    children: [
      {
        path: 'recommended-blogs',
        pathMatch: 'full',
        component: RecommendedBlogsComponent
      },
      {
        path: 'blogs/:id/:title',
        pathMatch: 'full',
        component: BlogsViewComponent,
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HomeRoutingModule { }
