<script context="module">
	import { get } from 'svelte/store';

	export const load = async ({ fetch }) => {
		const res = await fetch('/api/friend-requests', {
			credentials: 'include'
		});
		let FriendRequests = [];
		if (res.status == 200) {
			FriendRequests = await res.json();
		}
		let likes = [];
		const newliked = await fetch('/api/new-liked');
		if (newliked.status == 200) {
			const data = await newliked.json();
			likes = JSON.parse(JSON.stringify(data.Likes));
			likes = likes.reverse();
			Notification.set(get(Notification) - data.newlikes);
		}
		return { props: { FriendRequests, likes } };
	};
</script>

<script>
	// @ts-nocheck
	import formatDistanceToNowStrict from 'date-fns/formatDistanceToNowStrict/index';
	import { loading, Notification, ws } from '../store';
	export let FriendRequests;
	export let likes;
	let friendRequests = [...FriendRequests];
	if ($ws) {
		$ws.onmessage = (e) => {
			const { notification } = JSON.parse(e.data);
			if (notification.NewFriendReq) {
				$Notification++;
				friendRequests.push(notification.NewFriendReq);
				friendRequests.sort((a, b) => {
					return a.RequestId > b.RequestId;
				});
				friendRequests = friendRequests;
			}
			if (notification.Liked) {
				likes.unshift(notification.Liked);
				likes = likes;
			}
			if (notification.Disliked) {
				likes = likes.filter((liked) => {
					return (
						liked.Username != notification.Disliked.Username &&
						liked.PostId != notification.Disliked.PostId
					);
				});
				likes = likes;
			}
		};
	}
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
			$Notification = $Notification - 1;
			$Notification = $Notification;
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
			$Notification = $Notification - 1;
			$Notification = $Notification;
		}
	}
</script>

<svelte:head>
	<title>Notification</title>
</svelte:head>
<div class="w-88 lg:w-128 my-2 py-4 flex flex-col">
	<!--Friend requests-->
	<div class="w-full">
		<h4 class="text-center font-semibold my-2 text-text">Friend Requests</h4>
		<section class="flex justify-between gap-4 flex-col">
			{#if friendRequests}
				{#each friendRequests as friendRequest}
					<section
						class="  py-2 px-3 flex my-2 items-center  justify-between 2-full border-2 border-border rounded-md "
						id={friendRequest.RequestId}
					>
						<!--name and username-->
						<section class="flex  items-center w-5/12  justify-between">
							<!--profile img-->
							<a href="/profile/{friendRequest.Username}">
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
							<a href="/profile/{friendRequest.username}" class="mx-2 w-fit">
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
	</div>
	<!-- New likes -->
	<div class="w-full ">
		<h4 class="text-center font-semibold my-2 text-text">Likes</h4>
		<section class="flex justify-between gap-4 flex-col">
			{#each likes as like}
				<section
					class="  py-2 px-3 flex my-2 items-center  justify-between w-full lg:border-2  border-border rounded-md h-24 "
					id={like.RequestId}
				>
					<!--name and username-->
					<section class="flex  items-center h-16">
						<!--profile img-->
						<a href="/profile/{like.Username}">
							<!-- svelte-ignore a11y-img-redundant-alt -->
							{#if like.ImgUrl}
								<img
									class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
									src="/api/images/{like.ImgUrl}"
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
					</section>
					<section class="w-48 lg:w-64 flex  justify-between flex-col  gap-2">
						<div class="text-text">
							<div class="flex items-center w-fit">
								<a href="/profile/{like.Username} ">
									<p class=" text-sm text-white hover:text-text-hover font-semibold ">
										{like.Username}
									</p>
								</a>
								<span class="text-xs mx-2"> liked </span>
								<a
									href="/post/{like.PostId}"
									class="text-sm font-semibold text-white hover:text-text-hover"
								>
									your post</a
								>
							</div>
						</div>
						<p class="text-xs  bg-inherit text-text text-right">
							{formatDistanceToNowStrict(new Date(like.CreatedAt), { addSuffix: true })}
						</p>
					</section>
					<section class=" h-16 w-16">
						{#if like.PostImgUrl}
							<a href="/post/{like.PostId}">
								<img
									src="/api/images/{like.PostImgUrl}"
									alt=""
									class=" h-full w-full object-cover"
								/>
							</a>
						{/if}
					</section>
				</section>
			{/each}
		</section>
	</div>
</div>
