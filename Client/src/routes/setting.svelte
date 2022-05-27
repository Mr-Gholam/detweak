<script>
	import { onMount } from 'svelte';
	import { loading } from '../store';
	let user;
	let firstName;
	let lastName;
	let bio;
	let profilePic;
	let profilePicInput;
	let countries = [
		{ text: 'Afghanistan', value: 'AF' },
		{ text: 'Åland Islands', value: 'AX' },
		{ text: 'Albania', value: 'AL' },
		{ text: 'Algeria', value: 'DZ' },
		{ text: 'American Samoa', value: 'AS' },
		{ text: 'Andorra', value: 'AD' },
		{ text: 'Angola', value: 'AO' },
		{ text: 'Anguilla', value: 'AI' },
		{ text: 'Antarctica', value: 'AQ' },
		{ text: 'Antigua and Barbuda', value: 'AG' },
		{ text: 'Argentina', value: 'AR' },
		{ text: 'Armenia', value: 'AM' },
		{ text: 'Aruba', value: 'AW' },
		{ text: 'Australia', value: 'AU' },
		{ text: 'Austria', value: 'AT' },
		{ text: 'Azerbaijan', value: 'AZ' },
		{ text: 'Bahamas', value: 'BS' },
		{ text: 'Bahrain', value: 'BH' },
		{ text: 'Bangladesh', value: 'BD' },
		{ text: 'Barbados', value: 'BB' },
		{ text: 'Belarus', value: 'BY' },
		{ text: 'Belgium', value: 'BE' },
		{ text: 'Belize', value: 'BZ' },
		{ text: 'Benin', value: 'BJ' },
		{ text: 'Bermuda', value: 'BM' },
		{ text: 'Bhutan', value: 'BT' },
		{ text: 'Bolivia', value: 'BO' },
		{ text: 'Bosnia and Herzegovina', value: 'BA' },
		{ text: 'Botswana', value: 'BW' },
		{ text: 'Bouvet Island', value: 'BV' },
		{ text: 'Brazil', value: 'BR' },
		{ text: 'British Indian Ocean Territory', value: 'IO' },
		{ text: 'Brunei Darussalam', value: 'BN' },
		{ text: 'Bulgaria', value: 'BG' },
		{ text: 'Burkina Faso', value: 'BF' },
		{ text: 'Burundi', value: 'BI' },
		{ text: 'Cambodia', value: 'KH' },
		{ text: 'Cameroon', value: 'CM' },
		{ text: 'Canada', value: 'CA' },
		{ text: 'Cape Verde', value: 'CV' },
		{ text: 'Cayman Islands', value: 'KY' },
		{ text: 'Central African Republic', value: 'CF' },
		{ text: 'Chad', value: 'TD' },
		{ text: 'Chile', value: 'CL' },
		{ text: 'China', value: 'CN' },
		{ text: 'Christmas Island', value: 'CX' },
		{ text: 'Cocos (Keeling) Islands', value: 'CC' },
		{ text: 'Colombia', value: 'CO' },
		{ text: 'Comoros', value: 'KM' },
		{ text: 'Congo', value: 'CG' },
		{ text: 'Congo, The Democratic Republic of the', value: 'CD' },
		{ text: 'Cook Islands', value: 'CK' },
		{ text: 'Costa Rica', value: 'CR' },
		{ text: "Cote D'Ivoire", value: 'CI' },
		{ text: 'Croatia', value: 'HR' },
		{ text: 'Cuba', value: 'CU' },
		{ text: 'Cyprus', value: 'CY' },
		{ text: 'Czech Republic', value: 'CZ' },
		{ text: 'Denmark', value: 'DK' },
		{ text: 'Djibouti', value: 'DJ' },
		{ text: 'Dominica', value: 'DM' },
		{ text: 'Dominican Republic', value: 'DO' },
		{ text: 'Ecuador', value: 'EC' },
		{ text: 'Egypt', value: 'EG' },
		{ text: 'El Salvador', value: 'SV' },
		{ text: 'Equatorial Guinea', value: 'GQ' },
		{ text: 'Eritrea', value: 'ER' },
		{ text: 'Estonia', value: 'EE' },
		{ text: 'Ethiopia', value: 'ET' },
		{ text: 'Falkland Islands (Malvinas)', value: 'FK' },
		{ text: 'Faroe Islands', value: 'FO' },
		{ text: 'Fiji', value: 'FJ' },
		{ text: 'Finland', value: 'FI' },
		{ text: 'France', value: 'FR' },
		{ text: 'French Guiana', value: 'GF' },
		{ text: 'French Polynesia', value: 'PF' },
		{ text: 'French Southern Territories', value: 'TF' },
		{ text: 'Gabon', value: 'GA' },
		{ text: 'Gambia', value: 'GM' },
		{ text: 'Georgia', value: 'GE' },
		{ text: 'Germany', value: 'DE' },
		{ text: 'Ghana', value: 'GH' },
		{ text: 'Gibraltar', value: 'GI' },
		{ text: 'Greece', value: 'GR' },
		{ text: 'Greenland', value: 'GL' },
		{ text: 'Grenada', value: 'GD' },
		{ text: 'Guadeloupe', value: 'GP' },
		{ text: 'Guam', value: 'GU' },
		{ text: 'Guatemala', value: 'GT' },
		{ text: 'Guernsey', value: 'GG' },
		{ text: 'Guinea', value: 'GN' },
		{ text: 'Guinea-Bissau', value: 'GW' },
		{ text: 'Guyana', value: 'GY' },
		{ text: 'Haiti', value: 'HT' },
		{ text: 'Heard Island and Mcdonald Islands', value: 'HM' },
		{ text: 'Holy See (Vatican City State)', value: 'VA' },
		{ text: 'Honduras', value: 'HN' },
		{ text: 'Hong Kong', value: 'HK' },
		{ text: 'Hungary', value: 'HU' },
		{ text: 'Iceland', value: 'IS' },
		{ text: 'India', value: 'IN' },
		{ text: 'Indonesia', value: 'ID' },
		{ text: 'Iran, Islamic Republic Of', value: 'IR' },
		{ text: 'Iraq', value: 'IQ' },
		{ text: 'Ireland', value: 'IE' },
		{ text: 'Isle of Man', value: 'IM' },
		{ text: 'Israel', value: 'IL' },
		{ text: 'Italy', value: 'IT' },
		{ text: 'Jamaica', value: 'JM' },
		{ text: 'Japan', value: 'JP' },
		{ text: 'Jersey', value: 'JE' },
		{ text: 'Jordan', value: 'JO' },
		{ text: 'Kazakhstan', value: 'KZ' },
		{ text: 'Kenya', value: 'KE' },
		{ text: 'Kiribati', value: 'KI' },
		{ text: "Korea, Democratic People'S Republic of", value: 'KP' },
		{ text: 'Korea, Republic of', value: 'KR' },
		{ text: 'Kuwait', value: 'KW' },
		{ text: 'Kyrgyzstan', value: 'KG' },
		{ text: "Lao People'S Democratic Republic", value: 'LA' },
		{ text: 'Latvia', value: 'LV' },
		{ text: 'Lebanon', value: 'LB' },
		{ text: 'Lesotho', value: 'LS' },
		{ text: 'Liberia', value: 'LR' },
		{ text: 'Libyan Arab Jamahiriya', value: 'LY' },
		{ text: 'Liechtenstein', value: 'LI' },
		{ text: 'Lithuania', value: 'LT' },
		{ text: 'Luxembourg', value: 'LU' },
		{ text: 'Macao', value: 'MO' },
		{ text: 'Macedonia, The Former Yugoslav Republic of', value: 'MK' },
		{ text: 'Madagascar', value: 'MG' },
		{ text: 'Malawi', value: 'MW' },
		{ text: 'Malaysia', value: 'MY' },
		{ text: 'Maldives', value: 'MV' },
		{ text: 'Mali', value: 'ML' },
		{ text: 'Malta', value: 'MT' },
		{ text: 'Marshall Islands', value: 'MH' },
		{ text: 'Martinique', value: 'MQ' },
		{ text: 'Mauritania', value: 'MR' },
		{ text: 'Mauritius', value: 'MU' },
		{ text: 'Mayotte', value: 'YT' },
		{ text: 'Mexico', value: 'MX' },
		{ text: 'Micronesia, Federated States of', value: 'FM' },
		{ text: 'Moldova, Republic of', value: 'MD' },
		{ text: 'Monaco', value: 'MC' },
		{ text: 'Mongolia', value: 'MN' },
		{ text: 'Montserrat', value: 'MS' },
		{ text: 'Morocco', value: 'MA' },
		{ text: 'Mozambique', value: 'MZ' },
		{ text: 'Myanmar', value: 'MM' },
		{ text: 'Namibia', value: 'NA' },
		{ text: 'Nauru', value: 'NR' },
		{ text: 'Nepal', value: 'NP' },
		{ text: 'Netherlands', value: 'NL' },
		{ text: 'Netherlands Antilles', value: 'AN' },
		{ text: 'New Caledonia', value: 'NC' },
		{ text: 'New Zealand', value: 'NZ' },
		{ text: 'Nicaragua', value: 'NI' },
		{ text: 'Niger', value: 'NE' },
		{ text: 'Nigeria', value: 'NG' },
		{ text: 'Niue', value: 'NU' },
		{ text: 'Norfolk Island', value: 'NF' },
		{ text: 'Northern Mariana Islands', value: 'MP' },
		{ text: 'Norway', value: 'NO' },
		{ text: 'Oman', value: 'OM' },
		{ text: 'Pakistan', value: 'PK' },
		{ text: 'Palau', value: 'PW' },
		{ text: 'Palestinian Territory, Occupied', value: 'PS' },
		{ text: 'Panama', value: 'PA' },
		{ text: 'Papua New Guinea', value: 'PG' },
		{ text: 'Paraguay', value: 'PY' },
		{ text: 'Peru', value: 'PE' },
		{ text: 'Philippines', value: 'PH' },
		{ text: 'Pitcairn', value: 'PN' },
		{ text: 'Poland', value: 'PL' },
		{ text: 'Portugal', value: 'PT' },
		{ text: 'Puerto Rico', value: 'PR' },
		{ text: 'Qatar', value: 'QA' },
		{ text: 'Reunion', value: 'RE' },
		{ text: 'Romania', value: 'RO' },
		{ text: 'Russian Federation', value: 'RU' },
		{ text: 'RWANDA', value: 'RW' },
		{ text: 'Saint Helena', value: 'SH' },
		{ text: 'Saint Kitts and Nevis', value: 'KN' },
		{ text: 'Saint Lucia', value: 'LC' },
		{ text: 'Saint Pierre and Miquelon', value: 'PM' },
		{ text: 'Saint Vincent and the Grenadines', value: 'VC' },
		{ text: 'Samoa', value: 'WS' },
		{ text: 'San Marino', value: 'SM' },
		{ text: 'Sao Tome and Principe', value: 'ST' },
		{ text: 'Saudi Arabia', value: 'SA' },
		{ text: 'Senegal', value: 'SN' },
		{ text: 'Serbia and Montenegro', value: 'CS' },
		{ text: 'Seychelles', value: 'SC' },
		{ text: 'Sierra Leone', value: 'SL' },
		{ text: 'Singapore', value: 'SG' },
		{ text: 'Slovakia', value: 'SK' },
		{ text: 'Slovenia', value: 'SI' },
		{ text: 'Solomon Islands', value: 'SB' },
		{ text: 'Somalia', value: 'SO' },
		{ text: 'South Africa', value: 'ZA' },
		{ text: 'South Georgia and the South Sandwich Islands', value: 'GS' },
		{ text: 'Spain', value: 'ES' },
		{ text: 'Sri Lanka', value: 'LK' },
		{ text: 'Sudan', value: 'SD' },
		{ text: 'Suriname', value: 'SR' },
		{ text: 'Svalbard and Jan Mayen', value: 'SJ' },
		{ text: 'Swaziland', value: 'SZ' },
		{ text: 'Sweden', value: 'SE' },
		{ text: 'Switzerland', value: 'CH' },
		{ text: 'Syrian Arab Republic', value: 'SY' },
		{ text: 'Taiwan, Province of China', value: 'TW' },
		{ text: 'Tajikistan', value: 'TJ' },
		{ text: 'Tanzania, United Republic of', value: 'TZ' },
		{ text: 'Thailand', value: 'TH' },
		{ text: 'Timor-Leste', value: 'TL' },
		{ text: 'Togo', value: 'TG' },
		{ text: 'Tokelau', value: 'TK' },
		{ text: 'Tonga', value: 'TO' },
		{ text: 'Trinidad and Tobago', value: 'TT' },
		{ text: 'Tunisia', value: 'TN' },
		{ text: 'Turkey', value: 'TR' },
		{ text: 'Turkmenistan', value: 'TM' },
		{ text: 'Turks and Caicos Islands', value: 'TC' },
		{ text: 'Tuvalu', value: 'TV' },
		{ text: 'Uganda', value: 'UG' },
		{ text: 'Ukraine', value: 'UA' },
		{ text: 'United Arab Emirates', value: 'AE' },
		{ text: 'United Kingdom', value: 'GB' },
		{ text: 'United States', value: 'US' },
		{ text: 'United States Minor Outlying Islands', value: 'UM' },
		{ text: 'Uruguay', value: 'UY' },
		{ text: 'Uzbekistan', value: 'UZ' },
		{ text: 'Vanuatu', value: 'VU' },
		{ text: 'Venezuela', value: 'VE' },
		{ text: 'Viet Nam', value: 'VN' },
		{ text: 'Virgin Islands, British', value: 'VG' },
		{ text: 'Virgin Islands, U.S.', value: 'VI' },
		{ text: 'Wallis and Futuna', value: 'WF' },
		{ text: 'Western Sahara', value: 'EH' },
		{ text: 'Yemen', value: 'YE' },
		{ text: 'Zambia', value: 'ZM' },
		{ text: 'Zimbabwe', value: 'ZW' }
	];
	let currentPic;
	let countrySelected;
	let birthday;
	let email;
	let username;
	let currentPassword;
	let password;
	let confirmPassword;
	async function load() {
		const response = await fetch('/api/userInfo');
		const data = await response.json();
		user = data.user;
		firstName = data.user.firstName;
		lastName = data.user.lastName;
		bio = data.user.bio;
		birthday = data.user.birthday.slice(0, 10);
		countrySelected = data.user.location;
		currentPic = data.user.profileImgUrl;
		email = data.user.email;
		username = data.user.username;
	}
	onMount(async () => {
		load();
	});
	function imageChange() {
		const file = profilePicInput.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', () => {
				profilePic.setAttribute('src', reader.result);
			});
			reader.readAsDataURL(file);
			currentPic = file;
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
	async function submitPersonal() {
		const formData = new FormData();
		formData.append('firstName', firstName);
		formData.append('lastName', lastName);
		formData.append('bio', bio);
		formData.append('birthday', birthday);
		formData.append('image', profilePicInput.files[0]);
		formData.append('location', countrySelected);
		formData.append('onlineTime', timelineSelected + amPm);
		const response = await fetch('/api/update-profile', {
			method: 'POST',
			body: formData
		});
		if (response.ok) {
			load();
		}
	}
	async function submitAccount() {
		const response = await fetch('/api/update-account', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				email,
				username
			})
		});
		if (response.status == 200) {
			switchAccount();
			load();
		}
	}
	async function changePassword() {
		const response = await fetch('/api/change-password', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				currentPassword,
				newPassword: password,
				newConfirm: confirmPassword
			})
		});
		if (response.status == 200) {
			load();
		}
	}
