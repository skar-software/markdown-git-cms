<script lang="ts">
  import "../../style.css";
  let err = "";

  async function handleSubmit(event: SubmitEvent) {
    const form = event.target as HTMLFormElement;
    const data = new FormData(form);

    let url = "/api/set/?";
    let tknVal = data.get("tkn")?.toString();
    if (tknVal === undefined) {
      tknVal = "";
    }
    const response = await fetch(url + new URLSearchParams({ tkn: tknVal }).toString(), {
      method: "Get",
    });
    if (response.status !== 200) {
      err = await response.text();
    } else {
      window.location.href = "/repos";
    }
  }
</script>

<main>
  <h1>Set up your gihub token!</h1>
  <form
    class="login-form"
    on:submit|preventDefault={handleSubmit}
  >
    <label>
      <span>Github token</span>
      <input
        class="file-input"
        type="password"
        name="tkn"
      />
    </label>
    <button
      class="login-button"
      type="submit">Set token</button
    >
    {#if err !== ""}
      <span style="color: #ff3e00;">{err}</span>
    {/if}
  </form>
</main>

<style>
  main {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 50vh;
  }
</style>
