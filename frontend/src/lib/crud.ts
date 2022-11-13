import axios from 'axios';
import { setUser, setToken, user, token } from './store';
import { get } from 'svelte/store'

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
        return res.data;
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
        console.log("Entr")
        const res = await axios.post(`/api/logout`, {mail: get(user), token: get(token)});
        console.log(res.data)
        if(res.data.success){
            console.log("success")
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

export async function postInteraction(){
    try {
        const res = await axios.get(`/api/topInteractors`);
        return res.data;
    } catch (err) {
        console.log(err);
        return false
    }
}