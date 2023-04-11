<script>
  import vaulterLogoLight from "./assets/vaulter-logo-white-black.svg";
  import vaulterLogoDark from "./assets/vaulter-logo-black-white.svg";
</script>

<main>
  <div>
    <a
      href="https://github.com/Chirag-And-Dheeraj/vaulter"
      target="_blank"
      rel="noreferrer"
    >
      <picture>
        <source srcset={vaulterLogoDark} media="(prefers-color-scheme: dark)" />
		<source srcset={vaulterLogoLight} media="(prefers-color-scheme: light)" />
		<img class="logo" alt="Vaulter logo">
      </picture>
    </a>
  </div>
  <h1>Vaulter</h1>

  <p>
    Check out <a
      href="https://docs.google.com/presentation/d/10mVA3UtZ7ypxRj_MYcE27NBlk2hSUNPV92l1zwNPKik"
      target="_blank"
      rel="noreferrer">Vaulter</a
    >, an open sourced, web-based file collection and storage server!
  </p>

  <p class="read-the-docs">
    Click on the Vaulter logo to take a look under the hood.
  </p>
</main>

<style>
  .logo {
    height: 10em;
    will-change: filter;
    transition: filter 250ms;
    border-radius: 10px;
  }
  .logo:hover {
    filter: drop-shadow(0 0 1em #dddddd);
  }
</style>
