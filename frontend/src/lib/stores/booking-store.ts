import { writable } from "svelte/store";

export const userId = writable(0);
export const roomId = writable(0);
export const checkIn = writable();
export const checkOut = writable();
export const totalGuest = writable(0);
export const bookingStatus = writable(0);
export const additionalItems = writable("");
export const approvalId = writable(0);