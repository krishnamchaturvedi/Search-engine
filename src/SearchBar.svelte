<script>
  export let inputField = "";
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  let listening = false;
  let recognition = null;

  // Function to handle search
  const handleMessage = (searchQuery) => {
    console.log(searchQuery);
    const eventData = {
      source: inputField
    };
    dispatch('cusEvent', eventData);
  };

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleMessage();
    }
  };
  const requestMicrophonePermission = async () => {
    try {
      await navigator.mediaDevices.getUserMedia({ audio: true });
    } catch (error) {
      // Microphone permission denied or not supported
      console.error("Error accessing microphone:", error);
      alert("Microphone permission denied or not supported.");
    }
  };

  // Function to start or stop speech recognition
  const toggleSpeechRecognition = () => {
    if (!recognition) {
      recognition = new (window.SpeechRecognition || window.webkitSpeechRecognition || window.mozSpeechRecognition || window.msSpeechRecognition)();
      recognition.lang = "en-US"; // Set the language you want to recognize

      recognition.onstart = () => {
        listening = true;
      };

      recognition.onend = () => {
        listening = false;
      };

      recognition.onresult = (event) => {
        const transcript = event.results[0][0].transcript;
        inputField = transcript; // Update the inputField value with the recognized speech
      };
    }
    listening = !listening;
    if (listening) {
      recognition.stop();
    } else {
      recognition.start();
    }
  };

  // Function to handle image scanner icon click
  const handleImageScannerClick = () => {
    const inputElement = document.getElementById('imageInput');
    inputElement.click(); // Trigger the file selection dialog
  };

  // Function to handle the selected image
  const handleImageSelect = (event) => {
    const file = event.target.files[0];
    if (file) {
      // Now you have access to the selected image file and can perform further processing if needed
      console.log(file);
    }
  };
  import { onMount } from 'svelte';
  onMount(() => {
    // Listen for file selection change
    document.getElementById('imageInput').addEventListener('change', handleImageSelect);
  });
</script>

<div class="search-bar">
  <img src="/bird.png" alt="Search Icon" class="search-icon" />
  <input type="text" bind:value={inputField} on:keydown={handleKeyDown} placeholder="Search for Schools, Colleges and More..." />
  <button on:click={()=>handleMessage(inputField)}>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      width="24"
      height="24"
      fill="grey"
    >
      <path
        d="M14.5 13h-.79l-.28-.27c1.39-1.54 2.07-3.59 1.74-5.67-.47-3.44-3.4-6.27-7.01-6.75A7.49 7.49 0 007.5 2a7.5 7.5 0 100 15 7.49 7.49 0 005.97-2.97l.28.27v.79l4.25 4.25 1.25-1.25-4.25-4.25zm-7 0a5.5 5.5 0 110-11 5.5 5.5 0 010 11z"
      />
    </svg>
  </button>
  <!-- Replace the button with the mic icon -->
  <button on:click={toggleSpeechRecognition} class="mic-scanner-icon">
    {#if listening}
      🎙️
    {:else}
      🎤
    {/if}
  </button>
  <input type="file" accept="image/*" style="display: none" id="imageInput" />
  <button on:click={handleImageScannerClick} class="mic-scanner-icon">
    <!-- Your SVG icon for the image scanner -->
    📷
  </button>
</div>

<style>
  .search-bar {
    display: flex;
    align-items: top;
    background-color: #f1f3f4;
    padding: 8px 16px;
    border-radius: 24px;
    width: 40%;
    margin: 10px auto;
    box-shadow: 0 0 5px 1px yellow !important;
  }

  .search-icon {
    width: 30px; /* Adjust the width of the search icon */
    height: 40px; /* Adjust the height of the search icon */
  }

  input {
    flex: 1;
    border: none;
    outline: none;
    font-size: 16px;
    padding: 8px;
    background-color: transparent;
  }

  button {
    background-color:  #f1f3f4;
    color: black;
    border: none;
    border-radius: 4px;
    margin-left: 4px;
    cursor: pointer;
    padding: 8px;
    display: flex;
    align-items: center;
  }

  button:hover {
    background-color: #f1f3f4;
  }
</style>
