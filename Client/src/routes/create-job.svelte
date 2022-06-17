<script>
	let languages = ['JavaScript', 'PHP', 'Java', 'C', 'C++', 'C#', 'Go', 'Python', 'Ruby'];
	let Fields = ['Front-end', 'Back-end', 'Full-stack', 'Game Developer'];
	let frontEndFrameWorks = ['React', 'Vue', 'Svelte', 'Anguler'];
	let projectname;
	let language;
	let frameWork;
	let deadline;
	let field;
	let budget;
	let title;
	let description;
	let passedtitle;
	let errortitle;
	let passedName;
	let errorName;
	let passedField;
	let errorField;
	let passedLanguage;
	let errorLanguage;
	let passedDeadline;
	let errorDeadline;
	let passedBudget;
	let errorBudget;
	let passedDes;
	let errorDes;
	function checkName() {
		if (projectname) {
			if (projectname.length > 2) {
				passedName = true;
				errorName = undefined;
			} else {
				errorName = 'Project name must be at least 3 characters';
				passedName = false;
			}
		} else {
			errorName = 'Project name must be at least 3 characters';
			passedName = false;
		}
	}
	function checkField() {
		if (field) {
			passedField = true;
			errorField = undefined;
		} else {
			passedField = false;
			errorField = 'Please select a Field';
		}
	}
	function checkLanguage() {
		if (language) {
			passedLanguage = true;
			errorLanguage = undefined;
		} else {
			passedLanguage = false;
			errorLanguage = 'Please select a language ';
		}
	}
	function checkDeadline() {
		if (deadline) {
			passedDeadline = true;
			errorDeadline = undefined;
		} else {
			passedDeadline = false;
			errorDeadline = 'Please select a deadline';
		}
	}
	function checkBudget() {
		if (budget) {
			passedBudget = true;
			errorBudget = undefined;
		} else {
			passedBudget = false;
			errorBudget = 'Please Enter your maximum budget';
		}
	}
	function checkDescription() {
		if (description) {
			passedDes = true;
			errorDes = undefined;
		} else {
			passedDes = false;
			errorDes = 'Description must be at least 15 characters';
		}
	}
	function checkTitle() {
		if (title) {
			if (title.length > 2) {
				passedtitle = true;
				errortitle = undefined;
			} else {
				passedtitle = undefined;
				errortitle = 'Title must be at least 3 characters';
			}
		} else {
			passedtitle = undefined;
			errortitle = 'Title must be at least 3 characters';
		}
	}
	async function submit() {
		checkName();
		checkField();
		checkLanguage();
		checkDeadline();
		checkBudget();
		checkDescription();
		checkTitle();
		if (
			passedBudget &&
			passedDeadline &&
			passedDes &&
			passedField &&
			passedLanguage &&
			passedName
		) {
			const res = await fetch('/api/create-job', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					title,
					ProjectName: projectname,
					language,
					field,
					frameWork: frameWork ? frameWork : undefined,
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
	}
</script>

<svelte:head>
	<title>Create a Job</title>
</svelte:head>

<div class="md:p-4 my-2 md:w-9/12">
	<div class="w-fit  mx-auto">
		<h4 class="text-white text-center text-xl  my-2">Create A Job</h4>
		<form
			action=""
			class=" flex-col gap-10 flex lg:flex-row  items-top"
			on:submit|preventDefault={submit}
		>
			<div class="flex-col gap-4 flex">
				<section class="w-80">
					<p class="text-error text-sm text-center {errortitle ? 'block' : 'hidden'}">
						{errortitle}
					</p>
					<label for="name" class="text-base text-text">Title</label>
					<input
						bind:value={title}
						type="text"
						on:change={checkTitle}
						class="outline-none border-2   rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text {errortitle
							? 'border-error'
							: 'border-border'}"
					/>
				</section>
				<section class="w-80">
					<p class="text-error text-sm text-center {errorName ? 'block' : 'hidden'}">{errorName}</p>
					<label for="name" class="text-base text-text">Project name </label>
					<input
						bind:value={projectname}
						type="text"
						on:change={checkName}
						class="outline-none border-2   rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text {errorName
							? 'border-error'
							: 'border-border'}"
					/>
				</section>
				<section class="w-80">
					<p class="text-error text-sm text-center {errorField ? 'block' : 'hidden'}">
						{errorField}
					</p>
					<label for="name" class="text-base text-text">Feild </label>
					<select
						name="framwork"
						id=""
						bind:value={field}
						on:change={checkField}
						class="block mx-auto w-11/12 text-center outline-none border-2 color-main  text-text border-solid rounded-lg  my-3 py-2.5 px-2.5 {errorField
							? 'border-error'
							: 'border-border'}"
					>
						<option value="" selected disabled />
						{#each Fields as Field}
							<option value={Field}>{Field}</option>
						{/each}
					</select>
				</section>
				<section class="w-80">
					<p class="text-error text-sm text-center {errorLanguage ? 'block' : 'hidden'}">
						{errorLanguage}
					</p>
					<label for="name" class="text-base text-text">Language </label>
					<select
						name="language"
						id=""
						bind:value={language}
						on:change={checkLanguage}
						class="block mx-auto w-11/12 text-center outline-none border-2 color-main  text-text border-solid rounded-lg  my-3 py-2.5 px-2.5 {errorLanguage
							? 'border-error'
							: 'border-border'}"
					>
						<option value="" selected disabled />
						{#each languages as language}
							<option value={language}>{language}</option>
						{/each}
					</select>
				</section>
				<section class="w-80">
					<label for="name" class="text-base text-text">FrameWork </label>
					<select
						name="framwork"
						id=""
						bind:value={frameWork}
						class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 py-2.5 px-2.5"
					>
						<option value="" selected disabled />
						{#each frontEndFrameWorks as frameWork}
							<option value={frameWork}>{frameWork}</option>
						{/each}
					</select>
				</section>
			</div>
			<div class="flex-col gap-4 flex">
				<section class="w-80">
					<p class="text-error text-sm text-center {errorDeadline ? 'block' : 'hidden'}">
						{errorDeadline}
					</p>
					<label for="name" class="text-base text-text">Deadline</label>
					<input
						bind:value={deadline}
						on:change={checkDeadline}
						type="date"
						class="outline-none border-2  border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text {errorDeadline
							? 'border-error'
							: 'border-border'}"
					/>
				</section>
				<section class="w-80 relative">
					<p class="text-error text-sm text-center {errorBudget ? 'block' : 'hidden'}">
						{errorBudget}
					</p>
					<label for="name" class="text-base text-text">Budget </label>
					<input
						bind:value={budget}
						on:change={checkBudget}
						type="number"
						class="outline-none border-2  border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text {errorBudget
							? 'border-error'
							: 'border-border'}"
					/>
					<p class="text-text absolute right-0 inset-y-1/2 mr-6 ">$</p>
				</section>
				<section class="w-80">
					<p class="text-error text-sm text-center {errorDes ? 'block' : 'hidden'}">
						{errorDes}
					</p>
					<label for="name" class="text-base text-text">Description </label>
					<textarea
						name="bio"
						id="bio"
						cols="20 "
						rows="5"
						bind:value={description}
						on:change={checkDescription}
						class="outline-none border-2  text-text border-solid
					 rounded-lg px-2.5 mx-auto block w-11/12 my-3 p-2 resize-none
					 {errorDes ? 'border-error' : 'border-border'}
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
			</div>
		</form>
	</div>
</div>
