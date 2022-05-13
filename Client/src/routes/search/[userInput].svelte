<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { loading, User } from '../../store';
	import { goto } from '$app/navigation';
	let loggedIn;
	User.subscribe((value) => (loggedIn = value.username));
	let suggestion = [];
	onMount(async () => {
		const userInput = $page.params.userInput;
		const response = await fetch(`/api/search/${userInput}`);
		const data = await response.json();
		suggestion = JSON.parse(JSON.stringify(data.usersFound));
		$loading = false;
	});
	async function addFriend(username) {
		if (loggedIn) {
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

<div class="w-96 py-4 my-2 lg:ml-32 lg:mx-auto lg:w-128">
	<div class="flex flex-col">
		{#each suggestion as suggedtedPeople}
			<section
				class="  py-2 px-3 flex my-2 items-center  justify-between w-full border-2 border-border rounded-md "
			>
				<!--name and username-->
				<section class="flex  items-center w-8/12 ">
					<!--profile img-->
					<a href="/profile/{suggedtedPeople.username}">
						<!-- svelte-ignore a11y-img-redundant-alt -->
						{#if suggedtedPeople.profileImg}
							<img
								class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
								src="/api/{suggedtedPeople.profileImg}"
								alt="Current profile photo"
							/>
						{:else}
							<div
								class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center"
							>
								<i class="fa-solid fa-user text-slate-400 text-2xl" />
							</div>
						{/if}
					</a>
					<!-- Name and username-->
					<a href="/profile/{suggedtedPeople.username}" class="mx-2 w-9/12">
						<h4 class=" font-semibold text-sm text-text hover:text-text-hover ">
							{suggedtedPeople.firstName}
							{suggedtedPeople.lastName}
						</h4>
						<h5 class=" text-xs  text-text hover:text-text-hover">
							@{suggedtedPeople.username}
						</h5>
					</a>
				</section>
				{#if !suggedtedPeople.myProfile}
					{#if suggedtedPeople.sentRequest}
						<button class="main-btn ">
							<h5 class="mx-2	 text-sm">Friend Request Sent</h5>
						</button>
					{:else if suggedtedPeople.isFriend}
						<button class="main-btn">Sent Massage</button>
					{:else}
						<button on:click={addFriend(suggedtedPeople.username)} class=" main-btn"
							>Add Friend
						</button>
					{/if}
				{/if}
			</section>
		{/each}
	</div>
</div>
