<script>
	import { IsLoggedIn } from '../store';
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import LeftSidebar from '../components/leftSidebar.svelte';
	import { onDestroy } from 'svelte';
	let loggedIn;
	const unsubscribe = IsLoggedIn.subscribe((value) => (loggedIn = value));

	onDestroy(unsubscribe);
</script>

<Navbar />
{#if loggedIn}
	<main class="flex  xl:w-8/12  lg:w-9/12 mx-auto items-start justify-center md:justify-between">
		<LeftSidebar />
		<slot />
	</main>
{:else}
	<slot />
{/if}
