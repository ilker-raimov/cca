import { toast } from '@zerodevx/svelte-toast'

type Theme = { [x: string]: string | number; }

const theme_success: Theme = {
    '--toastBackground': '#4CAF50', // Green
    '--toastColor': '#FFFFFF', // White
    '--toastBarBackground': '#2E7D32' // Darker Green
}
const theme_error: Theme = {
    '--toastBackground': '#F44336', // Red
    '--toastColor': '#FFFFFF', // White
    '--toastBarBackground': '#B71C1C' // Darker Red
}
const theme_info: Theme = {
    '--toastBackground': '#E0E0E0', // Light Gray (Whitish)
    '--toastColor': '#212121', // Dark Gray (Readable)
    '--toastBarBackground': '#9E9E9E' // Mid Gray
}
const theme_warning: Theme = {
    '--toastBackground': '#FF9800', // Orange
    '--toastColor': '#FFFFFF', // White
    '--toastBarBackground': '#E65100' // Darker Orange
}

const show = (message: string, theme: Theme) => toast.push(message, { theme: theme});

export const success = (message: string) => show(message, theme_success);
export const error = (message: string) => show(message, theme_error);
export const info = (message: string) => show(message, theme_info);
export const warning = (message: string) => show(message, theme_warning);
