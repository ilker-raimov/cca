<script lang="ts">
    import { onMount } from "svelte";
    import { Container, Card, CardBody, CardTitle, Spinner, Alert } from "sveltestrap";
    import { push } from 'svelte-spa-router';
    import { check_token_or_login } from "../common/util";
    import { error, warning } from "../common/toast";

    check_token_or_login()

    let username = "";
    let role = "";
    let loading = true;

    onMount(async () => {
        try {
            const response = await fetch("/api/user/me", {
                method: "GET"
            });

            if (!response.ok) {
                let data = await response.text();

                warning(data);

                return;
            }

            const data = await response.json();

            username = data.username;
            role = data.role;
        } catch (err: any) {
            error(err)
        } finally {
            loading = false;
        }
    });
</script>

<Container class="d-flex justify-content-center align-items-center vh-100">
    <Card class="p-4 shadow-lg text-center" style="max-width: 500px; width: 100%;">
        <CardBody>
            {#if loading}
                <Spinner color="primary" />
                <p class="mt-2">Loading dashboard...</p>
            {:else}
                <CardTitle>
                    <h2>Welcome, {username}! This is the {role} dashboard</h2>
                </CardTitle>
            {/if}
        </CardBody>
    </Card>
</Container>

  