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
      globalPath = encodeURIComponent((repoPath && "/") + repoPath + "/" + data.file);
      globalRepo = (data.rep.split("/") ?? [""])[0];
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
    if (response.status === 200) {
      alert("File successfully updated and committed to Github.");
      window.location.reload();
    } else {
      alert("Error updating file");
      console.error("Error:", response);
    }
  }
</script>

<main class="main">
  <h1 class="my-h1">File {data.rep}/{data.file}</h1>
  <MarkdownEditor
    {carta}
    bind:value
  />
  <button
    class="button"
    on:click={SendFile}>Update file</button
  >
</main>

<style>
    * {
    font-family: poppins;
  }
  :global(.my-h1) {
    color: #415a77;
    text-decoration: underline;
    font-size: 2em;
    font-weight: 600;
    font-family: poppins;
    text-align: center;
    margin-top: 0;
  }
  :global(.main) {
    padding: 1em;
    margin: 0 auto;
    background-color: #edede9;
    height: 100%;
    display: flex;
    flex-direction: column;

    text-align:left;
    max-width: 100vw;
    justify-content:flex-start;
    align-items:normal;
    margin-top: 0;
  }
  :global(.button) {
    padding: 15px 30px 15px 30px;
    background-color: #415a77;
    color: #edede9;
    border: none;
    outline: none;
    font-size: 1rem;
    font-weight: 600;
    border-radius: 24px;
    cursor: pointer;
    margin: 20px auto;
    margin-bottom: 70px;
  }
  :global(.carta-font-code) {
	font-family: '...', monospace;
	font-size: 1.1rem;
	line-height: 1.25rem;
}
</style>
