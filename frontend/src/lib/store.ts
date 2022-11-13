import { writable } from 'svelte/store';

export const user = writable(localStorage.getItem("user") || "");
export const token = writable(localStorage.getItem("token") || "");
export const emojis = writable(localStorage.getItem("emojis") || null);
export const loginModal = writable(localStorage.getItem("loginModal") || false);

export function setUser(newUser:string) {
    user.set(newUser);
}

export function setToken(newToken:string) {
    token.set(newToken);
}

export function setLoginModal(newEnable: boolean){
    loginModal.set(newEnable)
}