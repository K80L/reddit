export async function customFetch(path?: string, options?: RequestInit) {
  try {
    const response = await fetch(`http://localhost:8080/${path}`, {
      method: 'GET',
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
      credentials: 'include',
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    return await response.json();
  } catch (error) {
    console.error('Error:', error);
  }
}
