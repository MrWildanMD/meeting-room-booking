import type { Room } from "./rooms";
import type { User } from "./user";

export interface Booking {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: null | string;
    user_id: number;
    room_id: number;
    check_in: string;
    check_out: string;
    number_of_guests: number;
    booking_status: number;
    additional_item: string;
    approval_id: number;
    user: User;
    room: Room;
    notifications: null | any; // You can replace 'any' with a proper type if needed
  }