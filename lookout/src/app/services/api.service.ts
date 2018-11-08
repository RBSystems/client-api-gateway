import { Injectable } from "@angular/core";
import { Http, Headers, RequestOptions } from "@angular/http";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";
import { Building, Room } from "../objects";

@Injectable()
export class ApiService {
  url = "http://localhost:9100";
  // url: string = '';
  options: RequestOptions;
  headers: Headers;
  constructor(private http: Http) {
    this.headers = new Headers({
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*"
    });

    this.options = new RequestOptions({ headers: this.headers });
  }

  ///// BUILDING FUNCTIONS /////
  GetBuildingList(): Observable<Building[]> {
    return this.http
      .get(this.url + "/buildings", this.options)
      .pipe(map(response => response.json()));
  }

  ///// ROOM FUNCTIONS /////
  GetRoomList(building: string): Observable<Room[]> {
    return this.http
      .get(this.url + "/buildings/" + building + "/rooms", this.options)
      .pipe(map(response => response.json()));
  }

  GetRoomByID(roomID: string): Observable<Room> {
    return this.http
      .get(this.url + "/rooms/" + roomID, this.options)
      .pipe(map(response => response.json()));
  }

  GetAllRooms(): Observable<Room[]> {
    return this.http
      .get(this.url + "/rooms", this.options)
      .pipe(map(response => response.json()));
  }

  ///// STATE & EVENTS /////
  GetState(buildingID: string, roomName: string): Observable<any> {
    console.log("getting state of", buildingID + "-" + roomName);

    return this.http
      .get(
        this.url + "/buildings/" + buildingID + "/rooms/" + roomName,
        this.options
      )
      .pipe(map(response => response.json()));
  }

  SetState(
    buildingID: string,
    roomID: string,
    payload: string
  ): Observable<any> {
    return this.http
      .put(
        this.url + "/buildings/" + buildingID + "/rooms/" + roomID,
        payload,
        this.options
      )
      .pipe(map(response => response.json()));
  }

  SubscribeToRoom(buildingID: string, roomID: string): Observable<any> {
    console.log("subscribing to", buildingID + "-" + roomID);
    return this.http
      .get(
        this.url +
          "/buildings/" +
          buildingID +
          "/rooms/" +
          roomID +
          "/subscribe",
        this.options
      )
      .pipe(map(response => response.json()));
  }

  UnsubscribeToRoom(buildingID: string, roomID: string): Observable<any> {
    console.log("unsubscribing from", buildingID + "-" + roomID);
    return this.http
      .get(
        this.url +
          "/buildings/" +
          buildingID +
          "/rooms/" +
          roomID +
          "/unsubscribe",
        this.options
      )
      .pipe(map(response => response.json()));
  }
}
