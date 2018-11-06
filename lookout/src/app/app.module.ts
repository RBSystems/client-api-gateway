import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import { MatFormFieldModule, MatToolbarModule, MatButtonModule, MatInputModule, MatAutocompleteModule, MatOptionModule } from '@angular/material';
import { ApiService } from './services/api.service';
import { SocketService } from './services/socket.service';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MatSelectModule } from '@angular/material/select';


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatSelectModule,
    FormsModule,
    HttpModule,
    MatToolbarModule,
    MatButtonModule,
    MatInputModule,
    MatAutocompleteModule,
    MatOptionModule
  ],
  providers: [
    ApiService,
    SocketService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
