<script context="module">
	export const load = async ({ fetch, params }) => {
		const jobId = params.jobId;
		const res = await fetch(`/api/job/${jobId}`);
		let job;
		let title;
		if (res.status == 200) {
			job = await res.json();
			title = job.Title;
		}
		return { props: { job, title } };
	};
</script>

<script>
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	export let job;
	export let title;
	console.log(job);
</script>

<svelte:head>
	<title>{title ? title : 'job'}</title>
</svelte:head>
{#if job}
	<div class=" md:w-fit w-full lg:w-128 my-2 py-4">
		<div class="w-full lg:w-128  mx-auto">
			<div
				class="md:border-2 border-solid border-border  shadow-xl w-full rounded-md my-2 overflow-x-hidden p-2"
			>
				<section class="flex gap-2 items-center justify-between  border-b border-border p-2 mb-2">
					<!--profile img-->
					<a href="/profile/{job.OwnerUsername}" class="ml-2">
						{#if job.OwnerImgUrl}
							<!-- svelte-ignore a11y-img-redundant-alt -->
							<img
								class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
								src="/api/images/{job.OwnerImgUrl}"
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
					<a href="/profile/{job.OwnerUsername}" class="flex-1">
						<h4 class="mx-2 font-semibold text-text hover:text-text-hover">
							{job.OwnerFirstname}
							{job.OwnerLastname}
						</h4>
						<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
							@{job.OwnerUsername}
						</h5>
					</a>
					<h6 class="text-xs text-text mx-2">
						{formatDistanceToNow(new Date(job.CreatedAt), { addSuffix: true })}
					</h6>
				</section>
				<div class="flex justify-between items-end ">
					<section>
						<h4 class="mx-2 font-bold text-text ">
							{job.Title}
						</h4>
						<h5 class=" text-sm  text-text  mx-2 my-1 flex">
							<p class="font-bold mr-1">Project name :</p>
							{job.ProjectName}
						</h5>
						<h5 class=" text-sm  text-text  m-2 flex ">
							<p class="font-bold mr-1">Field :</p>
							{job.Field}
						</h5>
						<h5 class=" text-sm  text-text  m-2 flex ">
							<p class="font-bold mr-1">Requirements :</p>
							{job.Language}
							{#if job.FrameWork} - {job.FrameWork}{/if}
						</h5>
						<h5 class=" text-sm  text-text  mx-2 flex ">
							<p class="font-bold mr-1">Budget :</p>
							{job.Budget}$
						</h5>
						<h5 class=" text-sm  text-text  m-2  flex">
							<p class="font-bold mr-1">Deadline :</p>
							{job.Deadline}
						</h5>
						<h5 class=" text-sm  text-text  m-2  flex">
							<p class="font-bold mr-1">Description :</p>
							{job.Description}
						</h5>
					</section>
					{#if job.isOwner}
						<div class="mx-2 my-2">
							<button class="main-btn">Apply</button>
						</div>
					{/if}
				</div>
			</div>
		</div>
	</div>
{/if}
