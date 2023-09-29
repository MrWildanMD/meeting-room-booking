import { tokenAuth } from "$lib/stores/login-store.js";
import { typeUser } from "$lib/stores/user-store.js";
import { redirect } from "@sveltejs/kit";
import { get } from "svelte/store";
import type { Booking } from "../../../models/booking.js";

export const load = async ({ cookies, fetch }) => {
  let bearer: string;
  const token = cookies.get("token");
  if (token?.length == 0) {
    bearer = get(tokenAuth);
  }
  bearer = "Bearer "+token;

  if (!token) {
    throw redirect(301, "/login")
  }
  if (get(typeUser) != 1) {
    throw redirect(301, "/")
  }

  const res = await fetch("http://localhost:8000/api/booking", {
    method: "GET",
    headers: {
        Authorization: bearer,
    },
  });

  if (!res.ok) {
    throw new Error(res.statusText)
  }

  const bookingBody = await res.json();
  const bookings = (bookingBody.data) as Booking[];
  
};
