<script>
  import { browser } from "$app/environment";
  import { onMount } from "svelte";
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
  onMount(() => {
    if (browser) {
      // if (token === null) {
      window.location.href = "/login";
      // } else {
      //   window.location.href = "/repos";
      // }
    }
  });
</script>

<main></main>
