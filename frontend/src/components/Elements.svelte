<script>
    import { onMount } from 'svelte';
    import { getElements } from '../lib/crud';
    import Loading from './Loading.svelte';
    import EmojiCard from './EmojiCard.svelte'
    import { emojis } from '../lib/store'

    let emojisArr = [];

    onMount(async () => {
		await getElements();
        emojisArr = $emojis;
	});

    export async function update(){
        emojisArr = $emojis;
    }

</script>

<style>
    .elements{
        width: 100%;
        background-color: rgb(76, 76, 76);
        position: absolute;
        height: 88vh;
        margin-top: 6vh;
        display: flex;
        align-items: center;
        align-content: center;
        justify-content: center;
        gap: 5vw;
    }
</style>

<meta charset="UTF-8">
<div class="elements">
    {#each emojisArr as emoji}
        <EmojiCard bind:emoji={emoji.emoji} bind:counter={emoji.count} {update}/>
    {:else}
		<Loading/>
    {/each}
</div>