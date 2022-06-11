<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch('/api/setting');
		let user;
		let firstName;
		let lastName;
		let bio;
		let birthday;
		let email;
		let username;
		let profileImage;
		let selectedLanguage = null;
		let selectedField = null;
		let selectedFrameWork = null;
		let gitHubUserName = '';
		let experience = null;
		let countrySelected;
		let currentPic;
		if (res.status == 200) {
			const data = await res.json();
			user = JSON.parse(JSON.stringify(data));
			firstName = data.Firstname;
			lastName = data.Lastname;
			gitHubUserName = data.GitHub;
			bio = data.Bio;
			birthday = data.Birthday.slice(0, 10);
			countrySelected = data.Location;
			currentPic = data.ImgUrl;
			profileImage = data.ImgUrl;
			email = data.Email;
			username = data.Username;
			selectedLanguage = data.Language;
			selectedField = data.Field;
			selectedFrameWork = data.FrameWork;
			experience = data.Experience;
		}
		return {
			props: {
				user,
				firstName,
				lastName,
				gitHubUserName,
				bio,
				birthday,
				countrySelected,
				currentPic,
				profileImage,
				email,
				selectedLanguage,
				selectedField,
				selectedFrameWork,
				experience,
				username
			}
		};
	};
</script>

<script>
	// @ts-nocheck
	import { User } from '../store';
	import { loading } from '../store';
	export let user;
	export let firstName;
	export let lastName;
	export let bio;
	export let ProfilePic;
	export let selectedLanguage;
	export let selectedField;
	export let selectedFrameWork;
	export let gitHubUserName;
	export let experience;
	export let birthday;
	export let email;
	export let profileImage;
	export let username;
	export let currentPic;
	export let countrySelected;
	let profilePic = ProfilePic;
	let profilePicInput;
	let currentPassword;
	let password;
	let confirmPassword;
	let usernameLengthError = false;
	let emailError = false;
	let usernameError = false;
	let showPasswordRules = false;
	let uppercaseError = false;
	let lowercaseError = false;
	let passwordNumberError = false;
	let passwordLength = false;
	let passwordsNotMatch = false;
	let incorectPassword = false;
	let hasPhoto = false;
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
	let languages = ['', 'JavaScript', 'PHP', 'Java', 'C', 'C++', 'C#', 'Go', 'Python', 'Ruby'];
	let exOption = [
		'',
		'Less than 6 months',
		'Less than 1 year',
		'Between 1-2 years',
		'Between 2-4 years',
		'More than 4 years'
	];
	let Fields = ['', 'Front-end', 'Back-end', 'Full-stack', 'Game Developer'];
	let frontEndFrameWorks = ['', 'React', 'Vue', 'Svelte', 'Anguler'];

	function imageChange() {
		const file = profilePicInput.files[0];
		if (file) {
			const reader = new FileReader();
			reader.addEventListener('load', () => {
				profilePic.setAttribute('src', reader.result);
			});
			reader.readAsDataURL(file);
			currentPic = file;
			changePic();
			return;
		}
	}
	function switchPersonal() {
		const personal = document.getElementById('personal');
		const account = document.getElementById('account');
		const Perfessional = document.getElementById('professional');
		const PerfessionalBtn = document.getElementById('professionalBtn');
		const personalBtn = document.getElementById('personalBtn');
		const accountBtn = document.getElementById('accountBtn');
		personal.classList.replace('hidden', 'flex');
		account.classList.replace('flex', 'hidden');
		Perfessional.classList.replace('flex', 'hidden');
		personalBtn.classList.replace('text-text', 'text-white');
		accountBtn.classList.replace('text-white', 'text-text');
		PerfessionalBtn.classList.replace('text-white', 'text-text');
	}
	function switchAccount() {
		const personal = document.getElementById('personal');
		const account = document.getElementById('account');
		const personalBtn = document.getElementById('personalBtn');
		const accountBtn = document.getElementById('accountBtn');
		const Perfessional = document.getElementById('professional');
		const PerfessionalBtn = document.getElementById('professionalBtn');
		personal.classList.replace('flex', 'hidden');
		account.classList.replace('hidden', 'flex');
		Perfessional.classList.replace('flex', 'hidden');
		personalBtn.classList.replace('text-white', 'text-text');
		accountBtn.classList.replace('text-text', 'text-white');
		PerfessionalBtn.classList.replace('text-white', 'text-text');
	}
	function switchPerfessional() {
		const personal = document.getElementById('personal');
		const account = document.getElementById('account');
		const personalBtn = document.getElementById('personalBtn');
		const accountBtn = document.getElementById('accountBtn');
		const Perfessional = document.getElementById('professional');
		const PerfessionalBtn = document.getElementById('professionalBtn');
		personal.classList.replace('flex', 'hidden');
		account.classList.replace('flex', 'hidden');
		Perfessional.classList.replace('hidden', 'flex');
		personalBtn.classList.replace('text-white', 'text-text');
		accountBtn.classList.replace('text-white', 'text-text');
		PerfessionalBtn.classList.replace('text-text', 'text-white');
	}
	async function submitPersonal() {
		const btn = document.getElementById('submitPersonal');
		if (hasPhoto) {
			const formData = new FormData();
			formData.append('image', profilePicInput.files[0]);
			const image = await fetch('/api/update-profileImg', {
				method: 'post',
				body: formData
			});
			const imageJson = await image.json();
			if (imageJson.ImgUrl) {
				hasPhoto = false;
				document.getElementById('profileImg').value = '';
				profileImage = imageJson.ImgUrl;
				$User.ImgUrl = imageJson.ImgUrl;

				btn.value = 'User Updated';
				btn.classList.add('text-green', 'border-green');
				setTimeout(() => {
					btn.value = 'Save Changes';
					btn.classList.remove('text-green', 'border-green');
				}, 3000);
			}
		}
		const response = await fetch('/api/update-personal', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				Firstname: user.Firstname != firstName ? firstName : undefined,
				Lastname: user.Lastname != lastName ? lastName : undefined,
				Bio: user.Bio != bio ? bio : undefined,
				Location: user.Location != countrySelected ? countrySelected : undefined,
				Birthday: user.Birthday != birthday ? birthday : undefined
			})
		});
		if (response.ok) {
			btn.value = 'User Updated';
			btn.classList.add('text-green', 'border-green');
			setTimeout(() => {
				btn.value = 'Save Changes';
				btn.classList.remove('text-green', 'border-green');
			}, 3000);
		}
	}
	function checkUsername() {
		if (user.Username != username) {
			if (username.length <= 3) {
				usernameLengthError = true;
			} else {
				usernameLengthError = false;
			}
		}
	}
	async function submitAccount() {
		const btn = document.getElementById('accountSubmitBtn');
		const response = await fetch('/api/update-account', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				Email: user.Email != email ? email : undefined,
				Username: user.Username != username && username.length >= 3 ? username : undefined
			})
		});
		const json = await response.json();
		if (json.error) {
			if (json.error.message == 'Username already exists') {
				usernameError = true;
			}
			if (json.error.message == 'Email already exists') {
				emailError = true;
			}
		}
		if (response.ok) {
			emailError = false;
			usernameError = false;
			btn.value = 'User Updated';
			btn.classList.add('text-green', 'border-green');
			setTimeout(() => {
				btn.value = 'Save Changes';
				btn.classList.remove('text-green', 'border-green');
			}, 3000);
		}
	}
	function checkNewPassword() {
		showPasswordRules = true;
		if (!/[A-Z]/.test(password)) {
			uppercaseError = true;
		} else if (!/[a-z]/.test(password)) {
			lowercaseError = true;
		} else if (!/[0-9]/.test(password)) {
			passwordNumberError = true;
		} else if (password.length <= 7) {
			passwordLength = true;
		} else {
			uppercaseError = false;
			lowercaseError = false;
			passwordNumberError = false;
			passwordLength = false;
		}
	}
	function matchConfirm() {
		if (password !== confirmPassword) {
			passwordsNotMatch = true;
		} else {
			passwordsNotMatch = false;
		}
	}
	async function changePassword() {
		const btn = document.getElementById('passwordBtn');
		if (currentPassword == undefined || password == undefined || confirmPassword == undefined)
			return;
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
		const data = await response.json();
		if (data.error) {
			incorectPassword = true;
		}
		if (data.message) {
			btn.value = 'Password Updated';
			btn.classList.add('text-green', 'border-green');
			setTimeout(() => {
				btn.value = 'Save Changes';
				btn.classList.remove('text-green', 'border-green');
			}, 3000);
		}
	}
	async function removeImgProfile() {
		const response = await fetch('/api/remove-profileImg', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				ImgUrl: currentPic
			})
		});
		if (response.ok) {
			currentPic = '';
			profileImage = '';
		}
	}
	function changePic() {
		if (hasPhoto) {
			hasPhoto = false;
			currentPic = '';
			document.getElementById('profileImg').value = '';
		} else {
			hasPhoto = true;
		}
	}
	async function submitprofessional() {
		const btn = document.getElementById('professionalSubmitBtn');
		const response = await fetch('/api/update-professional', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				GitHub: user.gitHub != gitHubUserName ? gitHubUserName : undefined,
				Experience: user.Experience != experience ? experience : undefined,
				Language: user.Language != selectedLanguage ? selectedLanguage : undefined,
				Field: user.Field != selectedField ? selectedField : undefined,
				FrameWork: user.FrameWork != selectedFrameWork ? selectedFrameWork : undefined
			})
		});
		if (response.ok) {
			btn.value = 'User Updated';
			btn.classList.add('text-green', 'border-green');
			setTimeout(() => {
				btn.value = 'Save Changes';
				btn.classList.remove('text-green', 'border-green');
			}, 3000);
		}
	}
	async function deleteAccount() {
		const response = await fetch('/api/delete-account', { method: 'POST' });
		if (response.ok) {
			console.log('done');
			$User = null;
			location.replace('/');
		}
	}
