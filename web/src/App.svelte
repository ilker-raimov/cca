<script lang="ts">
	import { onMount } from 'svelte';
	import { Router, Route, Link, navigate } from 'svelte-routing';
    import { SvelteToast } from '@zerodevx/svelte-toast'

	import Login from './page/Login.svelte';
	import Dashboard from './page/Dashboard.svelte';

  	onMount(() => {
		if (window.location.pathname === '/') {
			navigate('/login');
		}
	});
</script>

<style>
	@import 'svelte-notifications/style.css';
</style>

<Router>
	<Route path="/login" component={Login} />
	<Route path="/dashboard" component={Dashboard} />
</Router>

<SvelteToast/>