<script>
    import Fa from 'svelte-fa/src/fa.svelte';
    import { faRightToBracket } from '@fortawesome/free-solid-svg-icons';
    import LoginModal from './LoginModal.svelte';
    import { loginModal, setLoginModal } from '../lib/store';

    function openLogin(){
        setLoginModal(true)
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
        font-size: 2vw;
        display: flex;
        align-items: center;
        align-content: flex-end;
        justify-content: flex-end;
        gap: 1vw;
    }
    .login:hover{
        cursor: pointer;

    }
    .icon{
        margin-right: 2%;
    }
    p{
        font-size: 1.6vw;
    }
</style>

<nav class="navbar">
    <div on:mousedown={() => openLogin()} class="login"><p>Login</p></div><div class="icon"><Fa icon={faRightToBracket} /></div>
</nav>
<!--El dolar es para decirle que esté atento a los cambios de la store-->
{#if $loginModal}
    <LoginModal/>
{/if}
