<script>
	// @ts-nocheck

	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { onlineFriends, User } from '../store';
	let suggestion = [
		{
			userName: 'mr-gholam',
			firstName: 'Mehdi',
			lastName: 'Gholami',
			onlineTime: '1-2 pm'
		},
		{
			userName: 'jafar',
			firstName: 'jafar',
			lastName: 'Gholami',
			onlineTime: '3-4 am'
		}
	];
	onMount(async () => {
		const response = await fetch('/api/online-friends');
		const data = await response.json();
		$onlineFriends = data;
		$onlineFriends = $onlineFriends;
	});
	function openChat(username) {
		goto(`/messages?username=${username}`);
	}
</script>

<!--Right part-->
<div
	class=" h-fit p-4 my-2  xl:max-w-3/12 {($User.username && $page.url.pathname === '/dashboard') ||
	$page.url.pathname === '/liked-posts' ||
	$page.params.username ||
	$page.params.postId ||
	$page.params.jobId ||
	$page.url.pathname === '/notification' ||
	$page.url.pathname === '/my-jobs' ||
	$page.url.pathname === '/my-offers' ||
	$page.url.pathname === '/find-job'
		? '2xl:block'
		: 'hidden'}"
>
	<!--online friends-->
	{#if $onlineFriends}
		<section>
			<h3 class="text-center text-lg font-semibold my-2 text-white ">Online Friends</h3>
			<section class="flex justify-between gap-4 flex-col">
				{#each $onlineFriends as onlinefriend}
					<section
						class="  py-2 px-3 flex my-2 items-center  justify-between xl:w-72 border-2 border-border rounded-md "
					>
						<!--name and username-->
						<section class="flex  items-center w-full ">
							<!--profile img-->
							<a href="/prfile/{onlinefriend.Username}" class="lg:w-6/12">
								<!-- svelte-ignore a11y-img-redundant-alt -->
								{#if onlinefriend.ImgUrl}
									<img
										class="h-12 w-12 object-cover rounded-full lg:mr-16 "
										src="/api/images/{onlinefriend.ImgUrl}"
										alt="Current profile photo"
									/>
								{:else}
									<div
										class="h-12 w-12 rounded-full  bg-main-bg flex items-center justify-center border-border border"
									>
										<i class="fa-solid fa-user text-slate-400 text-xl" />
									</div>
								{/if}
							</a>
							<!-- Name and username-->
							<a href="/profile/{onlinefriend.Username}" class="mx-2 w-full">
								<h4 class=" font-semibold text-sm text-text hover:text-text-hover ">
									{onlinefriend.Firstname}
									{onlinefriend.Lastname}
								</h4>
								<h5 class=" text-xs  text-text hover:text-text-hover  ">
									@{onlinefriend.Username}
								</h5>
							</a>
							<button
								class="main-btn  px-3 py-1.5 text-xs"
								on:click={openChat(onlinefriend.Username)}>Send message</button
							>
						</section>
					</section>
				{/each}
			</section>
		</section>
	{/if}
	<!--Suggestion-->
	<section class="my-4">
		<h3 class="text-center text-lg font-semibold my-2 text-white">People you might know</h3>
		<!--suggestion box-->
		<section class="flex justify-between gap-4 flex-col">
			{#each suggestion as suggedtedPeople}
				<section
					class="  py-2 px-3 flex my-2 items-center  justify-between lg:w-72 border-2 border-border rounded-md "
				>
					<!--name and username-->
					<section class="flex  items-center w-8/12 ">
						<!--profile img-->
						<a href="/profile/{suggedtedPeople.userName}" class="lg:w-6/12">
							<!-- svelte-ignore a11y-img-redundant-alt -->
							<img
								class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
								src="data:image/jpeg;base64,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"
								alt="Current profile photo"
							/>
						</a>
						<!-- Name and username-->
						<a href="/profile/{suggedtedPeople.userName}" class="mx-2 w-full">
							<h4 class=" font-semibold text-sm text-text hover:text-text-hover ">
								{suggedtedPeople.firstName}
								{suggedtedPeople.lastName}
							</h4>
							<h5 class=" text-xs  text-text hover:text-text-hover   ">
								@{suggedtedPeople.userName}
							</h5>
						</a>
					</section>
					<form
						action="/friend-request"
						method="post"
						class="text-xs font-semibold   text-center  w-4/12 "
					>
						<button class=" main-btn   px-3  py-1.5 ">Add Friend </button>
					</form>
				</section>
			{/each}
		</section>
	</section>
</div>
