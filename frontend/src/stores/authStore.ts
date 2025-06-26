import { writable } from 'svelte/store';
import { customFetch } from '../utils/fetch';

export const isLoading = writable(true);
export const isLoggedIn = writable(false);
export const user = writable<{ username: string } | null>(null);

export function login(userData: { username: string } | null) {
  isLoggedIn.set(true);
  isLoading.set(false);
  user.set(userData);
}

export function logout() {
  isLoggedIn.set(false);
  isLoading.set(false);
  user.set(null);
}

export async function checkLoginStatus() {
  try {
    const response = await customFetch('auth/status');
    console.log('Response:', response);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    if (response.status === 401) {
      isLoggedIn.set(false);
      isLoading.set(false);
      user.set(null);
      console.log('User is not logged in');
      return;
    }
    if (response.status === 200) {
      isLoggedIn.set(true);
      isLoading.set(false);
      user.set(response.data.user);
      console.log('User is logged in');
      return;
    }
  } catch (error) {
    isLoggedIn.set(false);
    isLoading.set(false);
    console.error('Error checking login status:', error);
  }
}
