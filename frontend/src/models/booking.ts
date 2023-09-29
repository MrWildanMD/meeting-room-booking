export interface Booking {
    ID: number,
    user_id: number,
    room_id: number,
    check_in: Date,
    check_out: Date,
    number_of_guest: number,
    booking_status: number,
    additional_item: string,
    CreatedAt: Date
}