</script>

<svelte:head>
	<title>Setting</title>
</svelte:head>
<div class="md:p-4 my-2 md:w-9/12 lg:flex  ">
	<section
		class="w-80 flex border-2 justify-evenly p-3  rounded-full text-white mx-auto mb-4 border-border lg:flex-col  h-52 lg:w-fit lg:mx-0 lg:fixed lg:top-24 lg:rounded-md lg:my-2 lg:border-0 "
	>
		<button
			class="hover:text-main  text-left font-semibold text-white "
			id="personalBtn"
			on:click={switchPersonal}>Personal Setting</button
		>
		<button
			class="hover:text-main text-left font-semibold text-text"
			id="professionalBtn"
			on:click={switchPerfessional}>Professional Setting</button
		>
		<button
			class="hover:text-main text-left font-semibold text-text"
			id="accountBtn"
			on:click={switchAccount}>Account Setting</button
		>
	</section>
	<!--Personal-->

	<form
		id="personal"
		action="/set-profile"
		method="post"
		class="flex flex-col justify-between p-4 gap-4 items-center  md:border-2 md:border-solid md:border-border w-fit mx-auto lg:my-6"
		enctype="multipart/form-data"
		on:submit|preventDefault={submitPersonal}
	>
		<section class="w-80">
			<h4 class="text-white text-xl text-center my-2">Personal Setting</h4>
		</section>
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
				id="birthday"
				bind:value={birthday}
				class="outline-none border-2 border-border border-solid rounded-lg px-2.5 mx-auto block w-11/12 p-2 my-3  text-text"
			/>
		</section>
		<section class="w-80">
			<label for="online-time" class="text-base text-text">Select your country</label>
			<section class="flex justify-center">
				<select
					bind:value={countrySelected}
					class="block mx-2 w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg my-3 p-2"
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
						src="/api/images/{profileImage}"
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
					file:hover:cursor-pointer {hasPhoto ? 'file:text-main file:border-main' : ''}
					"
				/>
				{#if hasPhoto}
					<button
						on:click={changePic}
						class=" w-6 h-6 rounded-full text-center text-text  hover:text-error"
					>
						<i class="fa-solid fa-x" />
					</button>
				{/if}
			</div>
		</section>
		{#if profileImage}
			<section class="w-80">
				<label for="online-time" class="text-base text-text">Delete Profile Picture</label>
				<button
					on:click={removeImgProfile}
					class="cursor-pointer text-lg rounded-lg  border-red-600 border-2 text-red-600 hover:text-white py-2 px-12 mx-auto block w-11/12 hover:bg-red-600 my-3"
				>
					<i class="fa-solid fa-camera bg-transparent mx-2" />
					Delete picture
				</button>
			</section>
		{/if}
		<section class="w-80">
			<input
				id="submitPersonal"
				type="submit"
				value="Save Changes"
				class="main-btn w-11/12 py-2 px-20 text-lg mx-auto block border-2"
			/>
		</section>
	</form>
	<!-- perfossinal -->
	<div
		id="professional"
		class="hidden flex-col justify-between  p-4 gap-3 items-center shadow-xl border-2  w-fit mx-auto border-border lg:my-6"
	>
		<form
			class="flex flex-col justify-between  gap-4 items-center"
			on:submit|preventDefault={submitprofessional}
		>
			<section class="w-80">
				<h4 class="text-white text-xl text-center my-2">Professional Setting</h4>
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
					id="experience"
					class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-2 px-2.5"
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
					id="field"
					class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-2 px-2.5"
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
					id="language"
					class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-2 px-2.5"
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
					id="frameWork"
					class="block mx-auto w-11/12 text-center outline-none border-2 color-main border-border text-text border-solid rounded-lg  my-3 p-2 px-2.5"
				>
					{#each frontEndFrameWorks as frameWork}
						<option value={frameWork}>{frameWork}</option>
					{/each}
				</select>
			</section>
			<section class="w-80">
				<input
					id="professionalSubmitBtn"
					type="submit"
					value="Save Changes"
					class=" text-lg py-2 px-20 mx-auto block w-11/12 hover:shadow-xl main-btn my-3 border-2  "
				/>
			</section>
		</form>
	</div>

	<!--auth-->
	<div
		id="account"
		class=" flex-col justify-between  p-4 gap-3 items-center shadow-xl border-2 hidden w-fit mx-auto border-border lg:my-6"
	>
		<!-- password setting -->
		<form
			id="changePassword"
			class="mb-8 flex flex-col gap-4"
			on:submit|preventDefault={changePassword}
		>
			<h1 class="text-lg font-semibold text-center text-white">Password Settting</h1>
			<section class="w-80 {incorectPassword ? 'block' : 'hidden'}">
				<h1 class="text-center text-error font-semibold">Incorect Password</h1>
			</section>
			<section class="w-80 ">
				<label class="text-base text-text" for="password">Currnet Password</label>
				<input
					bind:value={currentPassword}
					type="password"
					name="password"
					id="currentPassword"
					class="outline-none border-2  text-text border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 {incorectPassword
						? 'border-error'
						: 'border-border'}"
				/>
			</section>
			<section class="w-80 ">
				<label class="text-base text-text" for="password"> New Password</label>
				<input
					on:change={checkNewPassword}
					bind:value={password}
					type="password"
					name="password"
					id="password"
					class="outline-none border-2 border-border text-text border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				/>
			</section>
			<section class="w-80 {passwordsNotMatch ? 'block' : 'hidden'}">
				<h1 class="text-center text-error font-semibold">Passwords do not match</h1>
			</section>
			<section class="w-80 ">
				<label class="text-base text-text" for="confirmPassword">Confirm New Password</label>
				<input
					on:change={matchConfirm}
					bind:value={confirmPassword}
					type="password"
					name="confirmPassword"
					id="confirmPassword"
					class="outline-none border-2 text-text border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 {passwordsNotMatch
						? 'border-error'
						: 'border-border'}"
				/>
			</section>
			<section class="w-80 {showPasswordRules ? 'block' : 'hidden'}">
				<ul class=" px-4">
					<li class="text-xs my-1 {uppercaseError ? 'text-error' : 'text-text'}" id="uppercase">
						<p>At least one uppercase letter</p>
					</li>
					<li class="text-xs my-1 {lowercaseError ? 'text-error' : 'text-text'}" id="lowercase">
						<p>At least one lowercase letter</p>
					</li>
					<li class="text-xs my-1 {passwordNumberError ? 'text-error' : 'text-text'}" id="number">
						<p>contain at least one number</p>
					</li>
					<li class="text-xs my-1 {passwordLength ? 'text-error' : 'text-text'}" id="length">
						<p>should be more than 8 character</p>
					</li>
				</ul>
			</section>
			<section class="w-80">
				<input
					id="passwordBtn"
					type="submit"
					value="Change Password"
					class="main-btn py-3 px-20 mx-auto block w-11/12 my-3 border-2 "
				/>
			</section>
		</form>
		<!-- account setting  -->
		<form
			on:submit|preventDefault={submitAccount}
			class="gap-3 items-center flex flex-col justify-between mb-8"
		>
			<h1 class="text-lg font-semibold text-center text-white">Account Settting</h1>
			<!-- email error -->
			<section class="w-80 {emailError ? 'block' : 'hidden'}">
				<h1 class="text-center text-error font-semibold">Email already exists</h1>
			</section>
			<section class="w-80 ">
				<label class="text-base text-text" for="email">Email</label>
				<input
					bind:value={email}
					type="email"
					name="email"
					id="email"
					class="outline-none border-2  border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text {emailError
						? 'border-error'
						: 'border-border'}"
				/>
			</section>
			<!-- username Error -->
			<section class="w-80 {usernameError || usernameLengthError ? 'block' : 'hidden'}">
				<h1 class="text-center text-error font-semibold {usernameError ? 'block' : 'hidden'}">
					Username already exists
				</h1>
				<h1 class="text-center text-error font-semibold{usernameLengthError ? 'block' : 'hidden'}">
					Username must be at least 3 charcthers
				</h1>
			</section>
			<section class="w-80 ">
				<label class="text-base  text-text" for="username">Username</label>
				<input
					on:change={checkUsername}
					bind:value={username}
					type="text"
					name="username"
					id="username"
					class="outline-none border-2  text-text border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 {usernameError ||
					usernameLengthError
						? 'border-error'
						: 'border-border'}"
				/>
			</section>
			<section class="w-80">
				<input
					id="accountSubmitBtn"
					type="submit"
					value="Save Changes"
					class="main-btn block w-11/12 mx-auto px-20 py-3 my-3 border-2"
				/>
			</section>
		</form>
		<!-- delete account  -->
		<section class="w-80">
			<label class="text-base text-text" for="confirmPassword">Delete Account</label>
			<button
				on:click={deleteAccount}
				class="cursor-pointer text-lg rounded-lg  border-red-600 border-2 text-red-600 hover:text-white py-2 px-12 mx-auto block w-11/12 hover:bg-red-600 my-3 "
			>
				<i class="fa-solid fa-user-xmark mx-2  bg-transparent" />
				Delete Account
			</button>
		</section>
	</div>
</div>
