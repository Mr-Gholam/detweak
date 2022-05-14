<script>
	let languages = ['JavaScript', 'PHP', 'Java', 'C', 'C++', 'C#', 'Go', 'Python', 'Ruby'];
	let feilds = ['Front-end ', 'Back-end', 'Full-stack', 'Game Developer'];
	let frontEndFrameWorks = ['React', 'Vue', 'Svelte', 'Anguler'];
</script>

<div>
	<form action="" class="flex flex-col justify-between lg:mr-56  p-4 gap-4 items-center">
		<section class="w-80">
			<label for="felid" class="text-base text-text">Your Github username </label>
			<input
				type="text"
				class="outline-none border-2 border-border text-text border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80">
			<label for="felid" class="text-base text-text">Choose your feild</label>
			<select
				name=""
				id=""
				class="block mx-2 w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each feilds as feild}
					<option value={feild}>{feild}</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<label for="languages" class="text-base text-text">Choose your main language</label>
			<select
				name=""
				id=""
				class="block mx-2 w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each languages as language}
					<option value="language">{language}</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<label for="fram-work" class="text-base text-text">Choose your fram work</label>
			<select
				name=""
				id=""
				class="block mx-2 w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-1 px-2.5"
			>
				{#each frontEndFrameWorks as frameWork}
					<option value={frameWork}>{frameWork}</option>
				{/each}
			</select>
		</section>
	</form>
</div>
