import { writable } from 'svelte/store'



export const User = writable({ username: null })

export const loading = writable(false)

export const onlineFriend = writable([])
export const ws = writable(null)