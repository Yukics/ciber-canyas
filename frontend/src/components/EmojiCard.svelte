<script>
    import { parseEmojiToHtml } from '../lib/misc';
    import { postInteraction } from '../lib/crud' 
    export let emoji;
    export let counter;
    export let update;

    async function handleCount(emoji){
        if(await postInteraction(emoji)){
            await update();
            console.log(emoji, counter);
        }
    }
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
        user-select: none;
        font-size:10vw;
        margin: 1vw;
    }
    .count{
        user-select: none;
        color: aliceblue;
        font-size: 2vw;
    }
</style>

<div class="card" on:mousedown={() => handleCount(emoji)}>
    <p class="emoji">{parseEmojiToHtml(emoji)}</p>
    <p class="count">Cuenta {counter}</p>    
</div>