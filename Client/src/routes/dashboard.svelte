<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	import { loading, User } from '../store';
	import { goto } from '$app/navigation';
	let commentChecked = true;
	let hasPhoto;
	let postContent;
	let imageSrc;
	let postPicInput;
	let availablePosts = null;
	let onlineFriends = [
		{
			firstName: 'Sohrab',
			lastName: 'Ebrahimpourian',
			userName: 'sohrab_sosor',
			timeRemian: '30:12'
		},
		{
			firstName: 'Mahsa',
			lastName: 'Ebrahimpourian',
			userName: 'nana',
			timeRemian: '10:12'
		}
	];
	let suggestion = [
		{
			userName: 'saleh Dow',
			firstName: 'saleh',
			lastName: 'Mirzaeeeee',
			onlineTime: '1-2 pm'
		},
		{
			userName: 'jafar',
			firstName: 'jafar',
			lastName: 'Gholami',
			onlineTime: '3-4 am'
		}
	];
	// onMount
	onMount(async () => {
		if (!$User) {
			goto('/');
		}
		const response = await fetch('/api/availablePosts');
		const posts = await response.json();
		availablePosts = posts.availablePosts;
		$loading = false;
	});
	function toggleComment(event) {
		const target = event.target;

		const state = target.getAttribute('aria-checked');

		commentChecked = state === 'true' ? false : true;
	}
	//post new content
	async function newPost() {
		const postPic = document.getElementById('postPic');
		if (!postContent) return;
		const formData = new FormData();
		if (hasPhoto) {
			formData.append('image', postPicInput.files[0]);
		}
		const Response = await fetch('/api/create-post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				description: postContent,
				allowComments: commentChecked
			})
		});
		const postId = await Response.json();
		console.log(`/api/create-postImg/${postId.PostId}`);
		if (Response.status == 201) {
			if (hasPhoto) {
				const sendImage = await fetch(`/api/create-postImg/${postId.PostId}`, {
					method: 'POST',
					body: formData
				});
				if (sendImage.ok) {
					imageSrc.setAttribute('src', '');
				}
			}
			postContent = undefined;
			postPicInput = postPic;
			hasPhoto = false;
		}
	}
	// preview post image
	function previewImg() {
		const file = postPicInput.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', () => {
				imageSrc.setAttribute('src', reader.result);
			});
			reader.readAsDataURL(file);
			hasPhoto = true;
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
			likeBtn.classList.add('text-main');
			likeNumber.classList.add('text-main');
			likeBtn.classList.remove('text-gray-400');
			likeNumber.classList.remove('text-gray-400');
			likeNumber.innerText = Number(likeNumber.innerText) + 1;
		} else {
			likeBtn.classList.remove('text-main');
			likeNumber.classList.remove('text-main');
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
		const comment = document.getElementById(`cm-${postId}`);
		if (comment.value == '' || comment.value == undefined) return;
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
				comment: comment.value,
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
			comment.value = '';
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
			const posts = await response.json();
			availablePosts = posts.availablePosts;
		}
	}
</script>

<svelte:head>
	<title>Dashboard</title>
