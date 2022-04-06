<script>
	let firstName;
	let lastName;
	let bio;
	let profilePic;
	let profilePicInput;
	let timelineSelected;
	let timelines = [
		{ id: 1, value: '1 - 2' },
		{ id: 2, value: '2 - 3' },
		{ id: 3, value: '3 - 4' },
		{ id: 4, value: '4 - 5' },
		{ id: 5, value: '5 - 6' },
		{ id: 6, value: '6 - 7' },
		{ id: 7, value: '7 - 8' },
		{ id: 8, value: '8 - 9' },
		{ id: 9, value: '9 - 10' },
		{ id: 10, value: '10 - 11' },
		{ id: 11, value: '11 - 12' },
		{ id: 12, value: '12 - 13' },
		{ id: 13, value: '13 - 14' },
		{ id: 14, value: '14 - 15' },
		{ id: 15, value: '15 - 16' },
		{ id: 16, value: '16 - 17' },
		{ id: 17, value: '17 - 18' },
		{ id: 18, value: '18 - 19' },
		{ id: 19, value: '19 - 20' },
		{ id: 20, value: '20 - 21' },
		{ id: 21, value: '21 - 22' },
		{ id: 22, value: '22 - 21' },
		{ id: 23, value: '23 - 24' }
	];
	//preview  profile picture
	function imageChange() {
		const file = profilePicInput.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', () => {
				profilePic.setAttribute('src', reader.result);
			});
			reader.readAsDataURL(file);
			console.log(file);
			return;
		}
	}
	// checking firstName
	/**
	 * todo find a way to use error golbaly
	 */
	function firstNameCheck() {
		const el = document.getElementById('firstName');
		const error = document.createElement('p');
		error.classList.add('text-error', 'text-center', 'text-base', 'my-2');
		error.innerHTML = 'your first name should be at least 3 charecters';
		if (firstName.length <= 3) {
			el.classList.add('border-error');
			if (el.parentNode.querySelector('.text-error')) {
				return;
			}
			el.parentNode.insertBefore(error, el.previousElementSibling);
		} else {
			el.classList.remove('border-error');
			el.closest('section').removeChild(el.parentNode.querySelector('.text-error'));
		}
	}
	// checking lastName
	function lastNameCheck() {
		if (lastName.length <= 3) {
			document.getElementById('lastName').classList.add('border-error');
		} else {
			document.getElementById('lastName').classList.remove('border-error');
		}
	}
</script>

<main class=" flex justify-center m-auto py-8  items-center">
	<form
		action="/set-profile"
		method="post"
		class="flex flex-col justify-between mr-56  p-4 gap-4"
		enctype="multipart/form-data"
	>
		<section class="w-80">
			<label class="text-base" for="firstName">First Name</label>
			<input
				type="text"
				name="firstName"
				id="firstName"
				bind:value={firstName}
				on:change={firstNameCheck}
				class="outline-none border-2 border-main border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80">
			<label class="text-base" for="lastName">Last Name</label>

			<input
				type="text"
				name="lastName"
				id="lastName"
				bind:value={lastName}
				on:change={lastNameCheck}
				class="outline-none border-2 border-main border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80">
			<label class="text-base" for="bio">Bio</label>
			<textarea
				name="bio"
				id=""
				cols="20 "
				rows="5"
				bind:value={bio}
				class="outline-none border-2 border-main border-solid
                 rounded-lg px-2.5 mx-auto block w-11/12 my-3 p-2 resize-none
                 "
			/>
		</section>
		<section class="w-80">
			<label for="online-time" class="text-base">Choose daily online time</label>

			<select
				id="online-time"
				bind:value={timelineSelected}
				class="block mx-auto w-20 text-center outline-none border-2 color-main border-main border-solid rounded-lg p-1 my-3"
			>
				{#each timelines as timeline}
					<option value={timeline.value}>
						{timeline.value}
					</option>
				{/each}
			</select>
		</section>
		<section class="w-80">
			<label class="text-base" for="profilePic">Profile Picture</label>
			<div class="flex items-center justify-evenly my-3   ">
				<!-- svelte-ignore a11y-img-redundant-alt -->
				<img
					class="h-16 w-16 object-cover rounded-full "
					bind:this={profilePic}
					src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
					alt="Current profile photo"
				/>
				<input
					type="file"
					name="profilePic"
					bind:this={profilePicInput}
					on:change={imageChange}
					id=""
					accept="image/png, image/jpeg"
					class="block w-4/6 text-sm text-gray-500
                file:mr-3 file:py-2 file:px-3 
                file:text-sm file:font-semibold
                file:bg-violet-50 file:text-main
                file:border-main file:border-2
                file:border-solid
                file:rounded-full
                hover:file:bg-violet-100
                file:hover:cursor-pointer"
				/>
			</div>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Next"
				class="cursor-pointer text-lg rounded-lg text-white bg-main py-2 px-20 mx-auto block w-11/12 hover:shadow-xl hover:bg-main-darker  "
			/>
		</section>
	</form>
	<h1 class="ml-56  p-4 ">There will be pic here</h1>
</main>
