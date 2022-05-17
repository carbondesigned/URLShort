<script>
  import Item from './Item.svelte';
  import { onMount } from 'svelte';

  let shorts = [];
  onMount(async () => {
    const response = await fetch('http://127.0.0.1:3000/shorts');
    const data = await response.json();
    console.log(data.data);
    shorts = data.data;
  });
</script>

{#each shorts as short}
  <Item {short} />
{/each}
