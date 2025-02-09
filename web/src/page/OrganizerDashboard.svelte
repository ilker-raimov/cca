<script lang="ts">
    import { push } from "svelte-spa-router";
    import { Table, Button, Badge, Container } from "sveltestrap";
    import { check_token_or_login } from "../common/util";
    import { error, warning } from "../common/toast";

    type Competition = {
        id: string;
        title: string;
        public: boolean;
        description: string;
        language: number;
        use_overall_time: boolean;
        use_execution_time: boolean;
    };

    let loading = true;
    let languages: string[] = [];
    let competitions: Competition[] = [];

    async function getLanguages(): Promise<string[]> {
        try {
            const response: Response = await fetch("/api/competitions/languages");

            if (!response.ok) {
                const data: string = await response.text();

                warning(data);

                return [];
            }

            return await response.json() || {};
        } catch(err: any) {
            error(err)

            return [];
        }
    }

    async function getCompetitionIdList(token: string): Promise<string[]> {
        try {
            const response: Response = await fetch("/api/competitions?all", {
                headers: { "Authorization": token }
            });

            if (!response.ok) {
                let data: string = await response.text();

                warning(data);

                return [];
            }

            return await response.json();
        } catch(err: any) {
            error(err);

            return [];
        }
    }

    async function getCompetition(id: string, token: string): Promise<Competition | null> {
        try {
            const response: Response = await fetch(`/api/competitions/${id}`, {
                headers: { "Authorization": token }
            });

            if (!response.ok) {
                let data: string = await response.text();

                warning(data);

                return null;
            }

            return await response.json();
        } catch(err: any) {
            error(err);

            return null;
        }
    }

    async function getCompetitions() {
        check_token_or_login()

        languages = await getLanguages();

        const token = sessionStorage.getItem("token") ?? "";

        let competition_id_list: string[] = await getCompetitionIdList(token);
        
        for (let competition_id of competition_id_list) {
            const competition: Competition | null = await getCompetition(competition_id, token);

            if (competition == null) {
                continue;
            }

            competitions = [...competitions, competition]
        }

        loading = false;
    }

    function deleteCompetition(id: any) {
        alert(`Deleting competition with ID: ${id}`);
    }

    getCompetitions();
</script>

<Container class="mt-4">
    <div class="d-flex justify-content-between align-items-center mb-3">
        <h2>My Competitions</h2>
        <Button color="success" on:click={() => push("/competitions/create")}>Create Competition</Button>
    </div>

    {#if loading}
    <p>Loading competitions...</p>
    {:else}
    <Table striped bordered hover>
        <thead>
            <tr>
                <th>#</th>
                <th>Title</th>
                <th>Publicity</th>
                <th>Language</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each competitions as comp, index (comp.id)}
                <tr>
                    <td>№{index + 1}</td>
                    <td>{comp.title}</td>
                    <td>
                        <Badge color={comp.public ? "success" : "danger"}>{comp.public ? "Public" : "Private"}</Badge>
                    </td>
                    <td>{languages[comp.language]}</td>
                    <td>
                        <Button color="success" size="sm" on:click={() => push(`/competitions/${comp.id}/view`)}>View</Button>
                        <Button color="primary" size="sm" on:click={() => push(`/competitions/${comp.id}/edit`)}>Edit</Button>
                        <Button color="danger" size="sm" class="ms-2" on:click={() => deleteCompetition(comp.id)}>Delete</Button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </Table>
    {/if}
</Container>
