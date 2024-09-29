<header class="header">
  <div class="logo">LOGO</div>
  <div class="searchbar"><input type="text" placeholder="Search Reddit" /></div>
  <div>
    <nav class="nav-icons">
      <a href="A">A</a>
      <a href="B">B</a>
      <a href="C">C</a>
      <a href="D">D</a>
      <a href="E">E</a>
    </nav>
  </div>
</header>

<style>
  .header {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    padding: 0 1rem;
    border-bottom: 1px solid #ffffff;
  }

  .logo {
    grid-column: 1;
  }

  .searchbar {
    grid-column: 2;
    text-align: center;
    height: 56px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .searchbar input {
    width: 100%;
    max-width: 500px;
    padding: 0.5rem;
    border: none; /* <-- This thing here */
    border: solid 1px #ccc;
    border-radius: 20px;
    background-color: #333d41;
    color: #8aa2ad;
  }

  .searchbar input::placeholder {
    color: #8aa2ad;
  }

  .nav-icons {
    grid-column: 3;
    display: flex;
    gap: 15px;

    &a {
      text-decoration: none;
    }
  }
</style>
