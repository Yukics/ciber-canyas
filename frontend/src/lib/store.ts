import { writable } from 'svelte/store';

export const user = writable(localStorage.getItem("user") ||"");
export const token = writable(localStorage.getItem("token") ||"");

export function setUser(newUser:string) {
    user.set(newUser);
}

export function setToken(newToken:string) {
    token.set(newToken);
}