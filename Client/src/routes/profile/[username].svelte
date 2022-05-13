<script>
	// @ts-nocheck

	import { page } from '$app/stores';
	import { loading, User } from '../../store';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	let user;
	User.subscribe((value) => (user = value));
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	let comment;
	let Loading = true;
	let firstName;
	let month;
	let day;
	let lastName;
	let bio;
	let profileImg;
	let birthday;
	let onlineTime;
	let userName;
	let postCount;
	let friendCount;
	let location;
	let myProfile;
	let isFriend;
	let sentRequest;
	let posts = null;
	onMount(async () => {
		const username = $page.params.username;
		if (username === user.username) {
			myProfile = true;
		}
		const response = await fetch(`/api/profile/${username}`);
		const data = await response.json();
		if (data) {
			$loading = false;
		}
		firstName = data.firstName;
		lastName = data.lastName;
		userName = data.username;
		bio = data.bio;
		profileImg = data.profileImg;
		birthday = data.birthday;
		postCount = data.postCount;
		friendCount = data.friendCount;
		location = data.location;
		isFriend = data.isFriend;
		sentRequest = data.sentRequest;
		if (birthday) {
			day = new Date(birthday).getDate();
			month = new Date(birthday).getMonth() + 1;
		}
		onlineTime = data.onlineTime;
		posts = data.availablePosts;
	});
	async function addFriend() {
		if (user) {
			const response = await fetch('/api/add-friend', {
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
	// like post
	async function likePost(postId) {
		const likeBtn = document.getElementById(`like-${postId}`);
		const likeNumber = document.getElementById(`like-n-${postId}`);
		const response = await fetch('/api/like-post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				postId
			})
		});
		const data = await response.json();
		if (data.added) {
			likeBtn.classList.add('text-red-600');
			likeNumber.classList.add('text-red-600');
			likeBtn.classList.remove('text-gray-400');
			likeNumber.classList.remove('text-gray-400');
			likeNumber.innerText = Number(likeNumber.innerText) + 1;
		} else {
			likeBtn.classList.remove('text-red-600');
			likeNumber.classList.remove('text-red-600');
			likeBtn.classList.add('text-gray-400');
			likeNumber.classList.add('text-gray-400');
			likeNumber.innerText = Number(likeNumber.innerText) - 1;
		}
	}
	//  open post option
	function postOption(postId) {
		const option = document.getElementById(`${postId}`);
		if (option.classList.contains('hidden')) {
			option.classList.remove('hidden');
		} else {
			option.classList.add('hidden');
		}
	}
	// add Comment
	async function addComment(postId) {
		const commentIcon = document.getElementById(`comment-${postId}`);
		const postBtn = document.getElementById(`postBtn-${postId}`);
		const sent = document.createElement('i');
		sent.classList.add('fa-solid', 'fa-check', 'mr-1', 'text-sm');
		const Sent = document.createElement('p');
		Sent.classList.add('text-sm');
		Sent.innerText = 'Sent';
		const response = await fetch('/api/add-comment', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				comment,
				postId
			})
		});
		if (response.status == 200) {
			commentIcon.classList.add('pulse-icon');
			postBtn.classList.add('flex', 'items-center', 'border-green', 'text-green');
			postBtn.innerText = '';
			postBtn.appendChild(sent);
			postBtn.appendChild(Sent);
			setTimeout(() => {
				postBtn.removeChild(sent);
				postBtn.removeChild(Sent);
				postBtn.classList.remove('flex', 'items-center', 'border-green', 'text-green');
				postBtn.innerText = 'Post';
			}, 1400);
			comment = '';
		}
	}
	// open post page
	function openPost(postId) {
		goto(`/post/${postId}`);
	}
	async function deletePost(postId) {
		const main = document.getElementById('main');
		const postbody = document.getElementById(`body-${postId}`);
		const response = await fetch('/api/delete-post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				postId
			})
		});
		if (response.status == 200) {
			postOption(postId);
			postbody.classList.add('right-exit');
			setTimeout(() => {
				main.removeChild(postbody);
			}, 1000);
		}
	}
	// Open Edit post
	function openEdit(postId) {
		const des = document.getElementById(`des-${postId}`);
		const form = document.getElementById(`form-${postId}`);
		if (form.classList.contains('hidden')) {
			postOption(postId);
			form.classList.remove('hidden');
			des.classList.add('hidden');
		} else {
			des.classList.remove('hidden');
			form.classList.add('hidden');
		}
	}
	async function updatePost(postId) {
		const updatedDes = document.getElementById(`upDes-${postId}`);
		const res = await fetch('/api/update-post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				updatedDes: updatedDes.value,
				postId
			})
		});
		if (res.status == 200) {
			openEdit(postId);
			const response = await fetch('/api/availablePosts');
			const Posts = await response.json();
			posts = Posts.availablePosts;
		}
	}
	function openChat(username) {
		goto(`/messages?username=${username}`);
	}
