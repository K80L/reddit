<script lang="ts">
  import { run, self, createBubbler, stopPropagation } from 'svelte/legacy';

  const bubble = createBubbler();
  import { closeModal, showModal } from '../stores/modalStore';

  let dialog: HTMLDialogElement = $state();
  interface Props {
    className?: string;
    header?: import('svelte').Snippet;
    children?: import('svelte').Snippet;
  }

  let { className = '', header, children }: Props = $props();

  run(() => {
    if (dialog && $showModal) {
      dialog.showModal();
    } else if (dialog && !$showModal) {
      dialog.close();
    }
  });

  // Adding console logs to debug the state and dialog element
  run(() => {
    if (dialog && $showModal) dialog.showModal();
  });
</script>

<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_noninteractive_element_interactions -->
<dialog bind:this={dialog} onclose={closeModal} onclick={self(() => dialog.close())}>
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div class={className} onclick={stopPropagation(bubble('click'))}>
    {@render header?.()}
    <hr />
    {@render children?.()}
  </div>
</dialog>

<style>
  dialog {
    max-width: 32em;
    border-radius: 0.2em;
    border: none;
    padding: 0;
    width: 528px;
  }

  dialog * {
    color: #000000;
  }

  dialog::backdrop {
    background: rgba(0, 0, 0, 0.3);
  }

  dialog > div {
    padding: 2em;
  }

  dialog[open] {
    animation: zoom 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
  }

  .login__container {
    background-color: #171c1f;
  }

  @keyframes zoom {
    from {
      transform: scale(0.95);
    }
    to {
      transform: scale(1);
    }
  }

  dialog[open]::backdrop {
    animation: fade 0.2s ease-out;
  }

  @keyframes fade {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
</style>
