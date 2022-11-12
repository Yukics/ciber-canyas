<script>
    import { onMount } from 'svelte';
    import { getElements } from '../lib/crud';
    import { parseEmojiToHtml } from '../lib/misc'
    
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
    }
    h1{
        font-size: 10vw;
    }
    .button:hover{
        cursor: pointer;
    }
</style>

<meta charset="UTF-8">
<div class="elements">
    {#each elements as emoji}
        <div class="button">
            <h1>{parseEmojiToHtml(emoji.Emoji)}</h1>
        </div>
    {:else}
		<p>loading...</p>
    {/each}
</div>