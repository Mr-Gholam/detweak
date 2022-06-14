<script>
	let projectname;
	let language;
	let frameWork;
	let deadline;
	let budget;
	let description;
	async function submit() {
		const res = await fetch('/api/create-job', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				ProjectName: projectname,
				language,
				frameWork,
				deadline,
				budget,
				description
			})
		});
		if (res.ok) {
			projectname = '';
			language = '';
			frameWork = '';
			deadline = undefined;
			budget = '';
			description = '';
		}
	}
</script>

<svelte:head>
	<title>Create a Job</title>
</svelte:head>

<div class="md:p-4 my-2 md:w-9/12">
	<div class="w-fit mx-12">
		<h4 class="text-white text-center text-xl  my-2">Create A Job</h4>
		<form action="" class=" flex-col gap-4 flex" on:submit|preventDefault={submit}>
			<section class="w-80">
				<label for="name" class="text-base text-text">Project name </label>
				<input
					bind:value={projectname}
					type="text"
					class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
				/>
			</section>
			<section class="w-80">
				<label for="name" class="text-base text-text">Language </label>
				<input
					bind:value={language}
					type="text"
					class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
				/>
			</section>
			<section class="w-80">
				<label for="name" class="text-base text-text">FrameWork </label>
				<input
					bind:value={frameWork}
					type="text"
					class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
				/>
			</section>
			<section class="w-80">
				<label for="name" class="text-base text-text">Deadline</label>
				<input
					bind:value={deadline}
					type="date"
					class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
				/>
			</section>
			<section class="w-80 relative">
				<label for="name" class="text-base text-text">Budget </label>
				<input
					bind:value={budget}
					type="text"
					class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
				/>
				<p class="text-text absolute right-0 inset-y-1/2 mr-6 ">$</p>
			</section>
			<section class="w-80">
				<label for="name" class="text-base text-text">Description </label>
				<textarea
					name="bio"
					id="bio"
					cols="20 "
					rows="5"
					bind:value={description}
					class="outline-none border-2 border-border text-text border-solid
                 rounded-lg px-2.5 mx-auto block w-11/12 my-3 p-2 resize-none
                 "
				/>
			</section>
			<section class="w-80">
				<input
					id="submitPersonal"
					type="submit"
					value="Create"
					class="main-btn w-11/12 py-2 px-20 text-lg mx-auto block border-2"
				/>
			</section>
		</form>
	</div>
</div>
