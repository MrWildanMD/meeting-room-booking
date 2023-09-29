import { tokenAuth } from '$lib/stores/login-store.js';
import { redirect } from '@sveltejs/kit';
import { get } from 'svelte/store';
import type { Rooms } from '../../models/rooms.js';
import type { User } from '../../models/user.js';

export const load = async ({cookies, fetch}) => {
    let bearer: string;
    const token = cookies.get("token");
    if (token?.length == 0) {
        bearer = get(tokenAuth);
    }
    bearer = "Bearer "+token;

    if (!token) {
        throw redirect(301, "/login");
    }
    const [roomResponse, userResponse] = await Promise.all([
        fetch("http://localhost:8000/api/room", {
        headers: {
            Authorization: bearer,
        },
    }),
        fetch("http://localhost:8000/api/users/me", {
        headers: {
            Authorization: bearer,
        }
    }),
    ]);

    if (!roomResponse.ok) {
        throw new Error(roomResponse.statusText);
    }

    if (!userResponse.ok) {
        throw new Error(userResponse.statusText);
    }

    const roomBody = await roomResponse.json();
    const rooms = (roomBody.data) as Rooms[];
    const userBody = await userResponse.json();
    const user = (userBody.data) as User;
    return {
        rooms: rooms,
        user: user};
}

export const actions = {
    create: async ({ request, fetch, cookies }) => {
        const formData = await request.formData();
        const user_id = formData.get("user_id");
        const room_id = formData.get("room_id");
        const check_in = formData.get("check_in");
        const check_out = formData.get("check_out");
        const number_of_guest = formData.get("guest_total");
        const additional_item = formData.get("additional_items");
        const body = await JSON.stringify({"user_id":Number(user_id), "room_id":Number(room_id), check_in, check_out, "number_of_guests":Number(number_of_guest), additional_item});
        console.log(body);

        const token = cookies.get("token");
        if (!token) {
            throw redirect(301, "/login");
        }

        const res = await fetch("http://localhost:8000/api/booking", {
            body,
            method: "POST",
            headers: {
                "Content-Type":"application/json",
                Authorization: `Bearer ${token}`,
            },
        });

        if (!res.ok) {
            throw new Error(res.statusText);
        };

        return { success: true }
    }
}