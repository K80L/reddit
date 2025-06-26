<script lang="ts">
  import { createBubbler } from 'svelte/legacy';

  const bubble = createBubbler();
  import Modal from '../modal.svelte';
  import { closeModal } from '../../stores/modalStore';
  import { login } from '../../stores/authStore';

  let username = $state('');
  let password = $state('');
  let error: string | null = null;

  async function handleLogin(event: Event) {
    event.preventDefault();
    error = null;

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
        credentials: 'include',
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      await response.json();
      login({ username });
      closeModal();
    } catch (err) {
      if (err instanceof Error) {
        error = err.message;
      } else {
        error = 'An unknown error occurred';
      }

      console.error(error);
    }
  }
</script>

<Modal className="login__container">
  {#snippet header()}
    <h2 class="login login__title">Log In</h2>
  {/snippet}

  <form class="login login__form" method="POST" onsubmit={handleLogin}>
    <input
      class="login__input"
      autocomplete="off"
      type="text"
      placeholder="Username"
      bind:value={username} />
    <input
      class="login__input"
      autocomplete="off"
      type="password"
      placeholder="Password"
      bind:value={password} />
    <button class="login__btn" type="submit" onchange={bubble('change')}>Log In</button>
  </form>
</Modal>

<style>
  * {
    box-sizing: border-box;
    color: white;
    font-family: system-ui;
    line-height: 1.5;
    font-weight: 300;
  }

  .login,
  .login * {
    color: #000000;
  }

  .login__form {
    display: flex;
    flex-direction: column;
  }

  .login__input[type='text'],
  .login__input[type='password'] {
    background-color: #2c3236;
  }

  .login__input,
  .login__btn {
    margin-top: 1em;
  }

  .login__input {
    border: 5px solid #2c3236;
    outline: none;
    border-radius: 1em;
    color: #ffffff;
  }

  .login__title {
    font-weight: 800;
  }
</style>
