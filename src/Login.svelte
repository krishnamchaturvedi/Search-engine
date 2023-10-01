<!-- Login.svelte -->
<script>
  let username = '';
  let password = '';
  let message = '';

  function handleLogin() {
    if (username === 'admin' && password === 'password') {
      message = 'Login successful!';
    } else {
      message = 'Invalid username or password.';
    }
  }

  function resetForm() {
    username = '';
    password = '';
    message = '';
  }
</script>

<div class="login-form">
  <h2>Login</h2>
  {#if message}
    <p class="message">{message}</p>
  {/if}
  <form>
    <div class="form-group">
      <label for="username">Username/Email</label>
      <input type="text" id="username" bind:value={username} />
    </div>
    <div class="form-group">
      <label for="password">Password</label>
      <input type="password" id="password" bind:value={password} />
    </div>
    <div class="button-group">
      <button type="button" on:click={handleLogin}>Login</button>
      <button type="button" on:click={resetForm}>Reset</button>
    </div>
  </form>
</div>

<style>
  /* Add your custom styling for the login form here */

  .login-form {
    max-width: 300px;
    margin: 0 auto;
    padding: 20px;
    background-color: #f0f0f0;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  h2 {
    text-align: center;
    margin-bottom: 15px;
    color:black;
  }

  .form-group {
    margin-bottom: 15px;
  }

  label {
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
    color:black;
  }

  input {
    width: 100%;
    padding: 10px;
    border: 1px solid black;
    border-radius: 5px;
  }

  .button-group {
    display: flex;
    justify-content: space-between;
    margin-top: 20px;
  }

  button {
    flex: 1;
    padding: 10px;
    border: none;
    border-radius: 5px;
    background-color: #fce883;
    color: black;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.2s ease;
  }

  button:hover {
    background-color: grey;
  }

  .message {
    color: #ff0000;
    text-align: center;
    margin-bottom: 10px;
  }
</style>
