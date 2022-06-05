<script>
	// @ts-nocheck

	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { loading, User, ws } from '../../store';
	import { goto } from '$app/navigation';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	let suggestionUsers;
	let suggestedPosts;
	onMount(async () => {
		const userInput = $page.params.userInput;
		const response = await fetch(`/api/search/${userInput}`);
		const data = await response.json();
		console.log(data);
		suggestionUsers = JSON.parse(JSON.stringify(data.Users));
		suggestedPosts = JSON.parse(JSON.stringify(data.Posts));
		$loading = false;
	});
	async function addFriend(username) {
		if ($User) {
			const response = await fetch('/api/add-friend', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username
				})
			});
			const data = await response.json();
			if (response.ok) {
				const user = suggestionUsers.find((user) => user.Username == username);
				const post = suggestedPosts.find((post) => post.Username == username);
				user.IsFriend = data.status;
				post.IsFriend = data.status;
				suggestionUsers = suggestionUsers;
				suggestedPosts = suggestedPosts;
				$ws.send(JSON.stringify({ Target: username }));
			}
		} else {
			goto('/login');
		}
	}
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
		if (data.Added) {
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
			if (likeNumber.innerText === '0') return;
			likeNumber.innerText = Number(likeNumber.innerText) - 1;
		}
	}
	//  open post option
	function postOption(postId) {
		const option = document.getElementById(`${postId}`);
		if (option.classList.contains('hidden')) {
			option.classList.replace('hidden', 'flex');
		} else {
			option.classList.replace('flex', 'hidden');
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
		const oldDes = document.getElementById(`des-${postId}`);
		const updatedDes = document.getElementById(`upDes-${postId}`);
		const res = await fetch('/api/update-post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				newDescription: updatedDes.value,
				postId
			})
		});
		if (res.status == 200) {
			openEdit(postId);
			oldDes.innerText = updatedDes.value;
		}
	}
	function openChat(username) {
		goto(`/messages?username=${username}`);
	}
</script>

<svelte:head>
	<title>Search</title>
</svelte:head>

<div class="w-96 py-4 my-2  lg:mx-auto lg:w-128 {$User.username ? 'lg:ml-32' : 'lg:ml-72'}">
	<div class="flex flex-col">
		<div class="flex flex-col">
			{#if suggestionUsers}
				<h1 class="text-text text-center">Users</h1>
				{#each suggestionUsers as suggedtedPeople}
					<section
						class="  py-2 px-3 flex my-2 items-center  justify-between w-full border-2 border-border rounded-md "
					>
						<!--name and username-->
						<section class="flex  items-center w-8/12 ">
							<!--profile img-->
							<a href="/profile/{suggedtedPeople.Username}">
								<!-- svelte-ignore a11y-img-redundant-alt -->
								{#if suggedtedPeople.ImgUrl}
									<img
										class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
										src="/api/images/{suggedtedPeople.ImgUrl}"
										alt="Current profile photo"
									/>
								{:else}
									<div
										class="h-12 w-12 rounded-full hover:opacity-90  flex items-center justify-center border-border border-2"
									>
										<i class="fa-solid fa-user text-slate-400 text-2xl" />
									</div>
								{/if}
							</a>
							<!-- Name and username-->
							<a href="/profile/{suggedtedPeople.Username}" class="mx-2 w-9/12">
								<h4 class=" font-semibold text-sm text-text hover:text-text-hover ">
									{suggedtedPeople.Firstname}
									{suggedtedPeople.Lastname}
								</h4>
								<h5 class=" text-xs  text-text hover:text-text-hover">
									@{suggedtedPeople.Username}
								</h5>
							</a>
						</section>
						{#if suggedtedPeople.Username != $User.username}
							{#if suggedtedPeople.IsFriend == 'Pending'}
								<button class="main-btn ">
									<h5 class="mx-2	 text-sm">Friend Request Sent</h5>
								</button>
							{:else if suggedtedPeople.IsFriend == 'Friend'}
								<button class="main-btn" on:click={openChat(suggedtedPeople.Username)}
									>Sent Massage</button
								>
							{:else}
								<button on:click={addFriend(suggedtedPeople.Username)} class=" main-btn"
									>Add Friend
								</button>
							{/if}
						{/if}
					</section>
				{/each}
			{/if}
			{#if suggestionUsers == null}
				<h1 class="text-text text-center">Users not found</h1>
			{/if}
		</div>
		<div class="flex flex-col my-4">
			<h1 class="text-text text-center">Posts</h1>
			{#if suggestedPosts}
				{#each suggestedPosts as post}
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
												alt="Profile photo"
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
									{#if post.Username == $User.username}
										<div class="relative text-text hover:text-main">
											<i
												on:click={postOption(post.PostId)}
												class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main"
											/>
											<div
												id={post.PostId}
												class="hidden absolute bg-main-bg w-32  flex-col items-center  rounded p-3  option gap-2 z-10 border-2 border-border"
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
									{:else if post.IsFriend == 'Pending'}
										<button class="main-btn ">
											<h5 class="mx-2	 text-sm">Friend Request Sent</h5>
										</button>
									{:else if post.IsFriend == 'Friend'}
										<button class="main-btn" on:click={openChat(post.Username)}>Sent Massage</button
										>
									{:else}
										<button on:click={addFriend(post.Username)} class=" main-btn"
											>Add Friend
										</button>
									{/if}
								</section>
							</section>
							<!-- post-->
							<section class="w-full h-fit">
								{#if post.PostImgUrl}
									<img
										on:dblclick={likePost(post.PostId)}
										class="w-full h-fit  md:mx-auto  object-cover"
										src="/api/images/{post.PostImgUrl}"
										alt=""
									/>
								{/if}
								<div class="flex items-center my-4 mx-2">
									<a href="/profile/{post.Username} ">
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
								{#if post.Username == $User.username}
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
											class="text-base pl-2 my-2 focus:outline-hidden focus:outline-none block w-full h-fit text-text"
										/>
										<input type="submit" value="Edit" class="main-btn float-right mr-4 w-16" />
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
									{#if $User.ImgUrl}
										<img
											class="h-8 w-8 object-cover rounded-full"
											src="/api/images/{$User.ImgUrl}"
											alt="Current profile photo"
										/>
									{:else}
										<div
											class="h-8 w-8 rounded-full  bg-main-bg flex items-center justify-center border border-border"
										>
											<i class="fa-solid fa-user text-slate-400 text-md" />
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
										<button id="postBtn-{post.PostId}" class="main-btn w-16">Post</button>
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
