import axios from 'axios';
import { setUser, setToken } from './store';

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

export async function postLogout(user: string, token:string){
    try {
        const res = await axios.post(`/api/logout`, {mail: user, token: token});
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

export async function postInteraction(){
    try {
        const res = await axios.get(`/api/topInteractors`);
        return res.data;
    } catch (err) {
        console.log(err);
        return false
    }
}