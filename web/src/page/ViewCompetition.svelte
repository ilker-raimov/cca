<script lang="ts">
    import { Card, CardBody, CardTitle, CardText, Badge, Container, Button } from "sveltestrap";
    import { getCompetition, type Competition } from "../common/competition";
    import { check_token_or_login } from "../common/util";
    import { push } from "svelte-spa-router";

    export let params: any;

    let loading = true;
    let competition: Competition;

    async function load() {
        check_token_or_login();

        const token = sessionStorage.getItem("token") ?? "";

        let temp: Competition | null = await getCompetition(params.id, token);

        if (temp) {
            competition = temp;
        }

        loading = false;
    }

    load();
</script>

{#if loading}
<p>Loading competition...</p>
{:else}
<Container class="mt-5">
    <Card>
        <CardBody>
            <CardTitle>{competition.title}</CardTitle>
            <CardText>
                <strong>Description:</strong>
                {competition.description}
            </CardText>

            <CardText>
                <strong>Language:</strong>
                {competition.language}
            </CardText>

            <CardText>
                <strong>Public:</strong>
                {#if competition.public}
                <Badge color="success">Yes</Badge>
                {:else}
                <Badge color="danger">No</Badge>
                {/if}
            </CardText>
            <CardText>
                <strong>Overall Time:</strong>
                {#if competition.use_overall_time}
                <Badge color="info">Enabled</Badge>
                {:else}
                <Badge color="secondary">Disabled</Badge>
                {/if}
            </CardText>
            <CardText>
                <strong>Execution Time:</strong>
                {#if competition.use_execution_time}
                <Badge color="info">Enabled</Badge>
                {:else}
                <Badge color="secondary">Disabled</Badge>
                {/if}
            </CardText>

            <Button color="primary" on:click={() => push(`/competitions/${competition.id}/tasks/create`)}>Create Task</Button>
        </CardBody>
    </Card>
</Container>
{/if}