<!-- Chatbox.svelte -->
<script>
    import { onMount } from 'svelte';
  
    let online = false;
    let messages = [
      {
        text: 'This is Flyhigh.ai! How may I assist you today?',
        from: 'agent',
      },
    ];
  
    onMount(() => {
      // Simulate checking online status (you can replace this with actual logic)
      setTimeout(() => {
        online = true;
      }, 2000);
    });
  
    function sendMessage() {
      const input = document.getElementById('messageInput');
      const text = input.value.trim();
      if (text !== '') {
        messages = [
        ...messages,
        { text, from: 'user' },
      ];
        input.value = '';
      }
    }
  </script>
  
  <style>
    .chatbox {
      position: fixed;
      bottom: 20px;
      right: 20px;
      width: 300px;
      background-color: white;
      border-radius: 8px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      padding: 16px;
      font-family: Arial, sans-serif;
    }
    .header {
      display: flex;
      align-items: center;
      margin-bottom: 16px;
      font-weight: bold;
    }
  
    .icon {
      width: 24px;
      height: 24px;
      margin-right: 8px;
    }
  
    .online {
      width: 8px;
      height: 8px;
      border-radius: 50%;
      background-color: green;
      margin-right: 4px;
    }
  
    .offline {
      width: 8px;
      height: 8px;
      border-radius: 50%;
      background-color: grey;
      margin-right: 4px;
    }
  
    .message {
      margin-bottom: 8px;
      padding: 8px;
      border-radius: 8px;
      background-color: #f0f0f0;
    }
  
    .user-message {
      background-color: #e0f3ff;
      text-align: right;
    }
  </style>
  
  <div class="chatbox">
    <div class="header">
      {#if online}
        <div class="online"></div>
      {:else}
        <div class="offline"></div>
      {/if}
      <img class="icon" src="/bird.png" alt="Chat Icon" />
      Flyhigh.ai
    </div>
    {#each messages as message}
      <div class="message {message.from === 'user' ? 'user-message' : ''}">
        {message.text}
      </div>
    {/each}
    <input id="messageInput" placeholder="Type your message..." />
    <button on:click={sendMessage}>Send</button>
  </div>
  