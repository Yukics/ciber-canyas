<script>
    import Fa from 'svelte-fa/src/fa.svelte';
    import { faRightToBracket, faUserCircle } from '@fortawesome/free-solid-svg-icons';
    import LoginModal from './LoginModal.svelte';
    import { loginModal, setLoginModal, user, token } from '../lib/store';
    import { postLogout } from '../lib/crud';

    function openLogin(){
        setLoginModal(true)
    }

    function logout(){
        postLogout($user, $token)
    }

</script>

<style>
    nav{
        position: absolute;
        background-color: rgb(47, 47, 47);
        height: 6vh;
        width: 100%;
        margin-top: 0vh;
        color: aliceblue;
    }
    .login{
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        /* align-items: center; */
        align-content: center;
        gap: 1vw;
    }
    .login:hover{
        cursor: pointer;

    }
    .icon{
        margin-right: 2%;
        font-size: 2vw;
        margin-top: 0.8vh;
    }
    p{
        font-size: 1.6vw;
        margin-top: 1.2vh;
    }
</style>

{#if $user !== ""}
    <nav class="navbar">
        <div on:mousedown={() => logout()} class="login">
            <div><p>{$user}</p></div>
            <div class="icon"><Fa icon={faUserCircle}/></div>
        </div>
    </nav>
{:else}
    <nav class="navbar">
        <div on:mousedown={() => openLogin()} class="login">
            <div><p>Login</p></div>
            <div class="icon"><Fa icon={faRightToBracket} /></div>
        </div>
    </nav>
{/if}

<!--El dolar es para decirle que esté atento a los cambios de la store-->
{#if $loginModal}
    <LoginModal/>
{/if}
