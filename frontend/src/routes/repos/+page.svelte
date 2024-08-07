<svelte:options immutable={false} />

<script lang="ts">
  import { onMount } from "svelte";
  import "../../style.css";

  interface Rep {
    Name: string;
    Owner: string;
  }
  let page = 0;
  let r: Rep[] = [{ Name: "loading...", Owner: "" }];
  let end = false;
  const url: string = "/api/rep";
  async function getReps(): Promise<Rep[]> {
    try {
      const response = await fetch(url);
      const res = JSON.parse(await response.text()) as Rep[];
      r = res;
      page = 1;
    } catch (error) {
      console.error("Error:", error);
    }
    return r;
  }
  async function loadMore() {
    page++;
    const response = await fetch(url + "?page=" + page);
    const res = JSON.parse(await response.text()) as Rep[];
    r = r.concat(res);
    if (res.length === 0) {
      end = true;
    }
  }
  onMount(async () => {
    await getReps();
  });
</script>

<main>
  <h1>Choose your repository</h1>
  <ul>
    {#each r as rep}
      {#if rep.Name !== ""}
        <li>
          {#if rep.Name === "loading..."}
            <p class="button">{rep.Name}</p>
          {:else}
            <a
              href="/repos/{rep.Owner}/{rep.Name}/"
              class="button">{rep.Name}</a
            >
          {/if}
        </li>
        <!-- <tr><p>---------------------</p></tr> -->
      {/if}
    {/each}
    {#if !end}
      <button style="margin-bottom: 50px;" on:click={loadMore}>LOAD MORE</button>
    {/if}
  </ul>
</main>