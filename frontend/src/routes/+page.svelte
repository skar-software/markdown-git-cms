<script>
  import { browser } from "$app/environment";
  import "../style.css";

  // Function to parse cookies and extract the token
  // @ts-ignore
  function getCookie(name) {
    if (browser && document.cookie) {
      const value = `; ${document.cookie}`;
      const parts = value.split(`; ${name}=`);
      // @ts-ignore
      if (parts.length === 2) return parts.pop().split(";").shift();
    }
  }

  // Extract the token from cookies
  const token = browser ? getCookie("tkn") : null;
</script>

<main>
  <h1>Hey, login to your gihub!</h1>
  {#if token}
    <a href="/repos">go to repos</a>
  {:else}
    <a href="/login">login</a>
  {/if}
</main>