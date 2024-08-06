<svelte:options immutable={false} />

<script lang="ts">
  import { onMount } from "svelte";
  export let data;

  interface Dir {
    Path: string;
    Type: string;
  }

  let r: Dir[] = [{ Path: "loading...", Type: "" }];
  const url: string = "/api/files?";
  async function getReps(): Promise<Dir[]> {
    try {
      const response = await fetch(url + new URLSearchParams({ owner: data.owner, repo: data.rep }).toString());
      const res = JSON.parse(await response.text()) as Dir[];
      r = res;

    } catch (error) {
      console.error("Error:", error);
    }
    return [];
  }

  onMount(async () => {
    await getReps();
  });
  async function handleSubmit(event: SubmitEvent) {
    const form = event.target as HTMLFormElement;
    const d = new FormData(form);
    let repo = encodeURIComponent(data.rep);
    // console.log();
    window.location.href = `/repos/${data.owner}/${repo}/${d.get("file") + ".md"}`;
  }
  function OnClick(file: Dir) {
    // file.Path.indexOf("/") === 0 ? (file.Path = file.Path.slice(1)) : file.Path;
    data.rep = decodeURIComponent(data.rep).split("/").shift();
    file.Path = encodeURIComponent(data.rep + "/" + file.Path);
    if (file.Type === "dir") {
      window.location.href = `/repos/${data.owner}/${file.Path}`;
    } else if (file.Type === "file") {
      const lastIndex = file.Path.lastIndexOf("%2F");
      if (lastIndex !== -1) {
        file.Path = file.Path.substring(0, lastIndex) + "/" + file.Path.substring(lastIndex + 3);
      }
      window.location.href = `/repos/${data.owner}/${file.Path}`;
    }
  }
</script>

<main>
  <h1>Choose file from {data.rep} repository</h1>
  <form on:submit|preventDefault={handleSubmit}>
    <label>
      <span>New file name</span>
      <input name="file" />.md
    </label>
    <button>Create new file</button>
  </form>
  <table>
    <tr><p>---------------------</p></tr>
    {#each r as file}
      {#if file.Path !== ""}
        <tr>
          {#if file.Path === "loading..."}
            <p class="button">{file.Path}</p>
          {:else}
            <button
              class="button"
              on:click={() => OnClick(file)}>{file.Path}</button
            >
          {/if}
        </tr>
      {/if}
    {/each}
  </table>
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 3em;
    font-weight: 100;
  }

  table {
    color: #ff3e00;
    margin-left: auto;
    margin-right: auto;
    font-size: 1.5em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
