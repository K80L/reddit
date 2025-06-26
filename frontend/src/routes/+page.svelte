<script lang="ts">
  import { onMount } from 'svelte';
  import Card from '../components/card.svelte';
  import { customFetch as fetch } from '../utils/fetch';

  let data = $state([]);

  onMount(async () => {
    try {
      data = await fetch('api/post', {});
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  });
</script>

<main class="main-content">
  {#each data as post}
    <Card post={post} />
  {/each}
  <div><p>Hello</p></div>
</main>

<style>
  .main-content {
    grid-column: 3 / span 8;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
