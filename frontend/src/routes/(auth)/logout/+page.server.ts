import { redirect } from "@sveltejs/kit";
import { tokenAuth } from "$lib/stores/login-store.js";
import { get } from "svelte/store";

export const load = async (event) => {
    let bearer: string;
    const token = event.cookies.get("token");
    if (token?.length == 0) {
        bearer = get(tokenAuth)
    }
    bearer = "Bearer "+token;

    const res = await fetch("http://localhost:8000/api/users/logout", {
        method: "POST",
        headers: {
            Authorization: bearer,
        },
    });

    if (res.ok) {
        event.cookies.set("token", "", {
            path: "/",
        });
        throw redirect(301, "/login");
    } else {
        return new Response(await res.text());
    }
}