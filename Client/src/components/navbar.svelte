<script>
	import { goto } from '$app/navigation';
	// @ts-nocheck
	import { User } from '../store';
	let username;
	User.subscribe((value) => (username = value.username));
	let searchValue;
	import { page } from '$app/stores';
	const path = $page.url.pathname;
	let humberguer = false;
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
		if (response.status == 200) {
			User.set(null);
			location.replace('/');
		}
	}
</script>

<nav class=" bg-main-bg">
	<div class="h-12 items-center justify-between relative flex lg:w-10/12 md:mx-auto md:w-11/12">
		<a href="/" class="mx-2 text-main">Logo</a>
		<!--search bar-->
		<section class="lg:ml-2">
			<form on:submit|preventDefault={search}>
				<input
					type="text"
					bind:value={searchValue}
					placeholder="search"
					class="w-42  lg:w-96 md:w-72 border-2 rounded-md py-0.5  px-2 focus:outline-hidden focus:outline-none lg:w-128 lg:ml-28  "
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
						class=" {$page.url.pathname === '/dashboard' ? 'text-main' : ''}">Dashboard</a
					>
					<a
						on:click={openHumberguer}
						href="/profile/{username}"
						class={$page.url.pathname === `/profile/${username}` ? 'text-main' : ''}>My Profile</a
					>
					<a
						on:click={openHumberguer}
						href="/notification"
						class={$page.url.pathname === '/notification' ? 'text-main' : ''}>Notification</a
					>
					<a
						href="/messages"
						on:click={openHumberguer}
						class={$page.url.pathname === '/messages' ? 'text-main' : ''}>Messages</a
					>
					<a
						on:click={openHumberguer}
						href="/liked-posts"
						class=" {$page.url.pathname === '/liked-posts' ? 'text-main' : ''}"
					>
						Liked Posts</a
					>
					<a
						href="/setting"
						on:click={openHumberguer}
						class=" {$page.url.pathname === '/setting' ? 'text-main' : ''}">Setting</a
					>
					<button class=" w-fit font-semibold " on:click={openHumberguer} on:click={logout}
						>Logout</button
					>
				</section>
			</div>
			<button
				class=" md:rounded-lg md:border-main md:border-solid md:mx-2 md:border-2 md:px-2 md:py-1  md:text-base md:text-gray-300 md:hover:bg-gray-800 md:hover:shadow-xl hover:text-main hidden md:block"
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
						href="/signup"
						class=" md:rounded-lg md:border-main md:border-solid md:mx-2 md:border-2 md:px-2 md:py-1  md:text-base md:text-gray-300 md:hover:bg-gray-800 md:hover:shadow-xl hover:text-main"
						>Sign up</a
					>
					<a
						href="/login"
						class="md:rounded-lg md:border-main md:border-solid md:mx-2 md:border-2 md:px-2 md:py-1  md:text-base md:text-gray-300 md:hover:bg-gray-800 md:hover:shadow-xl hover:text-main"
						>Login</a
					>
				</section>
			</section>
		{/if}
	</div>
</nav>
