<script lang="ts">
  import { onMount } from 'svelte';

  import Header from '../components/header.svelte';
  import LoginModal from '../components/login/loginModal.svelte';
  import { isLoading, isLoggedIn } from '../stores/authStore';
  import { customFetch } from '../utils/fetch';

  interface Props {
    children?: import('svelte').Snippet;
  }

  let { children }: Props = $props();

  async function checkLoginStatus() {
    const response = await customFetch('api/user/authenticate');

    const isAuthenticated = response?.isAuthenticated ?? false;

    if (isAuthenticated) {
      isLoggedIn.set(true);
      isLoading.set(false);
    } else {
      isLoggedIn.set(false);
      isLoading.set(false);
    }
  }
  onMount(async () => {
    try {
      await checkLoginStatus();
    } catch (error) {
      console.error('Error checking login status:', error);
    }
  });
</script>

<div class="app">
  {#if $isLoading}
    <div></div>
  {:else}
    <Header />
    <div class="home">
      <nav class="navigation">
        <div>
          <ul class="no-bullets">
            <li><a href="/">Home</a></li>
            <li><a href="/popular">Popular</a></li>
            <li><a href="/explore">Explore</a></li>
            <li><a href="/all">All</a></li>
          </ul>
        </div>
      </nav>
      <div class="main">
        {@render children?.()}
      </div>
    </div>
  {/if}
</div>

<LoginModal />

<style>
  * {
    box-sizing: border-box;
    color: white;
    font-family: system-ui;
    font-size: 1rem;
    line-height: 1.5;
    font-weight: 300;
  }

  a {
    color: white;
    text-decoration: none;
  }

  .app {
    background-color: #000000;
    height: 100vh;
    padding: 0 0.5rem;
  }

  .home {
    display: grid;
    grid-template-areas: 'navigation main';
    grid-template-columns: 200px 1fr;
    height: 100%;
  }

  .main {
    grid-area: main;
    padding: 20px;
    overflow-y: auto;
    display: grid;
    grid-template-columns: repeat(12, 1fr);
  }

  .navigation {
    grid-area: navigation;
    border-right: 1px solid #3e4142;
    padding-left: 20px;
    padding-top: 20px;
    position: sticky;
    top: 56px;
    left: 0;
    bottom: 0;
    width: 260px;
  }

  ul.no-bullets {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
</style>
