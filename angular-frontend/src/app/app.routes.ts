import { Routes } from '@angular/router';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

export const routes: Routes = [
    {
        path: '',
        pathMatch: 'full',
        redirectTo: 'auth/login'
    },
    {
        path: 'auth',
        loadChildren: () => import('./auth/auth.module').then(m => m.AuthModule),
    },
    {
        path: "home",
        loadChildren: () => import('./home/home.module').then(m => m.HomeModule)
    },
    {
        path: "**",
        component: PageNotFoundComponent
    }
];
