<script lang="ts">
    import { onMount } from "svelte";
    import { tick } from "svelte";

    export let id: string = "editor";
    export let theme: string = "chrome";
    export let mode: string = "java";
    export let read_only: boolean = false;
    export let width: string = "100%";
    export let height: string = "400px";
    export let onCodeChange = (code: string) => {
        console.log(code);
    };

    let editor: any;
    let code = "";

    onMount(async () => {
        await tick();

        editor = window.ace.edit(id);

        editor.setTheme(`ace/theme/${theme}`);
        editor.session.setMode(`ace/mode/${mode}`);
        editor.setReadOnly(read_only);

        editor.session.on("change", () => {
            code = editor.getValue();

            onCodeChange(code);
        });
    });
</script>

<div id="{id}" style="width: {width}; height: {height}"></div>

<svelte:head>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.37.0/ace.js" type="text/javascript" charset="utf-8"></script>
</svelte:head>
  