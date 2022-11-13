<script>
    import { parseEmojiToHtml } from '../lib/misc';
    export let emoji;
    export let counter;
</script>

<style>
    .card{
        border-radius: 15px;
        background-color: rgba(0, 0, 0, 0.249);
        display: flex;
        flex-direction: column;
        justify-content:space-around;
        align-content: center;
        align-items: center;
        gap: 0vw;
    }
    .card:hover{
        cursor: pointer;
        background-color: rgba(0, 0, 0, 0.171);
    }
    .card:active{
        cursor: pointer;
        background-color: rgba(116, 136, 252, 0.171);
    }
    .emoji{
        font-size:10vw;
        margin: 1vw;
    }
    .count{
        color: aliceblue;
        font-size: 2vw;
    }
</style>

<div class="card">
    <p class="emoji">{parseEmojiToHtml(emoji)}</p>
    <p class="count">Cuenta {counter}</p>    
</div>