<script>
    import Fa from 'svelte-fa/src/fa.svelte'
    import { faRightToBracket } from '@fortawesome/free-solid-svg-icons'
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
    .icon{
        margin-right: 2%;
    }
    p{
        font-size: 1.6vw;
    }
</style>

<nav class="navbar">
    <div><p>Login</p></div><div class="icon"><Fa icon={faRightToBracket} /> </div>
</nav>