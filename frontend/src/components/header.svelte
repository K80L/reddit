<script lang="ts">
  import { isLoggedIn, logout } from '../stores/authStore';
  import { openModal } from '../stores/modalStore';
  import { customFetch as fetch } from '../utils/fetch';

  async function handleLogout() {
    try {
      const { loggedOut } = await fetch('api/user/logout');
      if (loggedOut) {
        logout();
      }
    } catch (error) {
      console.error('Error logging out:', error);
    }
  }
</script>

<header class="header">
  <div class="header__left">
    <img class="logo" src="/reddit-seeklogo.svg" width="55" height="55" alt="Reddit Logo" />
    <img class="logo-reddit" src="/reddit.svg" width="80" height="40" alt="Reddit" />
  </div>
  <div class="searchbar"><input type="text" placeholder="Search Reddit" /></div>
  <div>
    {#if $isLoggedIn}
      <nav class="nav-icons">
        <a href="A">A</a>
        <a href="B">B</a>
        <a href="C">C</a>
        <a href="D">D</a>
        <a href="E">Create +</a>
        <button onclick={handleLogout} class="logout__btn">Log Out</button>
      </nav>
    {:else}
      <div>
        <button onclick={openModal} class="login__btn">Log In</button>
      </div>
    {/if}
  </div>
</header>

<style>
  .header {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    padding: 0 1rem 0 0.25rem;
    border-bottom: 1px solid #3e4142;
  }

  .header__left {
    display: flex;
    align-items: center;
  }

  .login__btn {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 20px;
    background-color: #ae2c00;
    color: white;
    cursor: pointer;
    font-size: 0.9rem;
  }

  .logo {
    grid-column: 1;
  }

  .logo-reddit {
    filter: invert(1) sepia(1) saturate(1) hue-rotate(180deg);
  }

  .searchbar {
    grid-column: 2;
    text-align: center;
    height: 56px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .searchbar input {
    width: 100%;
    max-width: 500px;
    padding: 0.5rem;
    border: none; /* <-- This thing here */
    border: solid 1px #ccc;
    border-radius: 20px;
    background-color: #333d41;
    color: #8aa2ad;
  }

  .searchbar input::placeholder {
    color: #8aa2ad;
  }

  .nav-icons {
    grid-column: 3;
    display: flex;
    gap: 15px;

    a {
      text-decoration: none;
    }
  }
</style>
