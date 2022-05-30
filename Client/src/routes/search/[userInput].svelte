<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { loading, User } from '../../store';
	import { goto } from '$app/navigation';
	let suggestion = [];
	let input;
	onMount(async () => {
		const userInput = $page.params.userInput;
		input = $page.params.userInput;
		const response = await fetch(`/api/search/${userInput}`);
		const data = await response.json();
		console.log(data);
		suggestion = JSON.parse(JSON.stringify(data));
		$loading = false;
	});
	async function addFriend(username) {
		if ($User) {
			const response = await fetch('/api/add-friend', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					targetUsername: username
				})
			});
			if (response.status == 200) {
				const userInput = $page.params.userInput;
				const response = await fetch(`/search/${userInput}`);
				const data = await response.json();
				suggestion = JSON.parse(JSON.stringify(data.usersFound));
			}
		} else {
			goto('/login');
		}
	}
</script>

<svelte:head>
	<title>Search</title>
</svelte:head>

<div class="w-96 py-4 my-2  lg:mx-auto lg:w-128 {$User.username ? 'lg:ml-32' : 'lg:ml-72'}">
	<div class="flex flex-col">
		{#if suggestion}
			<h1 class="text-text text-center">Users found by : {input}</h1>
			{#each suggestion as suggedtedPeople}
				<section
					class="  py-2 px-3 flex my-2 items-center  justify-between w-full border-2 border-border rounded-md "
				>
					<!--name and username-->
					<section class="flex  items-center w-8/12 ">
						<!--profile img-->
						<a href="/profile/{suggedtedPeople.Username}">
							<!-- svelte-ignore a11y-img-redundant-alt -->
							{#if suggedtedPeople.ImgUrl}
								<img
									class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
									src="/api/images/{suggedtedPeople.ImgUrl}"
									alt="Current profile photo"
								/>
							{:else}
								<div
									class="h-12 w-12 rounded-full hover:opacity-90  flex items-center justify-center border-border border-2"
								>
									<i class="fa-solid fa-user text-slate-400 text-2xl" />
								</div>
							{/if}
						</a>
						<!-- Name and username-->
						<a href="/profile/{suggedtedPeople.Username}" class="mx-2 w-9/12">
							<h4 class=" font-semibold text-sm text-text hover:text-text-hover ">
								{suggedtedPeople.Firstname}
								{suggedtedPeople.Lastname}
							</h4>
							<h5 class=" text-xs  text-text hover:text-text-hover">
								@{suggedtedPeople.Username}
							</h5>
						</a>
					</section>
					{#if suggedtedPeople.Username != $User.username}
						{#if suggedtedPeople.IsFriend == 'Pending'}
							<button class="main-btn ">
								<h5 class="mx-2	 text-sm">Friend Request Sent</h5>
							</button>
						{:else if suggedtedPeople.IsFriend == 'Friend'}
							<button class="main-btn">Sent Massage</button>
						{:else}
							<button on:click={addFriend(suggedtedPeople.Username)} class=" main-btn"
								>Add Friend
							</button>
						{/if}
					{/if}
				</section>
			{/each}
		{/if}
		{#if !suggestion}
			<h1 class="text-text text-center">Users not found</h1>
		{/if}
	</div>
</div>
