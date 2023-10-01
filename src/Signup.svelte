<!-- Signup.svelte -->
<script>
    let name = '';
    let email = '';
    let number = '';
    let password = '';
    let confirmPassword = '';
    let signupOption = '';
  
    let passwordMatch = false;
  
    function handleSignup() {
      if (password === confirmPassword) {
        passwordMatch = true;
        // Perform signup process here
        showMessage('Congrats! You signed up successfully.');
      } else {
        passwordMatch = false;
      }
    }
  
    function showMessage(message) {
      alert(message);
    }
  
    function handleSignupWithGmail() {
      // Implement signup with Gmail here
      showMessage('Signup using Gmail is not implemented yet.');
    }
  </script>
  
  <main>
    <div class="signup-options">
      <button on:click={() => (signupOption = 'manual')}>Signup (Manual)</button>
      <button on:click={handleSignupWithGmail}>Signup using Gmail</button>
    </div>
  
    {#if signupOption === 'manual'}
      <div class="signup-form">
        <h2>Manual Signup</h2>
        <form>
          <div class="form-group">
            <label for="name">Name</label>
            <input type="text" id="name" bind:value={name} />
          </div>
          <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" bind:value={email} />
          </div>
          <div class="form-group">
            <label for="number">Number</label>
            <input type="text" id="number" bind:value={number} />
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" bind:value={password} />
            {#if passwordMatch}
              <span class="icon">&#10004;</span>
            {:else}
              <span class="icon">&#10006;</span>
            {/if}
          </div>
          <div class="form-group">
            <label for="confirmPassword">Confirm Password</label>
            <input type="password" id="confirmPassword" bind:value={confirmPassword} />
            {#if passwordMatch}
              <span class="icon">&#10004;</span>
            {:else}
              <span class="icon">&#10006;</span>
            {/if}
          </div>
          <button type="button" on:click={handleSignup}>Signup</button>
        </form>
      </div>
    {/if}
  </main>
  
  <style>
    main {
      display: flex;
      flex-direction: column;
      align-items: center;
      padding-top: 20px;
      color:black;
    }
  
    .signup-options {
      display: flex;
      justify-content: center;
      gap: 10px;
      margin-bottom: 20px;
      color: black;
    }
  
    button {
      padding: 10px 20px;
      background-color:#fce883;
      color: black;
      border: none;
      border-radius: 5px;
      cursor: pointer;
      font-weight: bold;
    }
  
    .signup-form {
      max-width: 400px;
      background-color: #f0f0f0;
      padding: 20px;
      border-radius: 5px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
  
    .form-group {
      display: flex;
      align-items: center;
      margin-bottom: 15px;
    }
  
    label {
      flex: 1;
      margin-right: 10px;
    }
  
    input {
      flex: 3;
      padding: 5px;
      border: 1px solid black;
      border-radius: 5px;
    }
  
    .icon {
      margin-left: 10px;
    }
  </style>
  