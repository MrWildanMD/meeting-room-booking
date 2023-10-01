import { tokenAuth } from "$lib/stores/login-store.js";
import { typeUser } from "$lib/stores/user-store.js";
import { redirect } from "@sveltejs/kit";
import { get } from "svelte/store";
import type { Booking } from "../../../models/booking.js";
import type { Room } from "../../../models/rooms.js";

let url = "http://localhost:8000/api/"

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

  const [bookingRes, roomRes] = await Promise.all([
    fetch(url+"booking", {
      headers: {
        Authorization: bearer,
      },
    }),
    fetch(url+"room", {
      method: "GET"
    }),
  ]);

  if (!bookingRes.ok) {
    throw new Error(bookingRes.statusText)
  }
  if (!roomRes.ok) {
    throw new Error(roomRes.statusText)
  }

  const bookingBody = await bookingRes.json();
  const bookings = (bookingBody.data) as Booking[];
  const roomBody = await roomRes.json();
  const rooms = (roomBody.data) as Room[];
  return { bookings, rooms }
};

export const actions = {
  create: async(event) => {
    const token = event.cookies.get("token");
    const formData = await event.request.formData();
    const title = formData.get("title");
    const description = formData.get("description");
    const location = formData.get("location");
    const type = formData.get("type");
    const capacity = formData.get("capacity");
    const body = await JSON.stringify({"title":title, "description":description, "location":location, "type":Number(type), "capacity":Number(capacity), "status":0});

    const res = await fetch(url+"room", {
      body,
      method: "POST",
      headers: {
        "Content-Type":"application/json",
        Authorization: `Bearer ${token}`
      },
    });

    if (res.ok) {
      const resBody = await res.json();
      return { message: resBody.message }
    }

    return {
      error: await res.text(),
    }
  },
  update: async(event) => {
    const token = event.cookies.get("token");
    const formData = await event.request.formData();
    const id = formData.get("id");
    const title = formData.get("title");
    const description = formData.get("description");
    const location = formData.get("location");
    const type = formData.get("type");
    const capacity = formData.get("capacity");
    const body = await JSON.stringify({"title":title, "description":description, "location":location, "type":Number(type), "capacity":Number(capacity)});

    const res = await fetch(url+"room/"+id, {
      body,
      method: "PUT",
      headers: {
        "Content-Type":"application/json",
        Authorization: `Bearer ${token}`
      },
    });

    if (res.ok) {
      const resBody = await res.json();
      return { message: resBody.message }
    }

    return {
      error: await res.text(),
    }
  },
  delete: async(event) => {
    const token = event.cookies.get("token");
    const formData = await event.request.formData();
    const id = formData.get("id");

    const res = await fetch(url+"room/"+id, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`
      }
    });

    return {
      error: await res.text(),
    }
  },
  approve: async(event) => {
    const token = event.cookies.get("token");
    const formData = await event.request.formData();
    const booking_id = formData.get("booking_id");
    const room_id = formData.get("room_id");
    const bookingBody = JSON.stringify({"booking_status":1});
    const roomBody = JSON.stringify({"status":1});

    const [bookingRes, roomRes] = await Promise.all([
      fetch(url+"booking/"+booking_id, {
        body: bookingBody,
        method: "PUT",
        headers: {
          "Content-Type":"application/json",
          Authorization: `Bearer ${token}`
        }
      }),
      fetch(url+"room/"+room_id, {
        body: roomBody,
        method: "PUT",
        headers: {
          "Content-Type":"application/json",
          Authorization: `Bearer ${token}`
        }
      }),
    ])
    // const res = await fetch(url+"booking/"+booking_id, {
    //   body,
    //   method: "PUT",
    //   headers: {
    //     "Content-Type":"application/json",
    //     Authorization: `Bearer ${token}`
    //   }
    // });

    // if (res.ok) {
    //   const resBody = await res.json();
    //   return { message: resBody.message }
    // }

    return {
      error: await bookingRes.text(),
    }
  },
  decline: async(event) => {
    const token = event.cookies.get("token");
    const formData = await event.request.formData();
    const id = formData.get("id");

    const res = await fetch(url+"booking/"+id, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`
      }
    });

    return {
      error: await res.text(),
    }
  }
}
