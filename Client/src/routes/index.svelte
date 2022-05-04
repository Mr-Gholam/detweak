<script>
	import { goto } from '$app/navigation';

	import { onMount } from 'svelte';

	onMount(async () => {
		const response = await fetch('/api/', { method: 'POST' });
		if (response.status == 200) {
			goto('/dashboard');
		}
	});
</script>

<h1 class="bg-red-900">Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentat</p>
