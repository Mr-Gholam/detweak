<script>
	import { goto } from '$app/navigation';
	// @ts-nocheck
	import { User, Notification } from '../store';
	let searchValue;
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	const path = $page.url.pathname;
	let humberguer = false;
	onMount(async () => {
		const response = await fetch('/api/friend-requests');
		const data = await response.json();
		if (data) {
			$Notification = data.length;
		}
	});
	function openHumberguer() {
		const firstBurger = document.getElementById('firstBurger');
		const midtBurger = document.getElementById('midBurger');
		const lastBurger = document.getElementById('lastBurger');
		const menu = document.getElementById('menu');
		const auth = document.getElementById('auth');
		if (humberguer) {
			humberguer = false;
			firstBurger.classList.remove('rotate-45', 'absolute', 'top-5');
			lastBurger.classList.remove('rotate-N45', 'absolute', 'bottom-1');
			midtBurger.classList.remove('hidden');
			if ($User.username) {
				menu.classList.add('hidden');
			} else {
				auth.classList.add('hidden');
			}
		} else {
			humberguer = true;
			firstBurger.classList.add('rotate-45', 'absolute');
			firstBurger.style.top = '15px';
			firstBurger.style.right = '18px';
			lastBurger.classList.add('rotate-N45', 'absolute');
			lastBurger.style.top = '7px';
			lastBurger.style.left = '18px';
			midtBurger.classList.add('hidden');
			if ($User.username) {
				menu.classList.remove('hidden');
			} else {
				auth.classList.remove('hidden');
			}
		}
	}
	async function search() {
		if (searchValue.length > 2) {
			const value = searchValue;
			searchValue = '';
			goto(`/search/${value}`);
		}
	}
	async function logout() {
		const response = await fetch('/api/logout', { method: 'POST' });
		if (response.ok) {
			$User.username = null;
			location.replace('/');
		}
	}
</script>

<nav class=" bg-main-bg border-b border-border sticky top-0 z-10">
	<div
		class="h-16 items-center justify-between relative flex lg:w-10/12 md:mx-auto md:w-11/12 bg-inherit xl:w-8/12 "
	>
		<a href="/" class="mx-2 bg-inherit ">
			<img class=" object-cover" src="../../static/main-logo.png" alt="" />
		</a>
		<!--search bar-->
		<section class="lg:ml-2 bg-inherit">
			<form on:submit|preventDefault={search} class="bg-inherit">
				<input
					type="text"
					bind:value={searchValue}
					placeholder="search"
					class="w-42  md:w-72 border-2 border-border rounded-md py-0.5  px-2 focus:outline-hidden focus:outline-none lg:w-128  bg-inherit xl:mr-40 text-text"
				/>
			</form>
		</section>
		<!--humberguer menu-->
		<div
			class="p-4 space-y-2  rounded  md:hidden  w-14 relative hover:animate-pulse"
			on:click={openHumberguer}
		>
			<span class="block w-5 h-0.5 bg-main transition ease-in-out" id="firstBurger" />
			<span class="block w-5 h-0.5 bg-main transition ease-in-out" id="midBurger" />
			<span class="block w-5 h-0.5 bg-main transition ease-in-out" id="lastBurger" />
		</div>

		{#if $User.username}
			<!--menu-->
			<div class=" p-4   hidden absolute top-10 left-0 bg-main-bg w-full h-fit z-10 " id="menu">
				<section class=" flex justify-evenly flex-col text-lg gap-2 p-2 font-semibold  ">
					<a
						href="/dashboard"
						on:click={openHumberguer}
						class=" {$page.url.pathname === '/dashboard' ? 'text-white' : 'text-text'}">Dashboard</a
					>
					<a
						on:click={openHumberguer}
						href="/profile/{$User.username}"
						class={$page.url.pathname === `/profile/${$User.username}` ? 'text-white' : 'text-text'}
						>My Profile</a
					>
					<a
						on:click={openHumberguer}
						href="/notification"
						class=" flex  justify-between items-center {$page.url.pathname === '/notification'
							? 'text-white'
							: 'text-text'}"
						>Notification
						{#if $Notification}
							<div
								class="bg-main w-5 h-5 rounded-full text-xs flex justify-center items-center text-white mr-1"
							>
								{$Notification}
							</div>
						{/if}</a
					>
					<a
						href="/messages"
						on:click={openHumberguer}
						class={$page.url.pathname === '/messages' ? 'text-white' : 'text-text'}>Messages</a
					>
					<a
						on:click={openHumberguer}
						href="/liked-posts"
						class=" {$page.url.pathname === '/liked-posts' ? 'text-white' : 'text-text'}"
					>
						Liked Posts</a
					>
					<a
						href="/setting"
						on:click={openHumberguer}
						class=" {$page.url.pathname === '/setting' ? 'text-white' : 'text-text'}">Setting</a
					>
					<button
						class=" w-fit font-semibold text-text "
						on:click={openHumberguer}
						on:click={logout}>Logout</button
					>
				</section>
			</div>
			<button
				class=" md:rounded-lg  md:mx-2 m md:px-2 md:py-1  md:text-base    hover:text-main hidden md:block text-text"
				on:click={logout}>Logout</button
			>
		{:else}
			<!--auth -->
			<section
				class="justify-betweenitems-center absolute top-12 left-0 bg-main-bg w-full hidden p-4  md:block md:static  md:p-0 md:w-fit"
				id="auth"
			>
				<section
					class="flex flex-col md:flex-row justify-between md:justify-end gap-2 p-2 text-lg font-semibold md:p-0"
				>
					<a
						href="/login"
						class=" md:rounded-lg  md md:mx-2 border-2 border-main-bg	  md:px-2 md:py-1  md:text-base  hover:text-main text-text  hover:border-main"
						>Login</a
					>
					<a
						href="/signup"
						class="md:rounded-lg md:border-slate-400 md:border-solid md:mx-2 md:border-2 md:px-2 md:py-1  md:text-base text-slate-400   hover:text-main hover:border-main"
						>Sign up</a
					>
				</section>
			</section>
		{/if}
	</div>
</nav>
