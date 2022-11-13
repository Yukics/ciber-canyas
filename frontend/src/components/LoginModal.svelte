<script>
    import { setLoginModal } from '../lib/store';
    import { postLogin } from '../lib/crud';
    let mail = null;

    function closeLogin(event){        
        if (!document.getElementById('form').contains(event.target)){
            if(!document.getElementById('button').contains(event.target) || !document.getElementById('mail').contains(event.target)){
                setLoginModal(false)
            }
        }
    }

    async function sendLogin(){
        const res = await postLogin(mail)
        if(res === true){
            setLoginModal(false)
        }
    }

</script>

<style>
    .login{
        height: 100%;
        width: 100%;
        position: absolute;
        z-index: 10;
        background-color: rgba(0, 0, 0, 0.361);
        display: flex;
        justify-content: center;
        align-items: center;
        align-content: center;
    }
    #form{
        width: auto;
        height: auto;
        background-color: rgb(213, 191, 191);
        border-radius: 5px;
        display: flex;
        flex-direction: column;
    }
    input[type=text] {
        width: auto;
        padding: 12px 20px;
        margin: 1vh;
        display: inline-block;
        border: 1px solid #ccc;
        border-radius: 4px;
        box-sizing: border-box;
    }
    input[type=button] {
        width: auto;
        background-color: rgb(76, 122, 175);
        color: white;
        padding: 14px 20px;
        margin: 1vh;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

</style>

<div class="login" on:mousedown={(event) => closeLogin(event)}>
    <div id="form">
        <input type="text" id="mail" name="mail" placeholder="Mail del Borja" bind:value={mail}/>
        <input type="button" id="button" value="Login" on:mousedown={() => sendLogin()}/>
    </div>
</div>