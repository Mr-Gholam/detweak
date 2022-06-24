<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch('/api/find-jobs', {
			credentials: 'include'
		});
		let jobs;
		if (res.status == 200) {
			jobs = await res.json();
		}
		return { props: { jobs } };
	};
</script>

<script>
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	export let jobs;
	let title;
	let budget;
	function openFilter() {
		const filter = document.getElementById('filter');
		if (filter.classList.contains('hidden')) {
			filter.classList.replace('hidden', 'block');
		} else {
			filter.classList.replace('block', 'hidden');
		}
	}
</script>

<svelte:head>
	<title>Find job</title>
</svelte:head>

<div class=" md:w-fit w-full lg:w-128 my-2 py-4">
	<div class="w-full lg:w-128  mx-auto">
		<!-- search -->
		<button class="main-btn w-24" on:click={openFilter}> Filter </button>
		<form class="my-4 hidden" id="filter">
			<section class="w-80">
				<input
					placeholder="Tilte,Feild,Language or Framwork"
					bind:value={title}
					type="text"
					class="outline-none border-2   rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text border-border"
				/>
			</section>

			<section class="w-80 relative">
				<input
					bind:value={budget}
					type="number"
					placeholder="Budget"
					class="outline-none border-2  border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text border-border"
				/>
				<p class="text-text absolute right-0 inset-y-1/2 mr-6 top-3">$</p>
			</section>
			<section class="w-80 mx-auto">
				<input
					id="submitPersonal"
					type="submit"
					value="Search"
					class="main-btn w-11/12 py-2 px-20 text-lg mx-auto block border-2"
				/>
			</section>
		</form>
		<!-- jobs -->
		<div class="my-4">
			<h1 class="text-center text-lg font-semibold my-2 text-white">Available Jobs</h1>
			{#if jobs}
				{#each jobs as job}
					<div
						class="md:border-2 border-solid border-border  shadow-xl w-full rounded-md my-2 overflow-x-hidden p-2"
					>
						<section class="flex gap-2 items-center justify-between">
							<!--profile img-->
							<a href="/profile/{job.OwnerUsername}" class="ml-2">
								{#if job.OwnerImg}
									<!-- svelte-ignore a11y-img-redundant-alt -->
									<img
										class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
										src="/api/images/{job.OwnerImg}"
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
							<section class="flex-1">
								<a href="/job/{job.Id}" class="w-fit">
									<h4 class="mx-2 font-semibold text-text hover:text-text-hover ">
										{job.Title}
									</h4>
									<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
										{job.ProjectName}
									</h5>
								</a>
								<h5 class=" text-sm  text-text  mx-2 flex items-center justify-between ">
									<h6>
										<i class="fa-solid fa-money-bill-wave text-sm mr-2" />{job.Budget}$
									</h6>
									<h6 class="text-xs text-text mx-2">
										{formatDistanceToNow(new Date(job.CreatedAt), { addSuffix: true })}
									</h6>
								</h5>
							</section>
							<section class="flex flex-col justify-end items-center ">
								<a
									href="/job/{job.Id}"
									class="border-border p-2 border-2 rounded-full text-sm text-text hover:border-main hover:text-white"
									>See More</a
								>
							</section>
						</section>
					</div>
				{/each}
			{:else}
				<div class="my-8">
					<div class="text-white  mx-auto w-fit   ">
						<a
							href="/setting"
							class="text-center w-fit mx-auto border-border md:border-2 py-3 px-5 rounded-full text-sm  hover:text-main hover:border-main "
						>
							Update your professional information
						</a>
					</div>
				</div>
			{/if}
		</div>
	</div>
</div>
