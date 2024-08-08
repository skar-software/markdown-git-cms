<svelte:options immutable={false} />

<script lang="ts">
  import { onMount } from "svelte";
  import "../../../../style.css";
  export let data;

  import { getContext } from "svelte";

  const titleStore = getContext("title");
  // Update the title when this component is mounted
  titleStore.set(`Choose file from ${data.rep} repository`);

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

  async function CreateFile(filename: string) {
    let formData = new FormData();
    formData.append("file", decodeURIComponent(data.rep) + "/" + filename);
    formData.append("content", "");
    formData.append("owner", data.owner);
    let globalRepo = decodeURIComponent(data.rep).split("/").shift();
    formData.append("repo", globalRepo ?? "");
    const response = await fetch("/api/upl", {
      method: "POST",
      body: formData,
    });
    if (response.status === 200) {
      alert("File created successfully");
      window.location.reload();
    } else {
      alert("Error creating file");
      console.error("Error:", response);
    }
  }

  async function handleSubmit(event: SubmitEvent) {
    const form = event.target as HTMLFormElement;
    const d = new FormData(form);
    let repo = encodeURIComponent(data.rep);
    CreateFile(d.get("file")?.toString() ?? "");
    // window.location.href = `/repos/${data.owner}/${repo}/${d.get("file") + ".md"}`;
  }
  function OnClick(file: Dir) {
    titleStore.set("");
    // file.Path.indexOf("/") === 0 ? (file.Path = file.Path.slice(1)) : file.Path;
    data.rep = decodeURIComponent(data.rep).split("/").shift() ?? "";
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
  <form
    class="file-form"
    on:submit|preventDefault={handleSubmit}
  >
    <span>New file name</span>
    <input
      class="file-input"
      name="file"
      type="text"
      value=".md"
    />
    <button class="file-button">Create new file</button>
  </form>
  <table>
    <tr>
      <th>#</th>
      <th>Name</th>
    </tr>
    {#each r as file, i}
      {#if file.Path !== ""}
        <tr>
          {#if file.Path !== "loading..."}
            <td>{i + 1}</td>
            <td>
              <button
                on:click={() => OnClick(file)}
                class="repo-button"
              >
                {#if file.Type === "dir"}
                  <img
                    style="margin-right: 10px;"
                    src="/folder.svg"
                    alt="folder"
                  />
                {:else}
                  <img
                    style="margin-right: 10px;"
                    src="/markdown.svg"
                    alt="markdown"
                  />
                {/if}
                {file.Path}
              </button>
            </td>
          {/if}
        </tr>
      {/if}
    {/each}
  </table>
</main>

<style>
  .repo-button {
    border: none;
    background: none;
    cursor: pointer;
    font-size: medium;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .file-form {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
  }
</style>
