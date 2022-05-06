<script>
	import { onMount } from 'svelte';

	let friendRequests = [];
	onMount(async () => {
		const response = await fetch('/api/friend-requests');
		const data = await response.json();
		const orderedReq = JSON.parse(JSON.stringify(data.friendRequests));
		friendRequests = orderedReq.reverse();
	});
	async function acceptReq(requestId) {
		const response = await fetch('/api/accept-request', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				requestId
			})
		});
		if (response.status == 200) {
			const el = document.getElementById(`${requestId}`);
			el.parentNode.removeChild(el);
		}
	}
</script>

<svelte:head>
	<title>Notification</title>
</svelte:head>
<div class="w-96 lg:w-9/12 my-2 p-4">
	<!--Friend requests-->
	<section class="w-fit">
		<h4 class="text-center font-semibold my-2">Friend Requests</h4>
		<section class="flex justify-between gap-4 flex-col">
			{#each friendRequests as friendRequest}
				<section
					class="  py-2 px-3 flex my-2 items-center  justify-between lg:w-72 border-2 border-main-bg rounded-md "
					id={friendRequest.requestId}
				>
					<!--name and username-->
					<section class="flex  items-center w-8/12 ">
						<!--profile img-->
						<a href="/profile/{friendRequest.username}" class="lg:w-6/12">
							<!-- svelte-ignore a11y-img-redundant-alt -->
							<img
								class="h-12 lg:w-12 w-16 object-cover rounded-full hover:opacity-90  "
								src="/api/{friendRequest.profileImg}"
								alt="Current profile photo"
							/>
						</a>
						<!-- Name and username-->
						<a href="/profile/{friendRequest.username}" class="mx-2 w-full">
							<h4 class=" font-semibold text-sm text-gray-900 hover:text-gray-500 ">
								{friendRequest.firstName}
								{friendRequest.lastName}
							</h4>
							<h5 class=" text-xs  text-gray-900 hover:text-gray-500  ">
								@{friendRequest.username}
							</h5>
						</a>
					</section>
					<section class="flex justify-between w-10  items-center gap-3 text-lg ">
						<input type="hidden" value={friendRequest.requestId} id="requestId" />
						<button
							class=" w-7 h-7 rounded-full text-center  hover:text-green"
							on:click={acceptReq(friendRequest.requestId)}
						>
							<i class="fa-solid fa-check" />
						</button>
						<button class=" w-6 h-6 rounded-full text-center  hover:text-error">
							<i class="fa-solid fa-x" />
						</button>
					</section>
				</section>
			{/each}
		</section>
	</section>
</div>
