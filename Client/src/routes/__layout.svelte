<script>
	import { User } from '../store';
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import LeftSidebar from '../components/leftSidebar.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	onMount(async () => {
		const res = await fetch('/jwt');
		const u = await res.json();
		if (res.status == 200) {
			User.set(u);
		}
	});
</script>

<Navbar />
{#if $User.username}
	<main
		class="md:flex  xl:w-8/12  lg:w-9/12 md:mx-auto md:items-start md:justify-center md:justify-between "
	>
		<LeftSidebar />
		<slot />
	</main>
{:else}
	<slot />
{/if}
