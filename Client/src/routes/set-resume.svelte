<script>
	import { goto } from '$app/navigation';

	let selectedLanguage = null;
	let selectedField = null;
	let selectedFrameWork = null;
	let gitHubUserName = '';
	let experience = null;
	let languages = ['JavaScript', 'PHP', 'Java', 'C', 'C++', 'C#', 'Go', 'Python', 'Ruby'];
	let exOption = [
		'Less than 6 months',
		'Less than 1 year',
		'Between 1-2 years',
		'Between 2-4 years',
		'More than 4 years'
	];
	let Fields = ['Front-end', 'Back-end', 'Full-stack', 'Game Developer'];
	let frontEndFrameWorks = ['React', 'Vue', 'Svelte', 'Anguler'];
	async function submit() {
		if (
			selectedField == null &&
			selectedFrameWork == null &&
			gitHubUserName == '' &&
			selectedLanguage == null
		) {
			return goto('/dashboard');
		}
		const response = await fetch('/api/set-resume', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				githubUsername: gitHubUserName,
				language: selectedLanguage,
				field: selectedField,
				frameWork: selectedFrameWork,
				experience: experience
			})
		});
		if (response.ok) {
			goto('/dashboard');
		}
	}
</script>

<svelte:head>
	<title>Set Resume</title>
</svelte:head>

<div class=" md:flex md:justify-center md:m-auto md:py-8  md:items-center">
	<form
		action=""
		class="flex flex-col justify-between  p-4 gap-4 items-center"
		on:submit|preventDefault={submit}
	>
		<section class="w-80">
			<h4 class="text-white text-xl text-center my-2">Perfessional Information</h4>
		</section>
		<section class="w-80">
			<label for="felid" class="text-base text-text">Your Github username </label>
			<input
				bind:value={gitHubUserName}
				type="text"
				class="outline-none border-2 border-border text-text border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80">
			<label for="felid" class="text-base text-text">Coding Experience </label>
			<select
				bind:value={experience}
				name=""
				id=""
				class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each exOption as xp}
					<option value={xp}>{xp}</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<label for="felid" class="text-base text-text">Choose your Field</label>
			<select
				bind:value={selectedField}
				name=""
				id=""
				class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each Fields as field}
					<option value={field}>{field}</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<label for="languages" class="text-base text-text">Choose your main language</label>
			<select
				bind:value={selectedLanguage}
				name=""
				id=""
				class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each languages as language}
					<option value={language}>{language}</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<label for="fram-work" class="text-base text-text">Choose your fram work</label>
			<select
				bind:value={selectedFrameWork}
				name=""
				id=""
				class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each frontEndFrameWorks as frameWork}
					<option value={frameWork}>{frameWork}</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Next"
				class=" text-lg py-2 px-20 mx-auto block w-11/12 hover:shadow-xl main-btn my-3 border-2  "
			/>
		</section>
	</form>
</div>
