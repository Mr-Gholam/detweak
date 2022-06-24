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
	// @ts-nocheck

	import { page } from '$app/stores';

	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	import { goto } from '$app/navigation';
	export let job;
	export let title;
	console.log(job);
	let description = '';
	let descriptionErr;
	let applicationSent;
	let duplicateOffer;
	let jobId = $page.url.pathname.split('/')[2];
	async function submitApplication() {
		if (description.length > 15) {
			descriptionErr = undefined;
			const res = await fetch('/api/create-application', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					jobId,
					description
				})
			});
			if (res.ok) {
				description = undefined;
				applicationSent = true;
				setTimeout(() => {
					applicationSent = false;
				}, 3000);
			}
			if (res.status == 400) {
				duplicateOffer = 'You already applyed for this job';
			}
		} else {
			descriptionErr = 'Description must be at least 15 charecter';
		}
	}
	async function acceptOffer(offerId) {
		const box = document.querySelectorAll('.button-box');
		const offer = document.getElementById(`offer-${offerId}`);
		const res = await fetch('/api/reject-offer', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				offerId
			})
		});
		if (res.ok) {
			offer.classList.replace('border-border', 'border-green');
			box.forEach((el) => {
				el.classList.add('hidden');
			});
		}
	}
	async function rejectOffer(offerId) {
		const offerContainer = document.getElementById('offer-container');
		const offer = document.getElementById(`offer-${offerId}`);
		const res = await fetch('/api/reject-offer', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				offerId
			})
		});
		if (res.ok) {
			offer.classList.add('right-exit');
			setTimeout(() => {
				offerContainer.removeChild(offer);
			}, 1000);
		}
	}
	function sendMessage(username) {
		goto(`/messages?username=${username}`);
	}
</script>

<svelte:head>
	<title>{title ? title : 'job'}</title>
</svelte:head>
{#if job}
	<div class=" md:w-fit w-full lg:w-128 my-2 py-4">
		<div class="w-full lg:w-128  mx-auto h-full flex flex-col gap-4 pt-4">
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
				</div>
			</div>
			<!-- apply form -->
			{#if !job.isOwner}
				<div class="" id="application">
					<form on:submit|preventDefault={submitApplication}>
						<section class="w-full relative">
							<p class="text-sm text-error text-center my-2 {descriptionErr ? 'block' : 'hidden'}">
								{descriptionErr}
							</p>
							<p class="text-sm text-error text-center my-2 {duplicateOffer ? 'block' : 'hidden'}">
								{duplicateOffer}
							</p>
							<label class="text-base text-text " for="bio"
								>Describe why you are the best candidate</label
							>
							<textarea
								name="bio"
								id="bio"
								cols="20 "
								rows="5"
								placeholder="..."
								bind:value={description}
								class="outline-none border-2  text-text border-solid
								 rounded-lg px-2.5 mx-auto block w-11/12 my-3 p-2 resize-none {descriptionErr
									? 'border-error'
									: 'border-border'}"
							/>
						</section>
						<input type="submit" class="main-btn float-right mr-8" value="Apply" />
					</form>
				</div>
			{/if}
			{#if job.Offers}
				<div id="offer-container">
					{#each job.Offers as offer}
						<div
							id="offer-{offer.Id}"
							class="md:border-2 border-solid border-border  shadow-xl w-full rounded-md my-2 overflow-x-hidden p-2"
						>
							<!-- iformation -->
							<section
								class="flex gap-2 items-center justify-between  border-b border-border p-2 mb-2"
							>
								<a href="/profile/{offer.OwnerUsername}" class="ml-2">
									{#if offer.OwnerImgUrl}
										<!-- svelte-ignore a11y-img-redundant-alt -->
										<img
											class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
											src="/api/images/{offer.OwnerImgUrl}"
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
								<a href="/profile/{offer.OwnerUsername}" class="flex-1">
									<h4 class="mx-2 font-semibold text-text hover:text-text-hover">
										{offer.OwnerFirstname}
										{offer.OwnerLastname}
									</h4>
									<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
										@{offer.OwnerUsername}
									</h5>
								</a>
								<h6 class="text-xs text-text mx-2">
									{formatDistanceToNow(new Date(offer.CreatedAt), { addSuffix: true })}
								</h6>
							</section>
							<section>
								<h4 class="text-text mx-2 ">
									{offer.Description}
								</h4>
							</section>
							<section class="flex items-center mt-2">
								<button
									class=" main-btn text-sm border-1"
									on:click={sendMessage(offer.OwnerUsername)}
								>
									<i class="fa-solid fa-comments text-xs mr-2 " />contact for more information</button
								>
								<section
									class="flex justify-end w-10  items-center gap-3 text-lg flex-1 button-box"
								>
									<button
										on:click={rejectOffer(offer.Id)}
										class=" w-6 h-6 rounded-full text-center text-text  hover:text-error"
									>
										<i class="fa-solid fa-x" />
									</button>
									<button
										class=" w-7 h-7 rounded-full text-center text-text mt-1 hover:text-green"
										on:click={acceptOffer(offer.Id)}
									>
										<i class="fa-solid fa-check" />
									</button>
								</section>
							</section>
						</div>
					{/each}
				</div>
			{/if}
			{#if job.Others}
				<h3 class="text-text">People who also apply for this job</h3>
				<div class="flex">
					{#each job.Others as other}
						<a href="/profile/{other.OwnerUsername}" class="ml-2">
							{#if other.OwnerImgUrl}
								<!-- svelte-ignore a11y-img-redundant-alt -->
								<img
									class="h-12 w-12 object-cover rounded-full hover:opacity-90  "
									src="/api/images/{other.OwnerImgUrl}"
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
					{/each}
				</div>
			{/if}
		</div>
	</div>
{/if}
