<script>
    import { onMount } from 'svelte';
    import { getElements } from '../lib/crud';
    import Loading from './Loading.svelte';
    import EmojiCard from './EmojiCard.svelte'
    
    let elements = []

    onMount(async () => {
		elements = await getElements();
	});

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
    {#each elements as emoji}
    <EmojiCard emoji={emoji.emoji} counter={emoji.count}/>
    {:else}
		<Loading/>
    {/each}
</div>