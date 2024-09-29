<script>
  import { onMount } from 'svelte';
  import Card from '../components/card.svelte';

  onMount(async () => {
    try {
      const response = await fetch('http://localhost:8080/api/post', {});
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  });
</script>

<main class="main-content">
  <Card />
  <div>Hello</div>
</main>

<style>
  .main-content {
    grid-column: 3 / span 8;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
