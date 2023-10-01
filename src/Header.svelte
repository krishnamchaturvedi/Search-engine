<!-- Header.svelte -->
<script>
  import Modal from './Modal.svelte';
  import Login from './Login.svelte';
  import SignUp from './SignUp.svelte';

  let showModal = false;
  let currentMode = 'login'; 

  function toggleModal(mode) {
    showModal = !showModal;
    currentMode = mode;
  }
</script>

<header class="header-container">
  <div class="menu-container">
  <nav>
    <ul>
      <li><a href="https://www.flyhigh.ai/" target="_blank">About</a></li>
      <li><button on:click={() => toggleModal('login')}>Login</button></li>
      <li><button on:click={() => toggleModal('signup')}>Signup</button></li>
    </ul>
  </nav>
  </div>>

  {#if showModal}
    <Modal show>
      {#if currentMode === 'login'}
        <Login />
      {:else if currentMode === 'signup'}
        <SignUp />
      {/if}
    </Modal>
  {/if}
</header>

<style>
  .header-container {
    background-color: #fce883;
    color: white;
    padding: 10px 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: fixed;
    top: 0;
    left: 0;
    width: 97%;
    z-index: 1000;
  }

  .menu-container {
    display: flex;
    align-items: center;
    margin-left: auto; 
  }

  nav ul {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    gap: 20px;
  }

  nav a {
    color: black;
    text-decoration: none;
    font-weight: bold;
    font-size: 16px;
    margin-bottom: 0;
    padding:0;
  }

  nav button {
    background-color: transparent;
    border: none;
    color: black;
    font-weight: bold;
    font-size: 16px;
    cursor: pointer;
    margin-bottom: 0;
    padding:0;
  }

  nav button:hover {
    background-color: rgba(255, 255, 255, 0.1);
  }
</style>

