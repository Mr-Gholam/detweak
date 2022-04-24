<script>
	let firstName;
	let lastName;
	let bio;
	let profilePic;
	let profilePicInput;
	let timelineSelected;
	let amPm;
	let timelines = [
		{ id: 0, value: '0-1' },
		{ id: 1, value: '1-2' },
		{ id: 2, value: '2-3' },
		{ id: 3, value: '3-4' },
		{ id: 4, value: '4-5' },
		{ id: 5, value: '5-6' },
		{ id: 6, value: '6-7' },
		{ id: 7, value: '7-8' },
		{ id: 8, value: '8-9' },
		{ id: 9, value: '9-10' },
		{ id: 10, value: '10-11' },
		{ id: 11, value: '11-12' }
	];
	let passedName = false;
	let passedLastName = false;
	let passedBio = false;
	//preview  profile picture
	function imageChange() {
		const file = profilePicInput.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', () => {
				profilePic.setAttribute('src', reader.result);
			});
			reader.readAsDataURL(file);
			return;
		}
	}
	// checking firstName
	function checkFirstName() {
		const el = document.getElementById('firstName');
		const error = document.createElement('p');
		error.classList.add('error');
		error.innerHTML = 'First name should be at least 3 charecters';
		if (firstName.length <= 3) {
			el.classList.add('border-error');
			if (el.parentNode.querySelector('.error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedName = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.error'));
			}
			passedName = true;
		}
	}
	// checking lastName
	function checkLastName() {
		const el = document.getElementById('lastName');
		const error = document.createElement('p');
		error.classList.add('error');
		error.innerHTML = 'Last name should be at least 3 charecters';
		if (lastName.length <= 3) {
			el.classList.add('border-error');
			if (el.parentNode.querySelector('.text-error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedLastName = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.error'));
			}
			passedLastName = true;
		}
	}
	//Checking Bio
	function checkBio() {
		const el = document.getElementById('bio');
		const error = document.createElement('p');
		error.classList.add('error');
		error.innerHTML = 'Bio should be at least 15 charecters';
		if (bio.length <= 15) {
			el.classList.add('border-error');
			if (el.parentNode.querySelector('.error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedBio = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.error'));
			}
			passedBio = true;
		}
	}
	/**
	 * ? ya ali madad
	 */
	//submit
	function submit() {
		if (passedName) {
			if (timelineSelected) {
				if (amPm) {
					if (profilePicInput.file) {
					} else {
						document
							.getElementById('profileImg')
							.classList.add('file:border-error', 'file:text-error');
					}
				} else {
					document.getElementById('am-pm').classList.add('border-error');
				}
			} else {
				document.getElementById('online-time').classList.add('border-error');
			}
		}
	}
	//
</script>

<main class=" md:flex md:justify-center md:m-auto md:py-8  md:items-center">
	<form
		on:submit|preventDefault={submit}
		action="/set-profile"
		method="post"
		class="flex flex-col justify-between lg:mr-56  p-4 gap-4 items-center"
		enctype="multipart/form-data"
	>
		<section class="w-80">
			<label class="text-base" for="firstName">First Name</label>
			<input
				type="text"
				name="firstName"
				id="firstName"
				bind:value={firstName}
				on:change={checkFirstName}
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80">
			<label class="text-base" for="lastName">Last Name</label>

			<input
				type="text"
				name="lastName"
				id="lastName"
				bind:value={lastName}
				on:change={checkLastName}
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3 "
			/>
		</section>
		<section class="w-80">
			<label class="text-base" for="bio">Bio</label>
			<textarea
				name="bio"
				id="bio"
				cols="20 "
				rows="5"
				bind:value={bio}
				on:change={checkBio}
				class="outline-none border-2 border-main-bg border-solid
                 rounded-lg px-2.5 mx-auto block w-11/12 my-3 p-2 resize-none
                 "
			/>
		</section>
		<section class="w-80">
			<label for="online-time" class="text-base">Choose daily online time</label>
			<section class="flex justify-center">
				<select
					id="online-time"
					bind:value={timelineSelected}
					class="block mx-2 w-24 text-center outline-none border-2 color-main border-main-bg border-solid rounded-lg p-1 my-3"
				>
					<option selected disabled value=""> hour</option>
					{#each timelines as timeline}
						<option value={timeline.value}>
							{timeline.value}
						</option>
					{/each}
				</select>
				<select
					bind:value={amPm}
					name="am-pm"
					id="am-pm"
					class="block mx-2 w-24 text-center outline-none border-2 color-main border-main-bg border-solid rounded-lg p-1 my-3"
				>
					<option selected disabled value=""> Am-PM</option>
					<option value="AM">AM</option>
					<option value="PM">PM</option>
				</select>
			</section>
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
					id="profileImg"
					accept="image/png, image/jpeg"
					class="block w-4/6 text-sm text-gray-500
                file:mr-3 file:py-2 file:px-3 
                file:text-sm file:font-semibold
                file:bg-violet-50 file:text-main-bg
                file:border-main-bg file:border-2
                file:border-solid
                file:rounded-full
                file:hover:bg-main
		
                file:hover:cursor-pointer"
				/>
			</div>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Next"
				class="cursor-pointer text-lg rounded-lg hover:text-main text-white bg-main-bg py-2 px-20 mx-auto block w-11/12 hover:shadow-xl hover:bg-main-darker  "
			/>
		</section>
	</form>
	<h1 class="md:ml-56  md:p-4 hidden">There will be pic here</h1>
</main>
