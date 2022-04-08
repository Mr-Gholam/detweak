<main>
	<div>
		<!--post-->
		<div class="border-2 border-solid border-main w-96 rounded-md ">
			<!-- the top bar of the post-->
			<section class="flex justify-between items-center border-b-2 border-main p-2">
				<!-- svelte-ignore a11y-img-redundant-alt -->
				<section class=" flex justify-between items-center ">
					<a href="/:porfile" class="ml-2">
						<img
							class="h-12 w-12 object-cover rounded-full hover:opacity-90 "
							src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
							alt="Current profile photo"
						/>
					</a>
					<section class="mx-3">
						<h4 class="my-1 text-xl  font-bold">Post title</h4>
						<a href="/:profile" class="text-sm text-gray-500 hover:text-gray-900">
							<h5 class="my-1  font-semibold">UserName</h5>
						</a>
					</section>
				</section>
				<h6 class="text-xs text-gray-500 mx-2">1 hour ago</h6>
			</section>
			<section />
			<!-- middle part of the post-->
			<section class=" flex justify-between">
				<!-- left part-->
				<section class="p-2">
					<h3 class="text-base">
						Lorem ipsum dolor sit amet consectetur, adipisicing elit. Fugiat aliquam illo debitis
						harum ad.
					</h3>
				</section>
				<!-- right part-->
				<section class="border-l-2 border-main p-1 flex flex-col justify-evenly text-lg">
					<button class="text-gray-400 hover:text-red-600 "
						><i class="fa-solid fa-heart" id="likes" /></button
					>
					<button class="text-gray-400 hover:text-main-lighter" id="comments"
						><i class="fa-solid fa-comments" /></button
					>
					<button class="text-gray-400 hover:text-main-lighter" id="share"
						><i class="fa-solid fa-share" /></button
					>
				</section>
			</section>
		</div>
	</div>
</main>
