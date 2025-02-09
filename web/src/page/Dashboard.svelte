<script lang="ts">
    import { check_token_or_login, get_dashboard } from "../common/util";
    import { error, success, warning } from "../common/toast";
    import { push } from "svelte-spa-router";

    check_token_or_login()

    let dashboard: any = null;
    let loading = true;

    const load_dashboard = async () => {
        try {
            const response = await fetch("/api/users/roles");

            if (!response.ok) {
                const message = await response.text();

                warning(message);

                return;
            }

            const roles = await response.json();
            let username: string | null = sessionStorage.getItem("username");
            let email: string | null = sessionStorage.getItem("email");
            let role: string | null = sessionStorage.getItem("role");

            if (!username || !email || !role) {
                push("/login");
                error("Missing user data!");

                return;
            }

            let role_index: number = parseInt(role, 10);

            dashboard = await get_dashboard(role_index);
        } catch (err: any) {
            error(err);
        } finally {
            loading = false;
        }
    };

    load_dashboard();
</script>

{#if loading}
<div>Loading...</div>
{:else if dashboard !== null}
<svelte:component this={dashboard} />
{/if}

  