</script>

<svelte:head>
	<title>Setting</title>
</svelte:head>
<div class="md:p-4 my-2 md:w-9/12 ">
	<section
		class="w-80 flex border-2 justify-evenly p-3 bg-main-bg rounded-full text-white mx-auto mb-4 border-border"
	>
		<button class="hover:text-main text-main" id="personalBtn" on:click={switchPersonal}
			>Personal Setting</button
		>
		<button class="hover:text-main" id="accountBtn" on:click={switchAccount}>Account Setting</button
		>
	</section>
	<!--Personal-->

	<form
		id="personal"
		action="/set-profile"
		method="post"
		class="flex flex-col justify-between p-4 gap-4 items-center  md:border-2 md:border-solid md:border-border w-fit mx-auto"
		enctype="multipart/form-data"
		on:submit|preventDefault={submitPersonal}
	>
		<section class="w-80">
			<label class="text-base text-text" for="firstName">First Name</label>
			<input
				type="text"
				name="firstName"
				id="firstName"
				bind:value={firstName}
				class="outline-none border-2 border-border border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3 text-text"
			/>
		</section>
		<section class="w-80">
			<label class="text-base text-text" for="lastName">Last Name</label>

			<input
				type="text"
				name="lastName"
				id="lastName"
				bind:value={lastName}
				class="outline-none border-2 border-border border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3 text-text"
			/>
		</section>
		<section class="w-80">
			<label class="text-base text-text" for="bio">Bio</label>
			<textarea
				name="bio"
				id="bio"
				cols="20 "
				rows="5"
				bind:value={bio}
				class="outline-none border-2 border-border border-solid
                 rounded-lg px-2.5 mx-auto block w-11/12 my-3 p-2 resize-none
				text-text
                 "
			/>
		</section>
		<section class="w-80">
			<label class="text-base text-text" for="birthday">Bithday</label>
			<input
				type="date"
				name="birthday"
				id=""
				bind:value={birthday}
				class="outline-none border-2 border-border border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3  text-text"
			/>
		</section>
		<section class="w-80">
			<label for="online-time" class="text-base text-text">Select your country</label>
			<section class="flex justify-center">
				<select
					bind:value={countrySelected}
					class="block mx-2 w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg p-1 my-3 "
				>
					<option selected disabled value=""> Country</option>
					{#each countries as country}
						<option value={country.text}>
							{country.text}
						</option>
					{/each}
				</select>
			</section>
		</section>
		<section class="w-80">
			<label class="text-base text-text" for="profilePic">Profile Picture</label>
			<div class="flex items-center justify-evenly my-3   ">
				<!-- svelte-ignore a11y-img-redundant-alt -->

				{#if currentPic}
					<img
						class="h-16 w-16 object-cover rounded-full "
						bind:this={profilePic}
						src="/api/images/{currentPic}"
						alt="Current profile photo"
					/>
				{:else}
					<div
						class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center border-2 border-border"
					>
						<i class="fa-solid fa-user text-slate-400 text-2xl" />
					</div>
				{/if}

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
					file:bg-main-bg file:text-text
					file:border-border  file:border-2
					file:border-solid
					file:rounded-full
					file:hover:text-main
					file:hover:border-main
					file:hover:cursor-pointer"
				/>
			</div>
		</section>
		{#if !currentPic}
			<section class="w-80">
				<label for="online-time" class="text-base text-text">Delete Profile Picture</label>
				<button
					class="cursor-pointer text-lg rounded-lg  bg-red-600 hover:text-border py-2 px-12 mx-auto block w-11/12  text-white  my-3"
				>
					<i class="fa-solid fa-camera bg-transparent" />
					Delete picture
				</button>
			</section>
		{/if}
		<section class="w-80">
			<input
				type="submit"
				value="Save Changes"
				class="main-btn w-11/12 py-2 px-20 text-lg mx-auto block"
			/>
		</section>
	</form>

	<!--auth-->
	<div
		id="account"
		class="flex flex-col justify-between  p-4 gap-3 items-center shadow-xl border-2 hidden w-fit mx-auto"
	>
		<form
			on:submit|preventDefault={submitAccount}
			class="gap-3 items-center flex flex-col justify-between"
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
			<section class="w-80">
				<input
					type="submit"
					value="Save Changes"
					class="cursor-pointer text-lg rounded-lg  bg-main-bg hover:text-main py-2 px-20 mx-auto block w-11/12 hover:shadow-xl text-white "
				/>
			</section>
		</form>
		<form id="changePassword">
			<section class="w-80 ">
				<label class="text-base" for="password">Currnet Password</label>
				<input
					bind:value={currentPassword}
					type="password"
					name="password"
					id="currentPassword"
					class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				/>
			</section>
			<section class="w-80 ">
				<label class="text-base" for="password"> New Password</label>
				<input
					bind:value={password}
					type="password"
					name="password"
					id="password"
					class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				/>
			</section>
			<section class="w-80 ">
				<label class="text-base" for="confirmPassword">Confirm New Password</label>
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
					value="Change Password"
					class="cursor-pointer text-lg rounded-lg  bg-main-bg hover:text-main py-2 px-20 mx-auto block w-11/12 hover:shadow-xl text-white "
				/>
			</section>
		</form>
		<button
			class="cursor-pointer text-lg rounded-lg  bg-red-600 hover:text-black py-2 px-12 mx-auto block w-11/12  text-white "
		>
			<i class="fa-solid fa-user-xmark mx-2 " />
			Delete Account
		</button>
	</div>
</div>
