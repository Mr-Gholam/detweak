import { writable } from 'svelte/store'


export const IsLoggedIn = writable(false)

export const User = writable({})