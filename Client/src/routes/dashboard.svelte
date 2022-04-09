<script>
	let postContent;
	let imageSrc;
	let postPicInput;
	let avileblePosts = [
		{
			userName: 'mehdi',
			post: 'The w-auto utility can be useful if you need to remove an element’s assigned width under a specific condition, like at a particular breakpoint:',
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
	// preview post image
	function previewImg() {
		const file = postPicInput.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', () => {
				imageSrc.setAttribute('src', reader.result);
				console.log(reader.result);
			});
			reader.readAsDataURL(file);
		}
	}
</script>

<main class="flex justify-between 2xl:w-9/12 2xl:mx-auto items-start">
	<!--left part-->
	<div class=" h-fit p-4 my-2">
		<section class="border-2 border-main">
			<h4>post</h4>
			<h4>post</h4>
			<h4>post</h4>
			<h4>post</h4>
			<h4>post</h4>
		</section>
	</div>
	<!--Main part-->
	<div class="flex  justify-between items-center p-4 gap-4 flex-col w-6/12">
		<!--Creat Post-->
		<div
			class="border-2 border-solid border-gray-200 shadow-xl w-96 rounded-md my-2 overflow-hidden"
		>
			<form
				action="/create-post"
				method="post"
				class="flex-col flex"
				on:submit|preventDefault={newPost}
			>
				<img src="" bind:this={imageSrc} alt="" class="w-80 h80 rounded-sm mx-auto my-1" />
				<textarea
					name="postContent"
					id="postContent"
					bind:value={postContent}
					cols="20"
					rows="7"
					placeholder="What's up"
					class=" p-2 focus:outline-hidden focus:outline-none resize-none"
				/>
				<section class="flex justify-evenly items-center">
					<!-- file input-->
					<label for="postPic" class="text-2xl  hover:cursor-pointer "
						><i class="fa-solid fa-paperclip" /></label
					>
					<input
						bind:this={postPicInput}
						type="file"
						name="postPic"
						on:change={previewImg}
						id="postPic"
						class="hidden"
					/>
					<!--Submit button-->
					<input
						type="submit"
						value="Post "
						class="p-1 text-xl font-semibold hover:cursor-pointer py-2 px-5 
						m-2 rounded-md 
						w-24"
					/>
				</section>
			</form>
		</div>
		<!--post-->
		{#each avileblePosts as post}
			<!-- post outline-->
			<div class="border-2 border-solid border-gray-200 shadow-xl w-96 rounded-md my-2">
				<!-- svelte-ignore a11y-img-redundant-alt -->
				<section class=" flex justify-between items-center flex-col ">
					<!--post info-->
					<section class="  p-2 flex  items-center gap-2 justify-between w-full   ">
						<!--name and username-->
						<section class="flex gap-2 ">
							<!--profile img-->
							<a href="/{post.userName}">
								<img
									class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
									src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
									alt="Current profile photo"
								/>
							</a>
							<!-- Name and username-->
							<a href="/{post.userName}">
								<h4 class="mx-2 font-semibold text-gray-900 hover:text-gray-500">
									{post.firstName}
									{post.lastName}
								</h4>
								<h5 class=" text-sm font-semibold text-gray-900 hover:text-gray-500 mx-2 ">
									{post.userName}
								</h5>
							</a>
						</section>
						<h6 class="text-xs text-gray-500 mx-2">{post.time}</h6>
					</section>
					<!-- post-->
					<section>
						<img
							class="w-full h-64  mx-auto"
							src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQu9PvSvDgRqonJcU8peJcUrC8ydfa_lEm9ag&usqp=CAU"
							alt=""
						/>
						<h3 class="text-base mx-2 my-4">{post.post}</h3>
					</section>
					<!-- button  part-->
					<section class=" flex justify-end  text-lg  gap-2 w-full mr-4">
						<input type="hidden" value={post.postId} />
						<button class="text-gray-400 hover:text-red-600 text-2xl p-2"
							><i class="fa-solid fa-heart" id="likes" /></button
						>
						<button class="text-gray-400 hover:text-main-lighter text-2xl p-2" id="comments"
							><i class="fa-solid fa-comments" /></button
						>
						<button class="text-gray-400 hover:text-gray-800 text-2xl p-2" id="share"
							><i class="fa-solid fa-share" /></button
						>
					</section>
				</section>
			</div>
		{/each}
	</div>
	<!--Right part-->
	<div class=" h-fit p-4">
		<section class="border-2 border-main">
			<h3>friends</h3>
		</section>
	</div>
</main>
