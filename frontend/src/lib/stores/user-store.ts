import { get, writable } from "svelte/store";

export const id = writable(0);
export const name = writable("");
export const privyId = writable("");
export const typeUser = writable(0);
export const email = writable("");
export const teleId = writable("");

export function isAdmin() {
    return get(typeUser) == 1;
}