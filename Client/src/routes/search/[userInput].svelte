<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { User } from '../../store';
	import { goto } from '$app/navigation';
	let loggedIn;
	User.subscribe((value) => (loggedIn = value.username));
	let suggestion = [];
	onMount(async () => {
		const userInput = $page.params.userInput;
		const response = await fetch(`/search/${userInput}`);
		const data = await response.json();
		console.log(data);
		suggestion = JSON.parse(JSON.stringify(data.usersFound));
	});
	async function addFriend(userName) {
		if (loggedIn) {
			const response = await fetch('http://localhost:8585/add-friend', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					targetUsername: userName
				})
			});
		} else {
			goto('/login');
		}
	}
</script>

<div class="w-96 p-4 my-2 mx-auto lg:w-128">
	<div class="flex flex-col">
		{#each suggestion as suggedtedPeople}
			<section
				class="  py-2 px-3 flex my-2 items-center  justify-between w-full border-2 border-main-bg rounded-md "
			>
				<!--name and username-->
				<section class="flex  items-center w-8/12 ">
					<!--profile img-->
					<a href="/profile/{suggedtedPeople.username}">
						<!-- svelte-ignore a11y-img-redundant-alt -->
						<img
							class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
							src="http://localhost:8585/{suggedtedPeople.profileImg}"
							alt="Current profile photo"
						/>
					</a>
					<!-- Name and username-->
					<a href="/profile/{suggedtedPeople.username}" class="mx-2 w-9/12">
						<h4 class=" font-semibold text-sm text-gray-900 hover:text-gray-500 ">
							{suggedtedPeople.firstName}
							{suggedtedPeople.lastName}
						</h4>
						<h5 class=" text-xs  text-gray-900 hover:text-gray-500">
							@{suggedtedPeople.username}
						</h5>
					</a>
				</section>
				<button
					on:click={addFriend(suggedtedPeople.username)}
					class=" bg-main-bg hover:text-main hover:shadow-xl py-3 px-3 mx-auto rounded-xl  py-1.5 text-white"
					>Add Friend
				</button>
			</section>
		{/each}
	</div>
</div>
