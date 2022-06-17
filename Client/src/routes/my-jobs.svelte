<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch('/api/my-jobs', {
			credentials: 'include'
		});
		let jobs;
		if (res.status == 200) {
			jobs = await res.json();
			console.log(res);
		}
		return { props: { jobs } };
	};
</script>

<script>
	export let jobs;
	console.log(jobs);
</script>

<svelte:head>
	<title>My Job</title>
</svelte:head>
<div class=" md:w-fit w-full lg:w-128 my-2">
	<div class="w-full lg:w-128  mx-auto">
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
						<a href="/job/{job.Id}" class="flex-1">
							<h4 class="mx-2 font-semibold text-text hover:text-text-hover">
								{job.Title}
							</h4>
							<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
								{job.ProjectName}
							</h5>
							<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
								<i class="fa-solid fa-money-bill-wave text-sm mr-2" />{job.Budget}$
							</h5>
						</a>
						<section class="flex flex-col justify-end items-center relative ">
							<p
								class="w-5 h-5 bg-main rounded-full flex items-center justify-center text-white text-sm absolute -top-2 -right-2"
							>
								4
							</p>
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
			<h4 class="text-text">You haven't created a job yet</h4>
			<a href="/create-job" class="text-white hover:text-main">Create your first job</a>
		{/if}
	</div>
</div>
