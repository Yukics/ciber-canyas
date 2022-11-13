import axios from 'axios';
import { setUser, setToken, setEmojis, user, token } from './store';
import { get } from 'svelte/store';

export async function getInteractors(){
    try {
        const res = await axios.get(`/api/topInteractors`);
        return res.data;
    } catch (err) {
        console.log(err);
        return false
    }
}

export async function getElements(){
    try {
        const res = await axios.get(`/api/emojis`);
        if(res.data){
            setEmojis(res.data)
            return true
        } else {
            return false
        }
    } catch (err) {
        console.log(err);
        return false
    }
}

export async function postLogin(mail: string){
    try {
        const res = await axios.post(`/api/login`,{mail: mail});
        if(res.data.success){
            setUser(mail)
            setToken(res.data.token)
            return true
        } else {
            return false
        }
    } catch (err) {
        console.log(err);
        return false
    }
}

export async function postLogout(){
    try {
        const res = await axios.post(`/api/logout`, {mail: get(user), token: get(token)});
        if(res.data.success){
            setUser("")
            setToken("")
            return true
        } else {
            return false
        }
    } catch (err) {
        console.log(err);
        return false
    }
}

export async function postInteraction(emoji: string){
    try {
        const res = await axios.post(`/api/interaction`, {mail: get(user), token: get(token), emoji: emoji});
        if(res.data.success){
            await getElements();
            return true
        } else {
            return false
        }
    } catch (err) {
        console.log(err);
        return false
    }
}