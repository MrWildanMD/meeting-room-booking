export interface User {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: null | string;
    name: string;
    privy_id: string;
    type_user: number;
    email: string;
    tele_id: string;
    bookings: null | any; // You can replace 'any' with a proper type if needed
    notifications: null | any; // You can replace 'any' with a proper type if needed
  }