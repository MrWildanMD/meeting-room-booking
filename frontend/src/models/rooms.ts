export interface Room {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: null | string;
    title: string;
    description: string;
    location: string;
    type: number;
    capacity: number;
    status: number;
    offices: null | any; // You can replace 'any' with a proper type if needed
    bookings: null | any; // You can replace 'any' with a proper type if needed
  }