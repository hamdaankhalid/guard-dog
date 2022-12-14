import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { SignupComponent } from './auth/signup/signup.component';
import { HomeComponent } from './home/home/home.component';
import { StreamerComponent } from './home/streamer/streamer.component';
import { AiComponent } from './home/ai/ai.component';
import { NotificationCenterComponent } from './home/notification-center/notification-center.component';
import { LandingPageComponent } from './public/landing-page/landing-page.component';
import { AuthGuard } from './services/auth.guard';

const publicRoutes: Routes = [
  { path: '', component: LandingPageComponent},
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent }
];

const privateRoutes: Routes = [
  { path: 'home', component: HomeComponent, canActivate: [AuthGuard] },
  { path: 'streamer', component: StreamerComponent, canActivate: [AuthGuard] },
  { path: 'ai', component: AiComponent, canActivate: [AuthGuard] },
  { path: 'notifications', component: NotificationCenterComponent, canActivate: [AuthGuard] }
];

@NgModule({
  imports: [RouterModule.forRoot([...publicRoutes, ...privateRoutes])],
  exports: [RouterModule]
})
export class AppRoutingModule { 

}
