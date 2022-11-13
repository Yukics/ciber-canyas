import axios from 'axios';

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

export async function postLogin(){
    try {
        const res = await axios.get(`/api/topInteractors`);
        return res.data;
    } catch (err) {
        console.log(err);
        return false
    }
}

export async function postLogout(){
    try {
        const res = await axios.get(`/api/topInteractors`);
        return res.data;
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