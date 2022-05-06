<script>
	import { loading, User } from '../store';
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import LeftSidebar from '../components/leftSidebar.svelte';
	import Loading from '../components/loading.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { navigating } from '$app/stores';
	$: $loading = !!$navigating;
	onMount(async () => {
		const res = await fetch('/api/jwt');
		const u = await res.json();
		if (res.status == 200) {
			User.set(u);
		}
	});
</script>

<Navbar />
{#if $User.username}
	{#if $loading}
		<main
			class="md:flex  xl:w-8/12  lg:w-9/12 md:mx-auto md:items-start md:justify-center md:justify-between {$loading
				? 'blur-sm w-full'
				: ''}"
		>
			<LeftSidebar />
			<slot />
		</main>
		<Loading />
	{:else}
		<main
			class="md:flex  xl:w-8/12  lg:w-9/12 md:mx-auto md:items-start md:justify-center md:justify-between "
		>
			<LeftSidebar />
			<slot />
		</main>
	{/if}
{:else}
	<div class={$loading ? 'blur-sm ' : ''}>
		<slot />
	</div>
	<Loading />
{/if}
