<script>
    import {setUser, setToken, setLoginModal} from '../lib/store';
    function closeLogin(event){
        if (!document.getElementById('form').contains(event.target)){
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
    form{
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
    <form id="form">
        <input type="text" id="mail" name="mail" placeholder="Mail del Borja"/>
        <input type="button" value="Login"/>
    </form>
</div>