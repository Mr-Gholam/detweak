<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import { loading } from '../store';
	let friendRequests = [];
	onMount(async () => {
		const response = await fetch('/api/friend-requests');
		const data = await response.json();
		friendRequests = data;
		$loading = false;
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
	async function rejectReq(requestId) {
		const response = await fetch('/api/reject-request', {
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
		<h4 class="text-center font-semibold my-2 text-text">Friend Requests</h4>
		<section class="flex justify-between gap-4 flex-col">
			{#if friendRequests}
				{#each friendRequests as friendRequest}
					<section
						class="  py-2 px-3 flex my-2 items-center  justify-between lg:w-72 border-2 border-border rounded-md "
						id={friendRequest.RequestId}
					>
						<!--name and username-->
						<section class="flex  items-center w-8/12 ">
							<!--profile img-->
							<a href="/profile/{friendRequest.Username}" class="lg:w-6/12">
								<!-- svelte-ignore a11y-img-redundant-alt -->
								{#if friendRequest.ImgUrl}
									<img
										class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
										src="/api/images/{friendRequest.ImgUrl}"
										alt="Profile photo"
									/>
								{:else}
									<div
										class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center border-2 border-border "
									>
										<i class="fa-solid fa-user text-slate-400 text-2xl" />
									</div>
								{/if}
							</a>
							<!-- Name and username-->
							<a href="/profile/{friendRequest.username}" class="mx-2 w-full">
								<h4 class=" font-semibold text-sm text-text hover:text-text-hover">
									{friendRequest.Firstname}
									{friendRequest.Lastname}
								</h4>
								<h5 class=" text-xs text-text hover:text-text-hover ">
									@{friendRequest.Username}
								</h5>
							</a>
						</section>
						<section class="flex justify-between w-10  items-center gap-3 text-lg ">
							<input type="hidden" value={friendRequest.RequestId} id="requestId" />
							<button
								class=" w-7 h-7 rounded-full text-center text-text mt-1 hover:text-green"
								on:click={acceptReq(friendRequest.RequestId)}
							>
								<i class="fa-solid fa-check" />
							</button>
							<button
								on:click={rejectReq(friendRequest.RequestId)}
								class=" w-6 h-6 rounded-full text-center text-text  hover:text-error"
							>
								<i class="fa-solid fa-x" />
							</button>
						</section>
					</section>
				{/each}
			{/if}
		</section>
	</section>
</div>
