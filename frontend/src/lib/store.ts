import { writable } from 'svelte/store';

export const user = writable(sessionStorage.getItem("user") || "");
export const token = writable(sessionStorage.getItem("token") || "");
export const loginModal = writable(sessionStorage.getItem("loginModal") || false);

export function setUser(newUser:string) {
    user.set(newUser);
}

export function setToken(newToken:string) {
    token.set(newToken);
}

export function setLoginModal(newEnable: boolean){
    loginModal.set(newEnable)
}