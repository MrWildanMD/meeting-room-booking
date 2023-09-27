import { writable } from "svelte/store";

export const tokenAuth = writable("");
export const isLoggedIn = writable(false);