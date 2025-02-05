import { toast, type SvelteToastOptions } from '@zerodevx/svelte-toast'

type Theme = { [x: string]: string | number; }

const theme_success: Theme = {
    '--toastBackground': 'green',
    '--toastColor': 'white',
    '--toastBarBackground': 'red'
}
const theme_error: Theme = {
    '--toastBackground': 'green',
    '--toastColor': 'white',
    '--toastBarBackground': 'red'
}
const theme_info: Theme = {
    '--toastBackground': 'green',
    '--toastColor': 'white',
    '--toastBarBackground': 'red'
}
const theme_warning: Theme = {
    '--toastBackground': 'green',
    '--toastColor': 'white',
    '--toastBarBackground': 'red'
}

const show = (message: string, theme: Theme) => toast.push(message, { theme: theme});

export const success = (message: string) => show(message, theme_success);
export const error = (message: string) => show(message, theme_error);
export const info = (message: string) => show(message, theme_info);
export const warning = (message: string) => show(message, theme_warning);