</svelte:head>
<div class="  md:w-fit w-full lg:w-128 my-2">
	<!--Main part-->
	<div
		id="main"
		class="flex  justify-between items-center md:py-4  gap-4 flex-col  w-full   overflow-x-hidden	 "
	>
		<!--Creat Post-->
		<div
			class="md:border-2 md:border-solid md:border-border shadow-xl w-full rounded-md my-2 overflow-hidden"
		>
			<form
				action="/create-post"
				method="post"
				class="flex-col flex"
				on:submit|preventDefault={newPost}
			>
				{#if hasPhoto}
					<img
						src=""
						bind:this={imageSrc}
						alt=""
						class="w-full h-fit rounded-sm mx-auto my-1 object-cover"
					/>
				{/if}

				<textarea
					name="postContent"
					id="postContent"
					bind:value={postContent}
					cols="20"
					rows="7"
					placeholder="What's up"
					class=" p-2 focus:outline-hidden focus:outline-none resize-none text-text"
				/>
				<section class="flex justify-evenly items-center">
					<!-- comments option -->
					<label for="toggleB" class="flex items-center cursor-pointer">
						<!-- label -->
						<div class="mr-3 text-text font-medium">Comments</div>
						<!-- toggle -->
						<div class="relative">
							<!-- input -->
							<input
								on:change={toggleComment}
								type="checkbox"
								id="toggleB"
								class="sr-only"
								checked
								aria-checked={commentChecked}
							/>
							<!-- line -->
							<div class="block bg-main-bg w-11 h-6 rounded-full border border-border" />
							<!-- dot -->
							<div
								class="dot absolute left-0.5 top-0.5 bg-gray-500 w-5 h-5 rounded-full transition"
							/>
						</div>
					</label>

					<!-- file input-->
					<label for="postPic" class="text-2xl  hover:cursor-pointer hover:text-main  text-text"
						><i class="fa-solid fa-paperclip" /></label
					>
					<input
						bind:this={postPicInput}
						type="file"
						name="image"
						on:change={previewImg}
						id="postPic"
						class="hidden"
					/>
					<!--Submit button-->
					<input
						type="submit"
						value="Post "
						class="p-1 text-xl font-semibold hover:cursor-pointer py-2 px-5 
						text-text
						m-2 rounded-md 
						hover:text-main
						w-24"
					/>
				</section>
			</form>
		</div>
		<!--post-->
		{#if availablePosts}
			{#each availablePosts as post}
				<!-- post outline-->
				<div
					id="body-{post.postId}"
					class="md:border-2 border-solid border-border  shadow-xl w-full rounded-md my-2 overflow-x-hidden"
				>
					<!-- svelte-ignore a11y-img-redundant-alt -->
					<section class=" flex justify-between items-center flex-col ">
						<!--post info-->
						<section
							class="  p-2 flex  items-center gap-2 justify-between w-full border-b border-solid border-border"
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
											class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center border-2 border-border"
										>
											<i class="fa-solid fa-user text-slate-400 text-2xl" />
										</div>
									{/if}
								</a>
								<!-- Name and username-->
								<a href="/profile/{post.username}">
									<h4 class="mx-2 font-semibold text-text hover:text-text-hover">
										{post.firstName}
										{post.lastName}
									</h4>
									<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
										@{post.username}
									</h5>
								</a>
							</section>
							<section class="flex items-center ">
								{#if post.username == $User.username}
									<div class="relative text-text hover:text-main">
										<i
											on:click={postOption(post.postId)}
											class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main"
										/>
										<div
											id={post.postId}
											class="hidden absolute bg-main-bg w-32 flex flex-col items-center  rounded p-3  option gap-2 z-10 border-2 border-border"
										>
											<button
												on:click={openEdit(post.postId)}
												class="text-sm flex items-center w-full justify-start hover:text-main text-white"
											>
												<i class="fa-solid fa-pen-to-square mr-1 text-xs" /> Edit Post
											</button>
											<button
												on:click={deletePost(post.postId)}
												class="text-sm flex items-center w-full justify-start hover:text-main text-white"
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
							<div class="flex items-center my-4 mx-2">
								<a href="/profile/{post.username} ">
									<h3 class="font-semibold hover:text-text-hover text-text">{post.username}</h3>
								</a>
								<h3
									class="text-base mx-2 text-text"
									on:dblclick={likePost(post.postId)}
									id="des-{post.postId}"
								>
									{post.description}
								</h3>
							</div>
							{#if post.username == $User.username}
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
									class="hover:text-main text-2xl p-2 flex items-center w-16 {post.liked
										? 'text-main'
										: 'text-text'}"
									on:click={likePost(post.postId)}
									><i class="fa-solid fa-heart" id="like-{post.postId}" />
									<p class="text-sm ml-2 " id="like-n-{post.postId}">
										{post.likes}
									</p></button
								>
								{#if post.allowComments}
									<button
										on:click={openPost(post.postId)}
										class="text-text hover:text-border text-2xl p-2 w-16 "
										id="comment-{post.postId}"><i class="fa-solid fa-comment" /></button
									>
								{/if}
								<button class="text-text hover:text-gray-800 text-2xl p-2 w-16" id="share"
									><i class="fa-solid fa-share" /></button
								>
							</section>
							<!--time-->
							<h6 class="text-xs text-text mx-2">
								{formatDistanceToNow(new Date(post.createdAt), { addSuffix: true })}
							</h6>
						</section>
						<!-- comment part -->
						{#if post.allowComments}
							<section
								class="flex justify-between items-center w-full  p-2 border-t border-solid border-border "
							>
								{#if $User.profileImg}
									<img
										class="h-8 w-8 object-cover rounded-full"
										src="/api/{$User.profileImg}"
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
										id="cm-{post.postId}"
										class="w-9/12 py-0.5  px-2 focus:outline-hidden focus:outline-none text-text bg-inherit"
									/>
									<button id="postBtn-{post.postId}" class="main-btn w-16">Post</button>
								</form>
							</section>
						{/if}
					</section>
				</div>
			{/each}
		{/if}
	</div>
</div>

<style>
	input:checked ~ .dot {
		transform: translateX(100%);
		background-color: #fff;
	}
</style>
