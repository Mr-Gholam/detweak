<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch('/api/my-offers');
		let offers;
		if (res.status == 200) {
			offers = await res.json();
		}
		return { props: { offers } };
	};
</script>

<script>
	export let offers;
	console.log(offers);
</script>

<svelte:head>
	<title>My Offers</title>
</svelte:head>

<div class=" md:w-fit w-full lg:w-128 my-2 py-4">
	{#if offers}
		{#each offers as offer}
			<div
				class="md:border-2 border-solid border-border  shadow-xl w-full rounded-md my-2 overflow-x-hidden p-2"
			>
				<section class="flex gap-2 items-center justify-between">
					<!-- Name and username-->
					<a href="/job/{offer.JobId}" class="flex-1">
						<h4 class="mx-2 font-semibold text-text hover:text-text-hover">
							{offer.Title}
						</h4>
						<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
							{offer.ProjectName}
						</h5>
						<h5 class=" text-sm  text-text hover:text-text-hover mx-2 ">
							Status : <p
								class="inline {offer.AcceptedSeen
									? 'text-green'
									: offer.RejectedSeen
									? 'text-error'
									: 'text-text'} 
                                    hover:text-text-hover"
							>
								{offer.Status}
							</p>
						</h5>
					</a>
					<section class="flex flex-col justify-end items-center relative ">
						<a
							href="/job/{offer.JobId}"
							class="border-border p-2 border-2 rounded-full text-sm text-text hover:border-main hover:text-white"
							>See More</a
						>
					</section>
				</section>
			</div>
		{/each}
	{:else}
		<div class="flex flex-col gap-4">
			<p class="text-text text-center">You did not place any offers</p>
			<a
				href="/find-job"
				class="text-center w-fit mx-auto border-border md:border-2 py-3 px-5 rounded-full text-sm  text-text hover:text-main hover:border-main "
			>
				Place offers on jobs
			</a>
		</div>
	{/if}
</div>
