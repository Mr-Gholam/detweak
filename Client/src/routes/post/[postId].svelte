<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	import { User } from '../../store';
	let user;
	let post;
	let comment;
	let comments = [];
	User.subscribe((value) => (user = value));
	onMount(async () => {
		const postId = $page.params.postId;
		const response = await fetch(`/post/${postId}`);
		const data = await response.json();
		post = JSON.parse(JSON.stringify(data.postInfo[0]));
		comments = JSON.parse(JSON.stringify(data.comments));
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
		const option = document.getElementById(`${commentId}`);
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
		const response = await fetch('/like-post', {
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
		const response = await fetch('/add-comment', {
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
</script>

<div
	class="flex items-start justify-center md:justify-between md:py-4  gap-4 flex-col  w-96 lg:w-128  xl:max-w-9/12 md:mx-auto"
>
	<!-- post outline-->
	<div class="md:border-2 border-solid border-gray-200  shadow-xl w-full rounded-md my-2">
		<!-- svelte-ignore a11y-img-redundant-alt -->
		<section class=" flex justify-between items-center flex-col ">
			{#if post}
				<section class="  p-2 flex  items-center gap-2 justify-between w-full ">
					<!--post info-->
					<!--name and username-->
					<section class="flex gap-2 items-center">
						<!--profile img-->
						<a href="/profile/{post.username}" class="ml-2">
							{#if post.profileImg}
								<img
									class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
									src="http://localhost:8585/{post.profileImg}"
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
						{#if post.username == user}
							<div class="relative">
								<i
									on:click={postOption(post.postId)}
									class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main-bg"
								/>
								<div
									id={post.postId}
									class="hidden absolute bg-main-bg w-28 flex flex-col items-center  rounded p-3  option gap-2"
								>
									<button
										class="text-sm flex items-center w-full justify-start hover:text-main text-white"
									>
										<i class="fa-solid fa-pen-to-square mr-1 text-xs" /> Edit Post
									</button>
									<button
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
					<img
						class="w-full h-fit  md:mx-auto  object-cover"
						src="http://localhost:8585/{post.postImg}"
						alt=""
					/>
					<h3 class="text-base mx-2 my-4">{post.description}</h3>
				</section>
				<!--bottom part-->
				<section
					class="flex justify-between w-full items-center px-2   border-b border-solid border-gray-200"
				>
					<!-- button  part-->
					<section class=" flex justify-start  text-lg  gap-2 ">
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
						<button class="text-gray-400 hover:text-main-bg text-2xl p-2 w-16" id="comments"
							><i class="fa-solid fa-comments" /></button
						>
						<button class="text-gray-400 hover:text-gray-800 text-2xl p-2 w-16" id="share"
							><i class="fa-solid fa-share" /></button
						>
					</section>
					<!--time-->
					<h6 class="text-xs text-orange mx-2">
						{formatDistanceToNow(new Date(post.createdAt), { addSuffix: true })}
					</h6>
				</section>
			{/if}
			<!-- comments part -->
			<section class="flex flex-col gap-2 w-full">
				{#if comments}
					{#each comments as comment}
						<div class="flex items-center w-full px-4 my-1">
							{#if comment.profileImg}
								<a href="/profile/{comment.username}" class="flex w-fit  items-center ">
									<img
										class="h-6 w-6 object-cover rounded-full"
										src="http://localhost:8585/{comment.profileImg}"
										alt="Current profile photo"
									/>
									<h4 class="font-semibold text-base ml-1">
										{comment.username}
									</h4>
								</a>
							{:else}
								<a href="profile/{comment.username}" class="flex w-fit items-center">
									<div class="h-8 w-8 rounded-full  bg-main-bg flex items-center justify-center">
										<i class="fa-solid fa-user text-slate-400 text-sm" />
									</div>
									<h4 class="font-semibold text-base ml-1">
										{comment.username}
									</h4>
								</a>
							{/if}
							<h4 class="flex-1 text-sm mx-2 ">
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
											class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main-bg"
										/>
										<div
											id={comment.commentId}
											class="hidden absolute bg-main-bg w-32 flex flex-col items-center  rounded p-3  option gap-2 z-10"
										>
											<button
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
			<!-- add comment part -->
			<section
				class="flex justify-between items-center w-full  p-2 border-t border-solid border-gray-200 "
			>
				{#if user.profileImg}
					<img
						class="h-8 w-8 object-cover rounded-full"
						src="http://localhost:8585/{user.profileImg}"
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
						class="border-2 border-main-bg rounded-xl py-1 px-3 hover:bg-main-bg hover:text-white font-semibold"
						>Post</button
					>
				</form>
			</section>
		</section>
	</div>
</div>
