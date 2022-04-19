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
	let email;
	let username;
	let password;
	let confirmPassword;
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
	function switchPersonal() {
		const personal = document.getElementById('personal');
		const account = document.getElementById('account');
		const personalBtn = document.getElementById('personalBtn');
		const accountBtn = document.getElementById('accountBtn');
		personal.classList.remove('hidden');
		account.classList.add('hidden');
		personalBtn.classList.add('text-main');
		accountBtn.classList.remove('text-main');
	}
	function switchAccount() {
		const personal = document.getElementById('personal');
		const account = document.getElementById('account');
		const personalBtn = document.getElementById('personalBtn');
		const accountBtn = document.getElementById('accountBtn');
		personal.classList.add('hidden');
		account.classList.remove('hidden');
		personalBtn.classList.remove('text-main');
		accountBtn.classList.add('text-main');
	}
</script>

<div class="md:p-4 my-2 md:w-9/12 ">
	<section
		class="w-80 flex border-2 justify-evenly p-3 bg-main-bg rounded-full text-white mx-auto mb-4 "
	>
		<button class="hover:text-main text-main" id="personalBtn" on:click={switchPersonal}
			>Personal Setting</button
		>
		<button class="hover:text-main" id="accountBtn" on:click={switchAccount}>Account Setting</button
		>
	</section>
	<!--Personal-->
	<form
		action="/set-profile"
		method="post"
		class="flex flex-col justify-between  p-4 gap-4 items-center shadow-xl border-2 w-fit mx-auto"
		enctype="multipart/form-data"
		id="personal"
	>
		<h1 class="text-lg font-semibold text-center">Personal Setting</h1>
		<section class="w-80">
			<label class="text-base" for="firstName">First Name</label>
			<input
				type="text"
				name="firstName"
				id="firstName"
				bind:value={firstName}
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
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3"
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
            file:text-white
            file:bg-main-bg
            file:border-main-bg file:border-2
            file:border-solid
            file:rounded-full
            file:hover:text-main
    
            file:hover:cursor-pointer"
				/>
			</div>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Save Changes"
				class="cursor-pointer text-lg rounded-lg hover:text-main text-white bg-main-bg py-2 px-20 mx-auto block w-11/12 hover:shadow-xl hover:bg-main-darker  "
			/>
		</section>
	</form>
	<!--auth-->
	<!-- Making it  switchable -->
	<form
		action="/signup"
		method="post"
		class="flex flex-col justify-between  p-4 gap-3 items-center shadow-xl border-2 hidden w-fit mx-auto"
		id="account"
	>
		<h1 class="text-lg font-semibold text-center">Account Settting</h1>
		<section class="w-80 ">
			<label class="text-base" for="email">Email</label>
			<input
				bind:value={email}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base" for="username">Username</label>
			<input
				bind:value={username}
				type="text"
				name="username"
				id="username"
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base" for="password">Password</label>
			<input
				bind:value={password}
				type="password"
				name="password"
				id="password"
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base" for="confirmPassword">Confirm Password</label>
			<input
				bind:value={confirmPassword}
				type="password"
				name="confirmPassword"
				id="confirmPassword"
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Save Changes"
				class="cursor-pointer text-lg rounded-lg  bg-main-bg hover:text-main py-2 px-20 mx-auto block w-11/12 hover:shadow-xl text-white "
			/>
		</section>
	</form>
</div>
