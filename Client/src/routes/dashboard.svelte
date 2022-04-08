<script>
	let postContent;
	let avileblePosts = [
		{
			userName: 'mehdi',
			post: 'this is the test',
			time: '1 hour ago',
			postId: '553'
		},
		{
			userName: 'ali',
			post: 'this is the test2',
			time: '3 hour ago',
			postId: '123'
		}
	];
	//publish the new post
	function publish(data) {
		avileblePosts.push(data);
	}
	//post new content
	function newPost() {
		if (!postContent) return;
		// let data = {
		// 	userName: 'mamad',
		// 	post: postContent,
		// 	time: '1 min ago',
		// 	postId: '111'
		// };
		// publish(data);
		console.log(avileblePosts);
		fetch('http://localhost:8585/create-post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				postContent
			})
		})
			.then((res) => {
				res.json();
				// console.log(res);
			})
			.then((data) => {
				// console.log(data);
			});
	}
</script>

<main>
	<div class="flex  justify-between items-center p-4 gap-4 flex-col">
		<!--Creat Post-->
		<div class="border-2 border-solid border-main w-96 rounded-md my-2 overflow-hidden">
			<form
				action="/create-post"
				method="post"
				class="flex-col flex"
				on:submit|preventDefault={newPost}
			>
				<textarea
					name="postContent"
					id="postContent"
					bind:value={postContent}
					cols="20"
					rows="7"
					placeholder="What's up"
					class=" p-2 focus:outline-hidden focus:outline-none resize-none"
				/>
				<input
					type="submit"
					value="Post "
					class="p-1 text-lg font-semibold bg-slate-600 hover:bg-slate-700 border-t-2 border-main hover:text-main-darker"
				/>
			</form>
		</div>
		<!--post-->
		{#each avileblePosts as post}
			<div class="border-2 border-solid border-main w-96 rounded-md my-2">
				<!-- the top bar of the post-->
				<section class="flex justify-between items-center  ">
					<!-- svelte-ignore a11y-img-redundant-alt -->
					<section class=" flex justify-between items-center ">
						<section class="  p-2 flex flex-col items-center ">
							<a href="/{post.userName}">
								<img
									class="h-12 w-12 object-cover rounded-full hover:opacity-90 "
									src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
									alt="Current profile photo"
								/>
								<h5
									class="my-1  font-semibold text-sm text-gray-500 hover:text-gray-900 text-center"
								>
									{post.userName}
								</h5>
							</a>
							<h6 class="text-xs text-gray-500 mx-2">{post.time}</h6>
						</section>
						<h3 class="text-base mx-2">{post.post}</h3>
					</section>
					<!-- right part-->
					<section class=" p-2 flex flex-col justify-evenly text-lg border-l-2 border-main gap-1">
						<input type="hidden" value={post.postId} />
						<button class="text-gray-400 hover:text-red-600 "
							><i class="fa-solid fa-heart" id="likes" /></button
						>
						<button class="text-gray-400 hover:text-main-lighter" id="comments"
							><i class="fa-solid fa-comments" /></button
						>
						<button class="text-gray-400 hover:text-gray-800" id="share"
							><i class="fa-solid fa-share" /></button
						>
					</section>
				</section>
			</div>
		{/each}
	</div>
</main>
