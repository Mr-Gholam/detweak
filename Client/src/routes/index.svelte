<script>
	import { goto } from '$app/navigation';
	import { User } from '../store';

	import { onMount } from 'svelte';

	onMount(async () => {
		if ($User) {
			goto('/dashboard');
		}
	});
</script>

<svelte:head>
	<title>Detweak</title>
</svelte:head>
<h1 class="bg-red-900">Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentat</p>
