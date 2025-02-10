<script lang="ts">
    import { Container, Button, Form, FormGroup, Label, Input, Card, CardTitle, CardBody, Tooltip, Col, Row } from "sveltestrap";
    import { error, success, warning } from "../common/toast";
    import { push } from "svelte-spa-router";
    import { check_token_or_login } from "../common/util";
    
    // export let params: any;

    // const competition_id: string = params.id;
    // const is_edit: boolean = competition_id !== null && competition_id !== undefined

    type Competition = {
        title: string;
        public: boolean;
        description: string;
        language: number;
        use_overall_time: boolean;
        use_execution_time: boolean;
        start_date: string;
        start_time: string;
        end_date: string;
        end_time: string;
    };

    check_token_or_login()

    const token: string = sessionStorage.getItem("token") ?? "";
    let loading: boolean = true;
    let languages: string[];

    const load_languages = async () => {
        try {
            const response: Response = await fetch("/api/competitions/languages");

            if (!response.ok) {
                const data: string = await response.text();

                warning(data);

                return;
            }

            languages = await response.json() || [];
        } catch(err: any) {
            error(err)
        } finally {
            loading = false;
        }
    }

    let competition: Competition = {
        title: "",
        public: false,
        description: "",
        language: -1,
        use_overall_time: false,
        use_execution_time: false,
        start_date: "",
        start_time: "",
        end_date: "",
        end_time: ""
    };

    function checkInput(): boolean {
        if (!competition.title || !competition.description || competition.language === -1 ||
            !competition.start_date || !competition.start_time || !competition.end_date || !competition.end_time
        ) {
            warning("Empty competition properties.")

            return false;
        }

        return true;
    }

    async function createCompetition() {
        let is_input_ok: boolean = checkInput()

        if (!is_input_ok) {
            return;
        }

        fetch("/api/competitions", {
            method: "POST",
            headers: { "Content-Type": "application/json", "Authorization": token },
            body: JSON.stringify(competition)
        }).then(async e => {
            let data: string = await e.text();

            if (!e.ok) {
                warning(data);
            } else {
                success("Successfully created competition.")

                push('/dashboard');
            }
        }).catch((err: any) => error(err));
    }

    load_languages();
</script>

{#if loading}
<div>Loading...</div>
{:else}
<Container class="d-flex justify-content-center align-items-center vh-100" fluid>
    <Card class="p-4 shadow-lg" style="max-width: 600px; width: 100%;">
        <CardBody>
            <CardTitle>Competition</CardTitle>

            <Form>
                <FormGroup>
                    <Label for="title">Title</Label>
                    <Input type="text" id="title" bind:value={competition.title} placeholder="Enter competition title" />
                </FormGroup>
        
                <FormGroup>
                    <Label for="description">Description</Label>
                    <Input type="textarea" id="description" bind:value={competition.description} placeholder="Enter competition description" />
                </FormGroup>
        
                <FormGroup>
                    <Label for="language">Language</Label>
                    <Input type="select" id="language" bind:value={competition.language}>
                        {#each languages as language, index}
                        <option value="{index}">{language}</option>
                        {/each}
                    </Input>
                </FormGroup>

                <FormGroup>

                    <Label id="public_label" for="public">Is public?</Label>
                    <Tooltip target="public_label" placement="right">Sets whether the competition is visible to everyone.</Tooltip>
                    <Input type="checkbox" id="public" bind:checked={competition.public}/>

                    <Label id="use_overall_time_label" for="use_overall_time">Use overall time?</Label>
                    <Tooltip target="use_overall_time_label" placement="top">Takes into account overall passed time of the competition when scoring the results.</Tooltip>
                    <Input type="checkbox" id="use_overall_time" bind:checked={competition.use_overall_time}/>

                    <Label id="use_execution_time_label" for="use_execution_time">Use execution time?</Label>
                    <Tooltip target="use_execution_time_label" placement="right">Takes into account execution time of the submission when scoring the results.</Tooltip>
                    <Input type="checkbox" id="use_execution_time" bind:checked={competition.use_execution_time}/>
                </FormGroup>
                
                <Row>
                    <Col>
                        <FormGroup>
                            <Label for="start_date">Start date:</Label>
                            <Input type="date" id="start_date" bind:value={competition.start_date}/>
        
                            <Label for="start_time">Start time:</Label>
                            <Input type="time" id="start_time" bind:value={competition.start_time}/>
                        </FormGroup>
                    </Col>

                    <Col>
                        <FormGroup>
                            <Label for="end_date">End date:</Label>
                            <Input type="date" id="end_date" bind:value={competition.end_date}/>
        
                            <Label for="end_time">End time:</Label>
                            <Input type="time" id="end_time" bind:value={competition.end_time}/>
                        </FormGroup>
                    </Col>
                </Row>  
            </Form>
        
            <div class="d-flex justify-content-between mt-3">
                <Button color="success" on:click={createCompetition}>Create Competition</Button>
            </div>
        </CardBody>
    </Card>
</Container>
{/if}

