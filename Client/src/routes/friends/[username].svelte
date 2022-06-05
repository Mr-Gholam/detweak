<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { loading, User, ws } from '../../store';
	import { goto } from '$app/navigation';
	let friends = [];
	let username;
	onMount(async () => {
		username = $page.params.username;
		const response = await fetch(`/api/friends/${username}`);
		const data = await response.json();
		data.sort((a, b) => {
			if (a.IsFriend == 'Friend') {
				return -1;
			} else if (a.IsFriend == 'Not Friend') {
				return 1;
			} else {
				return 0;
			}
		});
		friends = data;
	});
	async function addFriend(username) {
		if ($User) {
			const response = await fetch('/api/add-friend', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username
				})
			});
			const data = await response.json();
			if (response.ok) {
				const user = friends.find((user) => user.Username == username);
				user.IsFriend = data.status;
				friends = friends;
				$ws.send(JSON.stringify({ Target: username }));
			}
		} else {
			goto('/login');
		}
	}
	function openChat(username) {
		goto(`/messages?username=${username}`);
	}
</script>

<svelte:head>
	<title>{username}'s Friends</title>
</svelte:head>

<div class="  md:w-fit w-full lg:w-128 my-2">
	<h1 class="text-white text-center text-lg py-4">{username}'s Friends</h1>
	{#each friends as friend}
		<section
			class="  py-2 px-3 flex my-2 items-center  justify-between w-full border-2 border-border rounded-md "
		>
			<!--name and username-->
			<section class="flex  items-center w-8/12 ">
				<!--profile img-->
				<a href="/profile/{friend.Username}">
					<!-- svelte-ignore a11y-img-redundant-alt -->
					{#if friend.ImgUrl}
						<img
							class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
							src="/api/images/{friend.ImgUrl}"
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
				<a href="/profile/{friend.Username}" class="mx-2 w-9/12">
					<h4 class=" font-semibold text-sm text-text hover:text-text-hover ">
						{friend.Firstname}
						{friend.Lastname}
					</h4>
					<h5 class=" text-xs  text-text hover:text-text-hover">
						@{friend.Username}
					</h5>
				</a>
			</section>
			{#if friend.Username != $User.username}
				{#if friend.IsFriend == 'Pending'}
					<button class="main-btn ">
						<h5 class="mx-2	 text-sm">Friend Request Sent</h5>
					</button>
				{:else if friend.IsFriend == 'Friend'}
					<button class="main-btn" on:click={openChat(friend.Username)}>Sent Massage</button>
				{:else}
					<button on:click={addFriend(friend.Username)} class=" main-btn">Add Friend </button>
				{/if}
			{/if}
		</section>
	{/each}
</div>
