<!-- Modal.svelte -->
<script>
  export let show = false;

  // Function to handle the keydown event
  function handleKeydown(event) {
    if (event.key === 'Escape') {
      show = false; // Close the modal when the Escape key is pressed
    }
  }
</script>

{#if show}
  <div class="modal" tabindex="-1" on:keydown={handleKeydown} on:click={() => (show = false)}>
    <div class="modal-content" on:click={(event) => event.stopPropagation()}>
      <span class="close" on:click={() => (show = false)}>&times;</span>
      <slot></slot>
    </div>
  </div>
{/if}
<style>
  /* Add your custom styling for the modal here */

  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5); /* Semi-transparent black background */
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999; /* Ensure the modal is on top of other elements */
  }

  .modal-content {
    background-color: #fff;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    max-width: 400px;
    width: 100%;
  }

  .close {
    position: absolute;
    top: 10px;
    right: 10px;
    font-size: 20px;
    cursor: pointer;
  }
</style>
