<script>
</script>

<div class="app">
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
  * {
    font-family: poppins;
  }
  .app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  main {
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
    text-decoration: none;
  }

  @media (min-width: 480px) {
    footer {
      padding: 12px 0;
    }
  }
</style>
