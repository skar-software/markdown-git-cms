<svelte:options immutable={false} />

<script lang="ts">
  import { onMount } from "svelte";
  export let data;
  import { Carta, MarkdownEditor } from "carta-md";
  import "carta-md/default.css";

  // Remember to use a sanitizer to prevent XSS attacks
  // {sanitizer: mySanitizer}
  const carta = new Carta();
  let value = "";
  let globalRepo: string;
  let globalPath: string;

  const url: string = "/api/file?";
  async function getReps(): Promise<Array<any>> {
    try {
      let repoPath = data.rep.substring(data.rep.split("/")[0].length + 1);
      globalPath = encodeURIComponent((repoPath && "/") + repoPath + "/" + data.file)
      globalRepo = data.rep.split("/").shift();
      const response = await fetch(url + new URLSearchParams({ owner: data.owner, repo: globalRepo, path: globalPath }).toString());
      if (response.status !== 200) {
      } else {
        value = await response.text();
      }
    } catch (error) {
      console.error("Error:", error);
    }
    return [];
  }

  onMount(async () => {
    await getReps();
  });

  let url2 = "/api/upl";
  async function SendFile() {
    let formData = new FormData();
    formData.append("file", globalPath);
    formData.append("content", value);
    formData.append("owner", data.owner);
    formData.append("repo", globalRepo);
    const response = await fetch(url2, {
      method: "POST",
      body: formData,
    });
    window.location.reload();
  }
</script>

<main>
  <h1 class="my-h1">File {data.rep}/{data.file}</h1>
  <MarkdownEditor
    {carta}
    bind:value
  />
  <button on:click={SendFile}>Update file</button>
</main>

<style>
  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 3em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
