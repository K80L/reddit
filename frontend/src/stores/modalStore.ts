import { writable } from 'svelte/store';

export const showModal = writable(false);

export function closeModal() {
  showModal.set(false);
}

export function openModal() {
  showModal.set(true);
}
