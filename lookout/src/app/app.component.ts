import { Component, OnInit } from "@angular/core";
import { Building, Room } from "./objects";
import { ApiService } from "./services/api.service";
import { SocketService } from "./services/socket.service";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.scss"]
})
export class AppComponent implements OnInit {
  title = "app";
  building: Building = new Building();
  room: Room = new Room();
  buildingList: Building[] = [];
  roomList: Room[] = [];

  roomState: any;
  payload: string;

  message: string;
  events: string[] = [];

  oldSub: string;
  newSub: string;

  constructor(private api: ApiService, private socket: SocketService) {
    this.UpdateFromEvents();
  }

  ngOnInit() {
    this.GetBuildings();
    this.payload = `{
  "displays": [
    {
      "name": "D1",
      "power": "on",
      "input": "HDMI1",
      "blanked": false
    }
  ]
}
    `;
  }

  GetBuildings() {
    this.api.GetBuildingList().subscribe(val => {
      if (val != null) {
        this.buildingList = val;
      }
    });
  }

  GetRoomList() {
    this.api.GetRoomList(this.building._id).subscribe(val => {
      if (val != null) {
        this.roomList = val;
      }
    });
  }

  GetState() {
    const roomSplit = this.room._id.split("-");
    this.api.GetState(roomSplit[0], roomSplit[1]).subscribe(val => {
      this.message = JSON.stringify(val, null, 2);
    });
  }

  SetState() {
    const roomSplit = this.room._id.split("-");
    const data = JSON.parse(this.payload);

    console.log("setting state", data);
    this.api.SetState(roomSplit[0], roomSplit[1], data).subscribe(val => {
      this.message = JSON.stringify(val, null, 3);
    });
  }

  UpdateFromEvents() {
    this.socket.getEventListener().subscribe(event => {
      if (event != null && event.data != null) {
        const e = event.data;

        this.events.push(JSON.stringify(e, null, 3));
      }
    });
  }

  UpdateRoomSubscription() {
    if (this.room._id === null) {
      return;
    }

    this.newSub = this.room._id;

    if (this.oldSub != null) {
      const r = this.oldSub.split("-");
      this.api.UnsubscribeToRoom(r[0], r[1]).subscribe(() => {
        this.events = [];
      });
    }

    if (this.newSub != null) {
      const r = this.newSub.split("-");
      this.api.SubscribeToRoom(r[0], r[1]).subscribe(() => {
        this.oldSub = this.newSub;
        this.events = [];

        this.GetState();
      });
    }
  }
}
