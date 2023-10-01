<!-- serp.svelte -->
<script>
  import { onMount, afterUpdate } from 'svelte';
  import SearchBar from './SearchBar.svelte';

  export let searchResults = [];
  let filteredResults = [];
  let searchQuery = '';

  // Function to update the filteredResults array based on the searchQuery
  function updateFilteredResults() {
    filteredResults = searchResults.filter(result =>
      result.title.toLowerCase().includes(searchQuery.toLowerCase())
    );
  }

  // Function to handle the user input in the search bar
  function handleSearchInput(event) {
    searchQuery = event.target.value;
    updateFilteredResults();
  }

  // Run the initial filtering when searchResults changes
  afterUpdate(updateFilteredResults);
  

</script>

<main>
    <SearchBar bind:searchQuery={searchQuery} on:cusEvent={handleSearchInput}/>
  {#if filteredResults.length > 0}
    <div class="results-container">
      {#each filteredResults as result}
        <div class="result-box">
          <h3>{result.title}</h3>
          <p>{result.description}</p>
          <a href={result.url} target="_blank" rel="noopener noreferrer">{result.url}</a>
        </div>
      {/each}
    </div>
  {:else}
    {#if searchQuery !== ''}
      <p>No results found for "{searchQuery}"</p>
    {/if}
  {/if}
</main>

<style>
  /* Your existing styles remain unchanged */
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 20px;
  }

  .results-container {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    grid-gap: 20px;
    max-width: 1200px;
    margin-top: 20px;
  }

  .result-box {
    background-color: #f9f9f9;
    border: 1px solid #ccc;
    padding: 10px;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .result-box h3 {
    margin-top: 0;
  }

  .result-box p {
    margin: 5px 0;
  }

  .result-box a {
    display: block;
    color: #0366d6;
    text-decoration: none;
  }

  .result-box a:hover {
    text-decoration: underline;
  }
  
</style>