</script>

<svelte:head>
	<title>{userName ? userName : 'profile'}</title>
</svelte:head>
<div class="flex md:w-9/12 items-start justify-center md:justify-between w-full">
	<!--Main part-->
	<div
		id="main"
		class="flex  justify-between items-center md:py-4  gap-4 flex-col  w-96 lg:w-128 md:mr-32 lg:mr-0 xl:max-w-9/12 overflow-x-hidden"
	>
		<div
			class="flex  justify-between items-center md:py-4  gap-4 flex-col  w-96 lg:w-128 md:mr-32 lg:mr-0"
		>
			<section
				class="w-full flex flex-col h-fit my-2  shadow-lg border-solid border-slate-500/30 md:border  rounded-md"
			>
				<!-- svelte-ignore a11y-img-redundant-alt -->
				<!-- profile img - first name - last name - online time -->
				<div class="flex items-center w-full justify-evenly border-b border-slate-500/30 p-2">
					{#if profileImg}
						<img
							class="h-20 w-20 object-cover rounded-full lg:mr-16 "
							src="/api/{profileImg}"
							alt="Current profile photo"
						/>
					{:else}
						<div class="h-20 w-20 rounded-full  bg-main-bg flex items-center justify-center">
							<i class="fa-solid fa-user text-slate-400 text-4xl" />
						</div>
					{/if}
					<div>
						<h4 class=" font-semibold text-base text-gray-900 my-1 lg:mr-16">
							{firstName}
							{lastName}
						</h4>
						<h5 class=" text-xs  text-gray-900 font-semibold  my-1  ">
							@{userName}
						</h5>
					</div>
					{#if onlineTime}
						<h4 class="text-green text-sm" id="timerCountdown">
							<i class="fa-solid fa-signal" />
							{onlineTime}
						</h4>
					{/if}
				</div>
				<!-- birthday - location - bio-->
				<div class="m-2 p-2">
					<!--birthday - location -->
					<section class="flex w-6/12 justify-between lg:w-7/12">
						{#if birthday}
							<h5 class=" text-xs  text-gray-900 font-semibold  my-1">
								<i class="fa-solid fa-cake-candles" />
								{month}/{day}
							</h5>
						{/if}
						{#if location}
							<h5 class=" text-xs  text-gray-900 font-semibold  my-1">
								<i class="fa-solid fa-location-dot" />
								{location}
							</h5>
						{/if}
					</section>
					{#if bio}
						<h5 class="text-sm my-2 ">
							{bio}
						</h5>
					{/if}
				</div>
				<!-- posts - friends - friend request or send massage-->
				<div class="flex justify-evenly w-full items-center p-2">
					<section class="flex flex-col items-center">
						<h4>{postCount}</h4>
						<h3 class="font-semibold">Posts</h3>
					</section>
					<section class="flex flex-col items-center">
						<h4>{friendCount}</h4>
						<h3 class="font-semibold">Friends</h3>
					</section>
					{#if !myProfile}
						{#if sentRequest}
							<button
								class="px-2 py-1 border border-main-bg rounded-md   bg-main-bg text-white hover:text-main flex items-center "
							>
								<h5 class="mx-2	 text-sm">Friend Request Sent</h5>
							</button>
						{:else if isFriend}
							<button
								class="px-2 py-1 border border-main-bg rounded-md   hover:bg-main-bg hover:text-white text-main-bg"
								on:click={openChat(userName)}>Sent Massage</button
							>
						{:else}
							<button
								class="px-2 py-1 border border-main-bg rounded-md   hover:bg-main-bg hover:text-white text-main-bg"
								on:click={addFriend}>Add Friend</button
							>
						{/if}
					{/if}
				</div>
			</section>
			<!--post-->
			{#if posts}
				{#each posts as post}
					<!-- post outline-->
					<div
						id="body-{post.postId}"
						class="md:border-2 border-solid border-gray-200  shadow-xl w-full rounded-md my-2 overflow-x-hidden"
					>
						<!-- svelte-ignore a11y-img-redundant-alt -->
						<section class=" flex justify-between items-center flex-col ">
							<!--post info-->
							<section
								class="  p-2 flex  items-center gap-2 justify-between w-full border-b border-solid border-gray-200"
							>
								<!--name and username-->
								<section class="flex gap-2 items-center">
									<!--profile img-->
									<a href="/profile/{post.username}" class="ml-2">
										{#if post.profileImg}
											<img
												class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
												src="/api/{post.profileImg}"
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
									<a href="/profile/{post.username}">
										<h4 class="mx-2 font-semibold text-gray-900 hover:text-gray-500">
											{post.firstName}
											{post.lastName}
										</h4>
										<h5 class=" text-sm  text-gray-900 hover:text-gray-500 mx-2 ">
											@{post.username}
										</h5>
									</a>
								</section>
								<section class="flex items-center ">
									<h6 class="text-xs text-orange mx-2 cursor-default">
										{post.onlineTime}
									</h6>
									{#if post.username == user.username}
										<div class="relative">
											<i
												on:click={postOption(post.postId)}
												class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main-bg"
											/>
											<div
												id={post.postId}
												class="hidden absolute bg-main-bg w-28 flex flex-col items-center  rounded p-3  option gap-2 z-10"
											>
												<button
													on:click={openEdit(post.postId)}
													class="text-sm flex items-center w-full justify-start hover:text-main text-white"
												>
													<i class="fa-solid fa-pen-to-square mr-1 text-xs" /> Edit Post
												</button>
												<button
													on:click={deletePost(post.postId)}
													class="text-sm flex items-center w-full justify-start hover:text-red-600 text-white"
												>
													<i class="fa-solid fa-trash mr-1 text-xs" /> Delete Post
												</button>
											</div>
										</div>
									{/if}
								</section>
							</section>
							<!-- post-->
							<section class="w-full h-fit">
								{#if post.postImg}
									<img
										on:dblclick={likePost(post.postId)}
										class="w-full h-fit  md:mx-auto  object-cover"
										src="/api/{post.postImg}"
										alt=""
									/>
								{/if}
								<h3
									class="text-base mx-2 my-4"
									on:dblclick={likePost(post.postId)}
									id="des-{post.postId}"
								>
									{post.description}
								</h3>
								{#if post.username == user.username}
									<form
										id="form-{post.postId}"
										method=" post"
										class="w-full h-fit hidden "
										on:submit|preventDefault={updatePost(post.postId)}
									>
										<input
											type="text"
											id="upDes-{post.postId}"
											placeholder={post.description}
											class="text-base pl-2 my-2 focus:outline-hidden focus:outline-none block w-full h-fit"
										/>
										<input
											type="submit"
											value="Edit"
											class="p-1 text-md font-semibold text-main-bg hover:cursor-pointer py-1 px-2 rounded-md hover:text-white hover:bg-main-bg w-16 border border-main-bg float-right mr-4 mb-1"
										/>
									</form>
								{/if}
							</section>
							<!--bottom part-->
							<section class="flex justify-between w-full items-center px-2">
								<!-- button  part-->
								<section class=" flex justify-start  text-lg  gap-2  ">
									<button
										class="hover:text-red-600 text-2xl p-2 flex items-center w-16 {post.liked
											? 'text-red-600'
											: 'text-gray-400'}"
										on:click={likePost(post.postId)}
										><i class="fa-solid fa-heart" id="like-{post.postId}" />
										<p class="text-sm ml-2 " id="like-n-{post.postId}">
											{post.likes}
										</p></button
									>
									{#if post.allowComments}
										<button
											on:click={openPost(post.postId)}
											class="text-gray-400 hover:text-main-bg text-2xl p-2 w-16 "
											id="comment-{post.postId}"><i class="fa-solid fa-comments" /></button
										>
									{/if}
									<button class="text-gray-400 hover:text-gray-800 text-2xl p-2 w-16" id="share"
										><i class="fa-solid fa-share" /></button
									>
								</section>
								<!--time-->
								<h6 class="text-xs text-orange mx-2">
									{formatDistanceToNow(new Date(post.createdAt), { addSuffix: true })}
								</h6>
							</section>
							<!-- comment part -->
							{#if post.allowComments}
								<section
									class="flex justify-between items-center w-full  p-2 border-t border-solid border-gray-200 "
								>
									{#if user.profileImg}
										<img
											class="h-8 w-8 object-cover rounded-full"
											src="/api/{user.profileImg}"
											alt="Current profile photo"
										/>
									{:else}
										<div class="h-8 w-8 rounded-full  bg-main-bg flex items-center justify-center">
											<i class="fa-solid fa-user text-slate-400 text-lg" />
										</div>
									{/if}
									<form
										method="post"
										on:submit|preventDefault={addComment(post.postId)}
										class="flex-1 flex justify-between  mx-2"
									>
										<input
											type="text"
											placeholder="Add a comment..."
											bind:value={comment}
											class="w-9/12 py-0.5  px-2 focus:outline-hidden focus:outline-none"
										/>
										<button
											id="postBtn-{post.postId}"
											class="border border-main-bg rounded-md py-1 px-3 hover:bg-main-bg hover:text-white font-semibold text-main-bg w-16"
											>Post</button
										>
									</form>
								</section>
							{/if}
						</section>
					</div>
				{/each}
			{/if}
		</div>
	</div>
</div>
