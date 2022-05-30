<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	import { loading, User } from '../store';
	import { goto } from '$app/navigation';
	let user;
	User.subscribe((value) => (user = value));
	let comment;
	let likedPosts = null;
	onMount(async () => {
		const response = await fetch('/api/liked-posts');
		const data = await response.json();
		likedPosts = data;
	});
	// like post
	async function likePost(postId) {
		const body = document.getElementById('body');
		const post = document.getElementById(`body-${postId}`);
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
		if (data.Added) {
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
			post.classList.add('right-exit');
			setTimeout(() => {
				body.removeChild(post);
			}, 1000);
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
			const response = await fetch('/api/likedPosts');
			const posts = await response.json();
			likedPosts = posts.likedPosts;
		}
	}
</script>

<svelte:head>
	<title>Liked Posts</title>
</svelte:head>

<!--posts-->
<div
	id="body"
	class="flex  justify-between items-center py-4 my-2  gap-4 flex-col  w-96 lg:w-128 md:mr-32 lg:mr-0 transition-all overflow-x-hidden"
>
	{#if likedPosts}
		{#each likedPosts as post}
			<!-- post outline-->
			<div
				id="body-{post.PostId}"
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
							<a href="/profile/{post.Username}" class="ml-2">
								{#if post.ProfileImg}
									<img
										class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
										src="/api/images/{post.ProfileImg}"
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
							<a href="/profile/{post.Username}">
								<h4 class="mx-2 font-semibold text-text hover:text-text-hover">
									{post.Firstname}
									{post.Lastname}
								</h4>
								<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
									@{post.Username}
								</h5>
							</a>
						</section>
						<section class="flex items-center ">
							{#if post.Username == user.username}
								<div class="relative text-text hover:text-main">
									<i
										on:click={postOption(post.PostId)}
										class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main"
									/>
									<div
										id={post.PostId}
										class="hidden absolute bg-main-bg w-32 flex flex-col items-center  rounded p-3  option gap-2 z-10 border-2 border-border"
									>
										<button
											on:click={openEdit(post.PostId)}
											class="text-sm flex items-center w-full justify-start hover:text-main text-white"
										>
											<i class="fa-solid fa-pen-to-square mr-1 text-xs" /> Edit Post
										</button>
										<button
											on:click={deletePost(post.PostId)}
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
						{#if post.PostImgUrl}
							<img
								on:dblclick={likePost(post.PostId)}
								class="w-full h-fit  md:mx-auto  object-cover"
								src="/api/{post.PostImgUrl}"
								alt=""
							/>
						{/if}
						<div class="flex items-center my-4 mx-2">
							<a href="/profile/{post.username} ">
								<h3 class="font-semibold hover:text-text-hover text-text">{post.Username}</h3>
							</a>
							<h3
								class="text-base mx-2 text-text"
								on:dblclick={likePost(post.PostId)}
								id="des-{post.PostId}"
							>
								{post.Description}
							</h3>
						</div>
						{#if post.Username == user.username}
							<form
								id="form-{post.PostId}"
								method=" post"
								class="w-full h-fit hidden "
								on:submit|preventDefault={updatePost(post.PostId)}
							>
								<input
									type="text"
									id="upDes-{post.PostId}"
									placeholder={post.Description}
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
								class="hover:text-main text-2xl p-2 flex items-center w-16 {post.Liked
									? 'text-main'
									: 'text-text'}"
								on:click={likePost(post.PostId)}
								><i class="fa-solid fa-heart" id="like-{post.PostId}" />
								<p class="text-sm ml-2 " id="like-n-{post.PostId}">
									{post.Likes}
								</p></button
							>
							{#if post.AllowComments}
								<button
									on:click={openPost(post.PostId)}
									class="text-text hover:text-border text-2xl p-2 w-16 "
									id="comment-{post.PostId}"><i class="fa-solid fa-comment" /></button
								>
							{/if}
							<button class="text-text hover:text-gray-800 text-2xl p-2 w-16" id="share"
								><i class="fa-solid fa-share" /></button
							>
						</section>
						<!--time-->
						<h6 class="text-xs text-text mx-2">
							{formatDistanceToNow(new Date(post.CreatedAt), { addSuffix: true })}
						</h6>
					</section>
					<!-- comment part -->
					{#if post.AllowComments}
						<section
							class="flex justify-between items-center w-full  p-2 border-t border-solid border-border "
						>
							{#if user.ImgUrl}
								<img
									class="h-8 w-8 object-cover rounded-full"
									src="/api/images/{user.ImgUrl}"
									alt="Current profile photo"
								/>
							{:else}
								<div class="h-8 w-8 rounded-full  bg-main-bg flex items-center justify-center">
									<i class="fa-solid fa-user text-slate-400 text-lg" />
								</div>
							{/if}
							<form
								method="post"
								on:submit|preventDefault={addComment(post.PostId)}
								class="flex-1 flex justify-between  mx-2"
							>
								<input
									type="text"
									placeholder="Add a comment..."
									id="cm-{post.PostId}"
									class="w-9/12 py-0.5  px-2 focus:outline-hidden focus:outline-none text-text bg-inherit"
								/>
								<button
									id="postBtn-{post.PostId}"
									class="border border-border rounded-md py-1 px-3 hover:border-main hover:text-main font-semibold text-text w-16 "
									>Post</button
								>
							</form>
						</section>
					{/if}
				</section>
			</div>
		{/each}
	{:else}
		<h4>ridi</h4>
	{/if}
</div>
