<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	import { User } from '../../store';
	let user;
	let post;
	let comment;
	let comments = [];
	User.subscribe((value) => (user = value));
	onMount(async () => {
		const postId = $page.params.postId;
		const response = await fetch(`/api/post/${postId}`);
		const data = await response.json();
		post = data.postInfo[0];
		comments = data.comments;
	});
	//  open post option
	function postOption(postId) {
		const option = document.getElementById(`${postId}`);
		if (option.classList.contains('hidden')) {
			option.classList.remove('hidden');
		} else {
			option.classList.add('hidden');
		}
	}
	//  open post option comment
	function postOptionComment(commentId) {
		const option = document.getElementById(`cm-${commentId}`);
		if (option.classList.contains('hidden')) {
			option.classList.remove('hidden');
		} else {
			option.classList.add('hidden');
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
	// add comment
	async function addComment(postId) {
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
		const data = await response.json();
		if (response.status == 200) {
			console.log(data.createdAt);
			const newComment = {
				commentContent: data.commentContent,
				profileImg: data.profileImg,
				username: data.username,
				createdAt: data.createdAt,
				commentId: data.commentId
			};
			comments = [...comments, newComment];
			comment = '';
		}
	}
	//  delete comment
	async function deleteComment(commentId) {
		const comments = document.getElementById('comments');
		const comment = document.getElementById(`c-${commentId}`);
		const option = document.getElementById(`${commentId}`);
		const response = await fetch('/api/delete-Comment', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				commentId
			})
		});
		if (response.status == 200) {
			option.classList.add('hidden');
			comment.classList.add('right-exit');
			setTimeout(() => {
				comment.parentNode.removeChild(comment);
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
			const response = await fetch(`/api/post/${postId}`);
			const data = await response.json();
			post = JSON.parse(JSON.stringify(data.postInfo[0]));
			comments = JSON.parse(JSON.stringify(data.comments));
		}
	}
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
</script>

<svelte:head>
	<title>Post</title>
</svelte:head>

{#if post}
	<div class="  md:w-fit w-full lg:w-128 my-2">
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
						{#if post.username == user.username}
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
				<!-- comments -->
				<section class="flex flex-col gap-2 w-full" id="comments">
					{#if comments}
						{#each comments as comment}
							<div class="flex items-center w-full px-4 my-1" id="c-{comment.commentId}">
								{#if comment.profileImg}
									<a href="/profile/{comment.username}" class="flex w-fit  items-center ">
										<img
											class="h-6 w-6 object-cover rounded-full"
											src="/api/{comment.profileImg}"
											alt="Current profile photo"
										/>
										<h4 class="font-semibold text-base ml-1 text-text hover:text-text-hover">
											{comment.username}
										</h4>
									</a>
								{:else}
									<a href="profile/{comment.username}" class="flex w-fit items-center">
										<div class="h-8 w-8 rounded-full  bg-main-bg flex items-center justify-center">
											<i class="fa-solid fa-user text-slate-400 text-sm" />
										</div>
										<h4 class="font-semibold text-base ml-1 text-text hover:text-text-hover">
											{comment.username}
										</h4>
									</a>
								{/if}
								<h4 class="flex-1 text-sm mx-2 text-text  ">
									{comment.commentContent}
								</h4>
								<section class="flex items-center ">
									<h6 class="text-xs text-gray-500 ">
										{formatDistanceToNow(new Date(comment.createdAt), { addSuffix: true })}
									</h6>
									{#if post.username == user.username}
										<div class="relative">
											<i
												on:click={postOptionComment(comment.commentId)}
												class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main text-text "
											/>
											<div
												id="cm-{comment.commentId}"
												class="hidden absolute bg-main-bg w-32 flex flex-col items-center  rounded p-3  option gap-2 z-10"
											>
												<button
													on:click={deleteComment(comment.commentId)}
													class="text-xs flex items-center w-full justify-start hover:text-red-600 text-white"
												>
													<i class="fa-solid fa-trash mr-1 text-xs" /> Delete Comment
												</button>
											</div>
										</div>
									{/if}
								</section>
							</div>
						{/each}
					{/if}
				</section>

				<!-- comment part -->
				{#if post.allowComments}
					<section
						class="flex justify-between items-center w-full  p-2 border-t border-solid border-border "
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
								id="cm-{post.postId}"
								class="w-9/12 py-0.5  px-2 focus:outline-hidden focus:outline-none text-text bg-inherit"
							/>
							<button id="postBtn-{post.postId}" class="main-btn w-16">Post</button>
						</form>
					</section>
				{/if}
			</section>
		</div>
	</div>
{/if}
