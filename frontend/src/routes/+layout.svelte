<script lang="ts">
  function handleOnClick(route: string, type = "normal") {
    if (type === "normal") {
      window.location.href = route;
    } else if (type === "help") {
      window.open(route, "_blank");
    }
  }
</script>

<div class="app">
  <header>
    <div>
      <button
        class="header-button"
        on:click={() => handleOnClick("/login")}>Connect</button
      >
      |
      <button
        class="header-button"
        on:click={() => handleOnClick("/repos")}>Repos</button
      >
      |
      <button
        class="header-button"
        on:click={() => handleOnClick("/repos")}>Markdown Editor</button
      >
      |
      <button
        class="header-button"
        on:click={() => handleOnClick("https://skar-software.github.io/pubdeskmd", "help")}>Help</button
      >
    </div>
  </header>
  <main>
    <slot />
  </main>
  <footer>
    <span style="font-weight: 400;">
      Powered by:
      <a
        class="foot"
        href="https://github.com/BearToCode/carta"
        >Carta Markdown Editor
      </a>
    </span>
    <span style="font-weight: 400;">
      credits:
      <a
        class="foot"
        href="https://github.com/sahildhingraa">@sahildhingra</a
      >
      <span> ,</span>
      <a
        class="foot"
        href="https://github.com/Wawilow">@wawilow</a
      >
    </span>
  </footer>
</div>

<style>
  header {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    padding: 10px;
    background-color: #edede9;
  }
  * {
    font-family: poppins;
  }
  .app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  main {
    margin: 0;
    padding: 0;
    width: 100%;
  }
  .foot,
  span {
    font-weight: 700;
    font-size: 1rem;
  }
  footer {
    overflow-x: hidden;
    position: fixed;
    bottom: 0%;
    width: 100vw;
    background: #778da9;
    color: aliceblue;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    align-items: center;
  }

  footer a {
    font-weight: bold;
    color: aliceblue;
    text-decoration-color: aliceblue;
  }

  @media (min-width: 480px) {
    footer {
      padding: 12px 0;
    }
  }
  header div .header-button {
    background-color: transparent;
    outline: none;
    border: none;
    font-size: 1rem;
    color: #415a77;
    font-weight: 600;
    padding: 0 10px;
  }
</style>
